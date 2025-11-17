package main

import (
	"log"
	"net/http"
	"time"

	"ecommerce/models"
	"ecommerce/routes"
	"ecommerce/services"
)

func main() {
	mux := http.NewServeMux()

	// services
	ps := services.NewProductsService()
	cs := services.NewCartService()
	os := services.NewOrdersService(ps)

	// seed products
	ps.Seed([]models.Product{
		{ID: "prod_1", Name: "Camiseta Classic", PriceCents: 1999, Stock: 20, Active: true, CreatedAt: time.Now().Unix()},
		{ID: "prod_2", Name: "Pantalón Denim", PriceCents: 4599, Stock: 15, Active: true, CreatedAt: time.Now().Unix()},
		{ID: "prod_3", Name: "Mug Coffee", PriceCents: 999, Stock: 50, Active: true, CreatedAt: time.Now().Unix()},
	})

	// register routes
	routes.Register(mux, ps, cs, os)

	// health check
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ecommerce API (in-memory). Endpoints: /api/products, /api/cart/{id}, /api/checkout/{id}, /api/orders"))
	})

	addr := ":8080"
	log.Printf("Server running at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
