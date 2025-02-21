package main

import (
	"log"
	"net/http"
	"warehouse/api/new/internal/database"
	"warehouse/api/new/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	productHandler := handlers.NewProductHandler(db)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/products", productHandler.CreateProduct)
	r.Post("/products/search", productHandler.SearchProductByName)
	r.Post("/shelves", productHandler.CreateShelf)

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
