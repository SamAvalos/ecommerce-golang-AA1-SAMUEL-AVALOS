package main

import (
	"log"
	"net/http"
<<<<<<< HEAD

	"ecommerce/internal/controller"
	"ecommerce/internal/repository"
	"ecommerce/internal/service"
	"ecommerce/models"
	"ecommerce/routes"
)

// ... arriba los imports ...

func main() {
	prodRepo := repository.NewProductRepository()
	cartRepo := repository.NewCartRepository()
	orderRepo := repository.NewOrderRepository()

	// Seed
	prodRepo.Seed([]models.Product{
		{ID: "prod_1", Name: "Camiseta Classic", PriceCents: 1999, Stock: 20, Active: true},
		{ID: "prod_2", Name: "Pantalón Denim", PriceCents: 4599, Stock: 15, Active: true},
		{ID: "prod_3", Name: "Mug Coffee", PriceCents: 999, Stock: 50, Active: true},
	})

	productService := service.NewProductService(prodRepo)
	cartService := service.NewCartService(cartRepo, prodRepo)
	checkoutService := service.NewCheckoutService(prodRepo, cartRepo, orderRepo)

	prodCtrl := controller.NewProductController(productService)
	cartCtrl := controller.NewCartController(cartService, productService)
	orderCtrl := controller.NewOrderController(orderRepo) // ← pasa el interface directamente
	checkoutCtrl := controller.NewCheckoutController(checkoutService)

	mux := http.NewServeMux()
	routes.Register(mux, prodCtrl, cartCtrl, orderCtrl, checkoutCtrl)

	log.Println("API corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
=======
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
>>>>>>> f1095d5ef8dad5b15f0350e9512114e8146417e7
}
