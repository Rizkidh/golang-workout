package main

import (
	"fmt"
	"log"
	"net/http"

	"golang-crud-clean-arch/config"
	httpHandler "golang-crud-clean-arch/delivery/http" // Aliaskan package httpHandler
	"golang-crud-clean-arch/internal/repository"
	"golang-crud-clean-arch/internal/usecase"

	"golang-crud-clean-arch/delivery/routes" // Import the routes package

	"github.com/go-chi/chi/v5"
)

func main() {
	// Inisialisasi database
	mongoClient := config.MongoConnect()

	// Inisialisasi repository
	repoRepo := repository.NewRepoRepository(mongoClient)
	userRepo := repository.NewUserRepository(mongoClient)

	// Inisialisasi usecase
	repoUsecase := usecase.NewRepositoryUsecase(repoRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)

	// Inisialisasi handler
	repoHandler := httpHandler.NewRepositoryHandler(repoUsecase, userUsecase) // Gunakan alias httpHandler
	userHandler := httpHandler.NewUserHandler(userUsecase)

	// Setup router
	r := chi.NewRouter()
	routes.SetupRepositoryRoutes(r, repoHandler)
	routes.SetupUserRoutes(r, userHandler)

	fmt.Println("Server berjalan di port :9000")
	log.Fatal(http.ListenAndServe(":9000", r))
}
