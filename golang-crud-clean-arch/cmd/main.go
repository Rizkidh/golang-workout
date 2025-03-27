package main

import (
	"fmt"
	"log"
	"net/http"

	"golang-crud-clean-arch/config"
	httpHandler "golang-crud-clean-arch/delivery/http"
	"golang-crud-clean-arch/delivery/routes"
	"golang-crud-clean-arch/internal/repository"
	"golang-crud-clean-arch/internal/usecase"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Inisialisasi database
	mongoClient := config.MongoConnect()

	// Inisialisasi Redis
	redisClient := config.ConnectRedis()

	// Inisialisasi repository
	repoRepo := repository.NewRepoRepository(mongoClient, redisClient)
	userRepo := repository.NewUserRepository(mongoClient, redisClient)

	// Inisialisasi usecase
	repoUsecase := usecase.NewRepositoryUsecase(repoRepo, redisClient)
	userUsecase := usecase.NewUserUsecase(userRepo, redisClient)

	// Inisialisasi handler
	repoHandler := httpHandler.NewRepositoryHandler(repoUsecase, userUsecase)
	userHandler := httpHandler.NewUserHandler(userUsecase)

	// Setup router
	r := chi.NewRouter()
	routes.SetupRepositoryRoutes(r, repoHandler)
	routes.SetupUserRoutes(r, userHandler)

	fmt.Println("Server berjalan di port :9000")
	log.Fatal(http.ListenAndServe(":9000", r))
}
