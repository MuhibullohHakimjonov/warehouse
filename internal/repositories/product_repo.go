package repositories

import (
	"database/sql"
	"fmt"
	"warehouse/api/new/internal/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(product *models.Product) error {
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM shelves WHERE path = $1)`
	err := r.db.QueryRow(checkQuery, product.Location).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("location '%s' does not exist in shelves", product.Location)
	}

	query := `INSERT INTO products (product_name, type, location) VALUES ($1, $2, $3) RETURNING id`
	return r.db.QueryRow(query, product.ProductName, product.Type, product.Location).Scan(&product.ID)
}

func (r *ProductRepository) GetProductByName(name string) (*models.Product, error) {
	product := &models.Product{}
	query := `SELECT id, product_name, type, location FROM products WHERE product_name = $1`
	err := r.db.QueryRow(query, name).Scan(&product.ID, &product.ProductName, &product.Type, &product.Location)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) CreateShelf(shelf *models.Shelf) error {
	query := `INSERT INTO shelves (shelf_name, path) VALUES ($1, $2) RETURNING id`
	return r.db.QueryRow(query, shelf.ShelfName, shelf.Path).Scan(&shelf.ID)
}
