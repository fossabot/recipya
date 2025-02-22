package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/reaper47/recipya/internal/app"
	"github.com/reaper47/recipya/internal/services"
	_ "github.com/reaper47/recipya/internal/templates"
	"github.com/reaper47/recipya/static"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
)

var imagesDir string

func init() {
	SessionData = make(map[uuid.UUID]int64)

	exe, err := os.Executable()
	if err != nil {
		return
	}
	imagesDir = filepath.Join(filepath.Dir(exe), "data", "images")
}

// NewServer creates a Server.
func NewServer(repository services.RepositoryService, email services.EmailService, files services.FilesService) *Server {
	srv := &Server{
		Repository: repository,
		Email:      email,
		Files:      files,
	}
	srv.mountHandlers()
	return srv
}

// Server is the web application's configuration object.
type Server struct {
	Router     *chi.Mux
	Repository services.RepositoryService
	Email      services.EmailService
	Files      services.FilesService
}

func (s *Server) mountHandlers() {
	r := chi.NewRouter()

	r.Get("/", s.indexHandler)
	r.Get("/recipes/{id:[1-9]([0-9])*}/share", s.recipeShareHandler)

	r.Route("/auth", func(r chi.Router) {
		r.With(s.mustBeLoggedInMiddleware).Post("/change-password", s.changePasswordHandler)
		r.Get("/confirm", s.confirmHandler)

		r.Route("/forgot-password", func(r chi.Router) {
			r.Get("/", s.forgotPasswordHandler)
			r.Post("/", s.forgotPasswordPostHandler)

			r.Route("/reset", func(r chi.Router) {
				r.Get("/", forgotPasswordResetHandler)
				r.Post("/", s.forgotPasswordResetPostHandler)
			})
		})

		r.Route("/login", func(r chi.Router) {
			r.Use(s.redirectIfLoggedInMiddleware)

			r.Get("/", loginHandler)
			r.Post("/", s.loginPostHandler)
		})

		r.Route("/register", func(r chi.Router) {
			r.Use(s.redirectIfLoggedInMiddleware)

			r.Get("/", registerHandler)
			r.Post("/", s.registerPostHandler)
			r.Post("/validate-email", s.registerPostEmailHandler)
			r.Post("/validate-password", s.registerPostPasswordHandler)
		})

		r.Post("/logout", s.logoutHandler)
	})

	r.Route("/recipes", func(r chi.Router) {
		r.Use(s.mustBeLoggedInMiddleware)

		r.Get("/", s.recipesHandler)

		r.Route("/{id:[1-9]([0-9])*}", func(r chi.Router) {
			r.Get("/", s.recipesViewHandler)
			r.Delete("/", s.recipeDeleteHandler)
			r.Post("/share", s.recipeSharePostHandler)

			r.Route("/edit", func(r chi.Router) {
				r.Get("/", s.recipesEditHandler)
				r.Put("/", s.recipesEditPostHandler)
			})
		})

		r.Route("/add", func(r chi.Router) {
			r.Get("/", recipesAddHandler)
			r.Post("/import", s.recipesAddImportHandler)

			r.Route("/manual", func(r chi.Router) {
				r.Get("/", s.recipeAddManualHandler)
				r.Post("/", s.recipeAddManualPostHandler)

				r.Route("/ingredient", func(r chi.Router) {
					r.Post("/", recipeAddManualIngredientHandler)
					r.Post("/{entry:[1-9]([0-9])*}", recipeAddManualIngredientDeleteHandler)
				})

				r.Route("/instruction", func(r chi.Router) {
					r.Post("/", recipeAddManualInstructionHandler)
					r.Post("/{entry:[1-9]([0-9])*}", recipeAddManualInstructionDeleteHandler)
				})
			})

			r.Post("/request-website", s.recipesAddRequestWebsiteHandler)
			r.Post("/website", s.recipesAddWebsiteHandler)
		})

		r.Get("/export", s.recipesExportHandler)
		r.Get("/supported-websites", s.recipesSupportedWebsitesHandler)
	})

	r.Group(func(r chi.Router) {
		r.Use(s.mustBeLoggedInMiddleware)

		r.Route("/settings", func(r chi.Router) {
			r.Get("/", s.settingsHandler)

			r.Post("/convert-automatically", s.settingsConvertAutomaticallyPostHandler)
			r.Post("/measurement-system", s.settingsMeasurementSystemsPostHandler)

			r.Route("/tabs", func(r chi.Router) {
				r.Get("/profile", settingsTabsProfileHandler)
				r.Get("/recipes", s.settingsTabsRecipesHandler)
			})
		})

		r.Get("/download/{tmpFile}", s.downloadHandler)
		r.Get("/user-initials", s.userInitialsHandler)
	})

	r.NotFound(notFoundHandler)

	staticFS := http.FileServer(http.FS(static.FS))
	r.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/static", staticFS).ServeHTTP(w, r)
	})

	r.Get("/data/images/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/data/images", http.FileServer(http.Dir(imagesDir))).ServeHTTP(w, r)
	})

	s.Router = r
}

// Run starts the web server.
func (s *Server) Run() {
	httpServer := &http.Server{
		Addr:              "0.0.0.0:" + strconv.Itoa(app.Config.Port),
		Handler:           s.Router,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       1 * time.Minute,
	}

	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		shutdownCtx, shutdownCancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer shutdownCancel()

		go func() {
			<-shutdownCtx.Done()
			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
				fmt.Println("forcing exit as graceful shutdown timed out")
				os.Exit(1)
			}
		}()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		serverStopCtx()
	}()

	fmt.Printf("Serving on %s\n", app.Config.Address())
	if err := httpServer.ListenAndServe(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	<-serverCtx.Done()
}
