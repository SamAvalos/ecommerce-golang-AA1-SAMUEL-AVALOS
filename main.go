package main

import (
	"ecommerce/internal/controller"
	"ecommerce/internal/repository"
	"ecommerce/internal/service"
	"ecommerce/models"
	"ecommerce/routes"
	"log"
	"net/http"
)

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
}
