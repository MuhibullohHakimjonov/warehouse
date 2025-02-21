package models

type Product struct {
	ID          int    `json:"id"`
	ProductName string `json:"product_name"`
	Type        string `json:"type"`
	Location    string `json:"location"`
}

type Shelf struct {
	ID        int    `json:"id"`
	ShelfName string `json:"shelf_name"`
	Path      string `json:"path"`
}
