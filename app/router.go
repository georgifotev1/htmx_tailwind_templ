package app

import (
	"net/http"

	components "github.com/georgifotev1/wms/internal/templates"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (a *App) newRouter() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		components.Hello().Render(r.Context(), w)
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}

		name := r.FormValue("name")
		component := components.HelloPost(name)
		err = component.Render(r.Context(), w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	a.router = r
}
