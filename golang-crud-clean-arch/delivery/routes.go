package routes

import (
	"golang-crud-clean-arch/m/delivery/http"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux, repoHandler *http.RepositoryHandler, userHandler *http.UserHandler) {
	r.Route("/repositories", func(r chi.Router) {
		r.Post("/", repoHandler.Create)
		r.Get("/{id}", repoHandler.GetAll)
		r.Put("/{id}", repoHandler.Update)
		r.Delete("/{id}", repoHandler.Delete)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.Create)
		r.Get("/{id}", userHandler.Get)
		r.Put("/{id}", userHandler.Update)    // Tambah endpoint update user
		r.Delete("/{id}", userHandler.Delete) // Tambah endpoint delete user
	})
}
