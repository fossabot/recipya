package server

import (
	"github.com/reaper47/recipya/internal/templates"
	"github.com/reaper47/recipya/internal/units"
	"net/http"
	"slices"
)

func (s *Server) settingsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Hx-Request") == "true" {
		templates.RenderComponent(w, "core", "settings", nil)
	} else {
		page := templates.SettingsPage
		templates.Render(w, page, templates.Data{
			IsAuthenticated: true,
			Title:           page.Title(),
		})
	}
}

func (s *Server) settingsConvertAutomaticallyPostHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int64)
	isConvert := r.FormValue("convert") == "on"
	if err := s.Repository.UpdateConvertMeasurementSystem(userID, isConvert); err != nil {
		w.Header().Set("HX-Trigger", makeToast("Failed to set setting.", errorToast))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	msg := "Recipe conversion disabled."
	if isConvert {
		msg = "Recipe conversion enabled."
	}
	w.Header().Set("HX-Trigger", makeToast(msg, infoToast))
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) settingsMeasurementSystemsPostHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int64)
	systems, settings, err := s.Repository.MeasurementSystems(userID)
	if err != nil {
		w.Header().Set("HX-Trigger", makeToast("Error fetching units systems.", errorToast))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	system := units.System(r.FormValue("system"))
	if !slices.Contains(systems, system) {
		w.Header().Set("HX-Trigger", makeToast("Measurement system does not exist.", errorToast))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if settings.MeasurementSystem == system {
		w.Header().Set("HX-Trigger", makeToast("System already set to "+system.String()+".", warningToast))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := s.Repository.SwitchMeasurementSystem(system, userID); err != nil {
		w.Header().Set("HX-Trigger", makeToast("Error switching units system.", errorToast))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", makeToast("Measurement system set to "+system.String()+".", infoToast))
	w.WriteHeader(http.StatusNoContent)
}

func settingsTabsProfileHandler(w http.ResponseWriter, _ *http.Request) {
	templates.RenderComponent(w, "core", "settings-tabs-profile", nil)
}

func (s *Server) settingsTabsRecipesHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int64)
	systems, settings, err := s.Repository.MeasurementSystems(userID)
	if err != nil {
		w.Header().Set("HX-Trigger", makeToast("Error fetching units systems.", errorToast))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	templates.RenderComponent(w, "core", "settings-tabs-recipes", templates.Data{
		Settings: templates.SettingsData{
			MeasurementSystems: systems,
			UserSettings:       settings,
		},
	})
}
