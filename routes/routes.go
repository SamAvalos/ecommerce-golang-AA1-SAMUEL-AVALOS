package routes

import (
	"ecommerce/controllers"
	"ecommerce/services"
	"fmt"
	"net/http"
)

func Register(mux *http.ServeMux, ps *services.ProductsService, cs *services.CartService, os *services.OrdersService) {
	pc := controllers.NewProductsController(ps)
	cc := controllers.NewCartController(cs, ps, os)
	oc := controllers.NewOrdersController(os)

	// Products
	mux.HandleFunc("/api/products", pc.List)

	// Cart example (userID hardcoded for simplicity)
	userID := "user_1"
	mux.HandleFunc(fmt.Sprintf("/api/cart/%s", userID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			cc.GetCart(w, r, userID)
			return
		}
		if r.Method == http.MethodPost {
			r.ParseForm()
			productID := r.Form.Get("product_id")
			qty := 1
			cc.AddToCart(w, r, userID, productID, qty)
			return
		}
		http.NotFound(w, r)
	})

	// Orders
	mux.HandleFunc("/api/orders", oc.List)
}
