"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[4781],{3905:(e,t,n)=>{n.d(t,{Zo:()=>c,kt:()=>m});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function s(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function o(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var l=r.createContext({}),p=function(e){var t=r.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):s(s({},t),e)),n},c=function(e){var t=p(e.components);return r.createElement(l.Provider,{value:t},e.children)},d="mdxType",u={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},h=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,i=e.originalType,l=e.parentName,c=o(e,["components","mdxType","originalType","parentName"]),d=p(n),h=a,m=d["".concat(l,".").concat(h)]||d[h]||u[h]||i;return n?r.createElement(m,s(s({ref:t},c),{},{components:n})):r.createElement(m,s({ref:t},c))}));function m(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=n.length,s=new Array(i);s[0]=h;var o={};for(var l in t)hasOwnProperty.call(t,l)&&(o[l]=t[l]);o.originalType=e,o[d]="string"==typeof e?e:a,s[1]=o;for(var p=2;p<i;p++)s[p]=n[p];return r.createElement.apply(null,s)}return r.createElement.apply(null,n)}h.displayName="MDXCreateElement"},6646:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>s,default:()=>u,frontMatter:()=>i,metadata:()=>o,toc:()=>p});var r=n(7462),a=(n(7294),n(3905));const i={},s="Add an Endpoint",o={unversionedId:"development/workflow/add-endpoint",id:"development/workflow/add-endpoint",title:"Add an Endpoint",description:"It is essential to grasp how the server works before bringing any modifications.",source:"@site/docs/development/workflow/add-endpoint.md",sourceDirName:"development/workflow",slug:"/development/workflow/add-endpoint",permalink:"/docs/development/workflow/add-endpoint",draft:!1,editUrl:"https://github.com/reaper47/recipya/tree/main/website/docs/docs/development/workflow/add-endpoint.md",tags:[],version:"current",frontMatter:{},sidebar:"docs",previous:{title:"Workflow",permalink:"/docs/category/workflow"},next:{title:"Support a Website",permalink:"/docs/development/workflow/import-website"}},l={},p=[{value:"The Server",id:"the-server",level:2},{value:"Example",id:"example",level:2},{value:"Router",id:"router",level:3},{value:"Test",id:"test",level:3},{value:"Handler",id:"handler",level:3},{value:"Repository",id:"repository",level:3}],c={toc:p},d="wrapper";function u(e){let{components:t,...n}=e;return(0,a.kt)(d,(0,r.Z)({},c,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"add-an-endpoint"},"Add an Endpoint"),(0,a.kt)("p",null,"It is essential to grasp how the server works before bringing any modifications.\nThen, we will guide you through the process of incorporating an HTTP endpoint into the server."),(0,a.kt)("h2",{id:"the-server"},"The Server"),(0,a.kt)("p",null,"Recipya's server code is located within the ",(0,a.kt)("inlineCode",{parentName:"p"},"internal/server")," package. The main file\nis ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/reaper47/recipya/blob/main/internal/server/server.go"},"server.go"),". It exports a single with\na receiver function and a corresponding struct."),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"Server"),": This struct holds the HTTP router, the repository, the email service and the files service. You can\nfind the declaration of each service in the ",(0,a.kt)("a",{parentName:"li",href:"https://github.com/reaper47/recipya/blob/main/internal/services/service.go"},"internal/services/service.go"),"\nfile."),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"Server.Run"),": Starts the web server."),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"NewServer"),": Creates a server that is ready for use. It requires the services to be passed as arguments.")),(0,a.kt)("p",null,"The HTTP router is initialized during the server's creation. We use ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/go-chi/chi"},"chi")," due to\nits simplicity in organizing endpoints. Please read the ",(0,a.kt)("inlineCode",{parentName:"p"},"mountHandlers")," function to observe the router in action."),(0,a.kt)("h2",{id:"example"},"Example"),(0,a.kt)("p",null,"Let's walk through an example aimed at adding an endpoint that searches for recipes."),(0,a.kt)("h3",{id:"router"},"Router"),(0,a.kt)("p",null,"The first step involves adding the endpoint to the router. A suitable endpoint is ",(0,a.kt)("inlineCode",{parentName:"p"},"GET /recipes/search?q=query"),". Open the\n",(0,a.kt)("a",{parentName:"p",href:"https://github.com/reaper47/recipya/blob/main/internal/server/server.go"},"internal/server/server.go")," file and include the\nendpoint within the ",(0,a.kt)("inlineCode",{parentName:"p"},"/recipes")," route block. The handler should be named ",(0,a.kt)("inlineCode",{parentName:"p"},"recipesSearchHandler"),", following the\n",(0,a.kt)("inlineCode",{parentName:"p"},"{resource}{LastWordEndpoint}{Handler}")," naming convention."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'r.Route("/recipes", func(r chi.Router) {\n  r.Use(s.mustBeLoggedInMiddleware)\n  \n  r.Get("/search", recipesSearchHandler)\n  ...\n})\n')),(0,a.kt)("p",null,"With the route established, it is time to declare the handler. Since we are dealing with the ",(0,a.kt)("inlineCode",{parentName:"p"},"/recipes")," resource,\nadd the handler to the ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/reaper47/recipya/blob/main/internal/server/handlers_recipes.go"},"handlers_recipes.go"),"\nfile. "),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'func (s *Server) recipesSearchHandler(w http.ResponseWriter, r *http.Request) {\n    panic("TODO: To implement")\n}\n')),(0,a.kt)("p",null,"We are now ready to create tests for our route."),(0,a.kt)("h3",{id:"test"},"Test"),(0,a.kt)("p",null,"Tests related to the server are written in the ",(0,a.kt)("inlineCode",{parentName:"p"},"handlers_{resource}_test.go")," files. The tests for our handlers are\nstored in the ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/reaper47/recipya/blob/main/internal/server/handlers_recipes_test.go"},"handlers_recipes_test.go"),"\nfile. The nomenclature for test function is ",(0,a.kt)("inlineCode",{parentName:"p"},"TestHandlers_{Resource}_{Endpoint}"),". Let's write the foundation\nfunction of our tests."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'func TestHandlers_Recipes_AddManual(t *testing.T) {\n    srv := newServerTest()\n\n    uri := "/recipes/search"\n}\n')),(0,a.kt)("p",null,"The subsequent step involves writing the different tests that add value to the users."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'func TestHandlers_Recipes_AddManual(t *testing.T) {\n    srv := newServerTest()\n\n    uri := "/recipes/search"\n\n    t.Run("must be logged in", func(t *testing.T) {\n        assertMustBeLoggedIn(t, srv, http.MethodGet, uri)\n    })\n    \n    t.Run("search fails", func(t *testing.T) {\n        t.Fail()\n    })\n    \n    t.Run("user has no recipes", func(t *testing.T) {\n        t.Fail()\n    })\n\n    t.Run("user searches empty string", func(t *testing.T) {\n        t.Fail()\n    })\n\n    testcases := []struct {\n        name string\n        in   string\n        want models.Recipes\n    }{\n        {name: "user searches empty string", in: "", want: ...},\n        {name: "user searches for lunch", in: "lunch", want: ...},\n        ...\n    }\n    for _, tc := range testcases {\n        t.Run(tc.name, func(t *testing.T) {\n            rr := sendHxRequestAsLoggedIn(srv, http.MethodGet, uri+"?q="+tc.in, noHeader, nil)\n            \n            ...\n        })\n    }\n}\n')),(0,a.kt)("p",null,"The body of the tests is omitted for brevity. Run the tests to ensure that they fail."),(0,a.kt)("h3",{id:"handler"},"Handler"),(0,a.kt)("p",null,"The next step entails crafting the handler's code. Return to the ",(0,a.kt)("inlineCode",{parentName:"p"},"handles_recipes.go")," file, and implement\nthe ",(0,a.kt)("inlineCode",{parentName:"p"},"recipesSearchHandler")," function that will make the tests go green. For instance, the implementation could resemble\nthe following. "),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'func (s *Server) recipesSearchHandler(w http.ResponseWriter, r *http.Request) {\n    userID := r.Context().Value("userID").(int64)\n    query := chi.URLParam(r, "q")\n    \n    recipes, err := s.Repository.SearchRecipes(query, userID)\n    if err != nil {\n        w.Header().Set("HX-Trigger", makeToast("Failed to search recipes.", errorToast))\n        w.WriteHeader(http.StatusInternalServerError)\n        return\n    }\n\n    templates.RenderComponent(w, "recipes", "search-recipes", templates.Data{Recipes: recipes})\n}\n')),(0,a.kt)("p",null,"This code gets the logged-in user's ID and the search query, then passes them to the ",(0,a.kt)("inlineCode",{parentName:"p"},"SearchRecipes")," function of the\nrepository. If this function encounters an error, an HTMX toast is sent to the user, accompanied by a HTTP 500\nstatus code. Otherwise, the HTML containing the recipes is sent."),(0,a.kt)("p",null,"The ",(0,a.kt)("inlineCode",{parentName:"p"},"templates.RenderComponent")," function displays a template from the ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/reaper47/recipya/tree/main/web/templates/components"},"web/templates/components"),"\ndirectory. Its second parameter is the name of one of the files within that directory, excluding the extension.\nThe third parameter is the name of the template within a file in that folder. Lastly, the fourth parameter is\na ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/reaper47/recipya/blob/main/internal/templates/data.go"},"struct")," containing data for the\nGoHTML template."),(0,a.kt)("h3",{id:"repository"},"Repository"),(0,a.kt)("p",null,"The final piece of the puzzle involves writing the ",(0,a.kt)("inlineCode",{parentName:"p"},"s.Repository.SearchRecipes")," function. The repository is an interface\nthat declares functions for interacting with a database. Currently, Recipya supports ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/reaper47/recipya/blob/main/internal/services/sqlite_service.go"},"sqlite"),"\nonly. To support other databases, we need define a struct satisf the ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/reaper47/recipya/blob/main/internal/services/service.go"},"RepositoryService"),"\ninterface."),(0,a.kt)("p",null,"Let's declare the function within the ",(0,a.kt)("inlineCode",{parentName:"p"},"RepositoryService")," interface. The functions are declared alphabetically."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"type RepositoryService interface {\n    // AddAuthToken adds an authentication token to the database.\n    AddAuthToken(selector, validator string, userID int64) error\n    \n    // AddRecipe adds a recipe to the user's collection.\n    AddRecipe(r *models.Recipe, userID int64) (int64, error)\n    \n    ...\n    \n    // SearchRecipes gets the user's recipes that include the search query.\n    SearchRecipes(query string, userID int64) (models.Recipes, error)\n    \n    ...\n    \n    // VerifyLogin checks whether the user provided correct login credentials.\n    // If yes, their user ID will be returned. Otherwise, -1 is returned.\n    VerifyLogin(email, password string) int64\n    \n    // Websites gets the list of supported websites from which to extract the recipe.\n    Websites() models.Websites\n}\n")),(0,a.kt)("p",null,"Subsequently, let's implement the function within the ",(0,a.kt)("inlineCode",{parentName:"p"},"sqlite_service.go")," file."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"func (s *SQLiteService) SearchRecipes(query string, userID int64) (models.Recipes, error) {\n    // s.Mutex.Lock()\n    // defer s.Mutex.Unlock()\n\n    ctx, cancel := context.WithTimeout(context.Background(), shortCtxTimeout)\n    defer cancel()\n\n    rows, err := s.DB.QueryContext(ctx, statements.SelectSearchRecipes, query, userID)\n    if err != nil {\n        return nil, err\n    }\n    defer rows.Close()\n    \n    var recipes models.Recipes\n    for rows.Next() {\n        // code to scan a recipe\n        ...\n        recipes = append(recipes, c)\n    }\n    return recipes, nil\n}\n")),(0,a.kt)("p",null,"Remember, invoking ",(0,a.kt)("inlineCode",{parentName:"p"},"s.Mutex.Lock()")," and ",(0,a.kt)("inlineCode",{parentName:"p"},"defer s.Mutex.Unlock()")," is necessary when inserting, updating, or deleting\ndatabase entries. However, in this scenario, we're merely fetching data, rendering the mutex unnecessary."),(0,a.kt)("p",null,"SQL statements are organized by action within the ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/reaper47/recipya/tree/main/internal/services/statements"},"internal/services/statements"),"\ndirectory. Their nomenclature is ",(0,a.kt)("inlineCode",{parentName:"p"},"{Action}{Resource}"),". In our case, a ",(0,a.kt)("inlineCode",{parentName:"p"},"SELECT")," statement for fetching recipes is termed\n",(0,a.kt)("inlineCode",{parentName:"p"},"SelectSearchRecipes")," and would reside in the\n",(0,a.kt)("a",{parentName:"p",href:"https://github.com/reaper47/recipya/tree/main/internal/services/statements"},"select.go")," file."),(0,a.kt)("p",null,"Every statement is a ",(0,a.kt)("inlineCode",{parentName:"p"},"const")," whose value is the SQLite statement itself. If Recipya ever supports other databases, we\nshall find a way to organize the statements per database type. "),(0,a.kt)("p",null,"You can call it a day and open a PR once you wrote your SQL and the tests pass."))}u.isMDXComponent=!0}}]);