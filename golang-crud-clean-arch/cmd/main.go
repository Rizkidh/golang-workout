package main

import (
	"fmt"
	"log"
	"net/http"

	"golang-crud-clean-arch/m/config"
	httpHandler "golang-crud-clean-arch/m/delivery/http" // Aliaskan package httpHandler
	"golang-crud-clean-arch/m/internal/repository"
	"golang-crud-clean-arch/m/internal/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"golang-crud-clean-arch/m/delivery/routes" // Import the routes package
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Menyajikan file statis di folder "views"
	e.Static("/static", "views")

	// Route untuk menampilkan index.html sebagai halaman utama
	e.GET("/", func(c echo.Context) error {
		return c.File("views/index.html")
	})

	// Tambahkan route untuk halaman repository
	e.GET("/repositories", func(c echo.Context) error {
		return c.File("views/repository.html")
	})
	// Inisialisasi database
	mongoClient := config.MongoConnect()

	// Inisialisasi repository
	repoRepo := repository.NewRepoRepository(mongoClient)
	userRepo := repository.NewUserRepository(mongoClient)

	// Inisialisasi usecase
	repoUsecase := usecase.NewRepositoryUsecase(repoRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)

	// Inisialisasi handler
	repoHandler := httpHandler.NewRepositoryHandler(repoUsecase) // Gunakan alias httpHandler
	userHandler := httpHandler.NewUserHandler(userUsecase)

	// Setup router
	r := chi.NewRouter()
	routes.SetupRoutes(r, repoHandler)
	routes.SetupUserRoutes(r, userHandler)

	fmt.Println("Server berjalan di port :9000")
	log.Fatal(http.ListenAndServe(":9000", r))
}
