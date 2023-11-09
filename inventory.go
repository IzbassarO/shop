package main

import "fmt"

type Product struct {
	ID    int
	Name  string
	Price float64
}

type InventoryManager struct {
	Products []Product
}

func NewInventoryManager() *InventoryManager {
	return &InventoryManager{
		Products: []Product{
			{ID: 1, Name: "Product 1", Price: 10.0},
			{ID: 2, Name: "Product 2", Price: 20.0},
			{ID: 3, Name: "Product 3", Price: 30.0},
			{ID: 4, Name: "Product 4", Price: 40.0},
			{ID: 5, Name: "Product 5", Price: 50.0},
			{ID: 6, Name: "Product 6", Price: 60.0},
		},
	}
}

func (im *InventoryManager) ListProducts() {
	for _, p := range im.Products {
		fmt.Printf("ID: %d, Name: %s, Price: $%.2f\n", p.ID, p.Name, p.Price)
	}
}

func (im *InventoryManager) GetProductByID(productID int) *Product {
	for _, p := range im.Products {
		if p.ID == productID {
			return &p
		}
	}
	return nil
}
