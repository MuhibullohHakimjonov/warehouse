package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"warehouse/api/new/internal/models"
	"warehouse/api/new/internal/repositories"
	"warehouse/api/new/internal/services"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(db *sql.DB) *ProductHandler {
	repo := repositories.NewProductRepository(db)
	service := services.NewProductService(repo)
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&product); err != nil {
		log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if product.ProductName == "" {
		http.Error(w, "Field 'product_name' is required", http.StatusBadRequest)
		return
	}
	if product.Type == "" {
		http.Error(w, "Field 'type' is required", http.StatusBadRequest)
		return
	}
	if product.Location == "" {
		http.Error(w, "Field 'location' is required", http.StatusBadRequest)
		return
	}
	if err := h.service.CreateProduct(&product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) SearchProductByName(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if requestBody.Name == "" {
		http.Error(w, "Field name is empty", http.StatusBadRequest)
		return
	}

	product, err := h.service.GetProductByName(requestBody.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Location string `json:"location"`
	}{
		Location: product.Location,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *ProductHandler) CreateShelf(w http.ResponseWriter, r *http.Request) {
	var shelf models.Shelf

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&shelf); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if shelf.ShelfName == "" {
		http.Error(w, "Field 'shelf_name' is required", http.StatusBadRequest)
		return
	}
	if shelf.Path == "" {
		http.Error(w, "Field 'path' is required", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateShelf(&shelf); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(shelf)
}
