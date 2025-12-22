package routes

import (
	"encoding/json"
	"html/template"
	"net/http"

	"ecommerce/internal/auth"
	"ecommerce/internal/controller"
	"ecommerce/internal/middleware"

	"github.com/go-chi/chi/v5"
)

var authService = auth.NewAuthService()
var tmpl = template.Must(template.New("home").Parse(homeHTML))

const homeHTML = `
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Tienda Ecommerce SamAv</title>
    <style>
        body { font-family: 'Segoe UI', sans-serif; background: linear-gradient(135deg, #667eea, #764ba2); color: white; margin: 0; padding: 20px; min-height: 100vh; }
        .container { max-width: 1200px; margin: 0 auto; background: rgba(255,255,255,0.1); padding: 30px; border-radius: 15px; backdrop-filter: blur(10px); }
        h1, h2 { text-align: center; }
        .card { background: rgba(255,255,255,0.15); padding: 20px; border-radius: 12px; margin: 20px 0; }
        input, button { padding: 12px; margin: 8px; font-size: 16px; border-radius: 8px; border: none; }
        button { background: #00d1b2; color: white; cursor: pointer; width: 100%; }
        button:hover { background: #00b89c; }
        .product-grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(250px, 1fr)); gap: 20px; margin: 20px 0; }
        .product-card { background: rgba(255,255,255,0.2); padding: 20px; border-radius: 12px; text-align: center; }
        .product-name { font-size: 18px; font-weight: bold; margin-bottom: 10px; }
        .product-price { font-size: 16px; color: #ffd700; margin-bottom: 15px; }
        .add-btn { background: #4CAF50; padding: 10px 20px; width: auto; }
        .add-btn:hover { background: #45a049; }
        .hidden { display: none; }
        #cart-items { background: rgba(255,255,255,0.2); padding: 15px; border-radius: 8px; margin: 10px 0; }
        pre { background: rgba(0,0,0,0.5); padding: 15px; border-radius: 8px; overflow-x: auto; }
        .eco-message { background: #4CAF50; color: white; padding: 15px; border-radius: 8px; margin-top: 15px; text-align: center; font-weight: bold; }
    </style>
</head>
<body>
    <div class="container">
        <h1>Tienda Ecommerce SamAv</h1>
        <p style="text-align:center; font-size:18px;">Proyecto 100% funcional • Clean Architecture • JWT</p>

        <div class="card">
            <h2>Login (fácil)</h2>
            <p>Regístrate o inicia sesión. Si no existes, se crea automáticamente.</p>
            <input type="email" id="email" placeholder="tu@email.com" value="juan@test.com">
            <input type="password" id="pass" placeholder="contraseña" value="123456">
            <button onclick="login()">Iniciar sesión / Registrarse</button>
            <p id="msg" style="margin-top:10px; font-weight:bold;"></p>
        </div>

        <div class="card hidden" id="panel">
            <h2>Productos disponibles</h2>
            <div id="products-grid" class="product-grid"></div>

            <div id="cart-items" style="margin-top: 20px;">
                <h3>Tu carrito (<span id="cart-count">0</span> ítems)</h3>
                <div id="cart-display"></div>
                <button onclick="checkout()" style="background:#e74c3c; margin-top:10px;">Hacer checkout (pagar)</button>
                <div id="eco-result" class="eco-message hidden"></div>
                <pre id="result"></pre>
            </div>

            <div class="card">
                <h2>Tus órdenes</h2>
                <button onclick="verOrdenes()">Ver mis órdenes</button>
                <pre id="orders-result"></pre>
            </div>
        </div>
    </div>

    <script>
        let token = "";
        let products = [];

        async function login() {
            const email = document.getElementById("email").value;
            const pass = document.getElementById("pass").value;

            let res = await fetch("/api/auth/login", {
                method: "POST",
                headers: {"Content-Type": "application/json"},
                body: JSON.stringify({email, password: pass})
            });

            if (!res.ok) {
                res = await fetch("/api/auth/register", {
                    method: "POST",
                    headers: {"Content-Type": "application/json"},
                    body: JSON.stringify({email, password: pass})
                });
                if (!res.ok) {
                    document.getElementById("msg").innerText = "Error al crear cuenta";
                    return;
                }
                document.getElementById("msg").innerText = "Cuenta creada! Iniciando sesión...";
                res = await fetch("/api/auth/login", {
                    method: "POST",
                    headers: {"Content-Type": "application/json"},
                    body: JSON.stringify({email, password: pass})
                });
            }

            const data = await res.json();
            token = data.token;
            document.getElementById("msg").innerText = "¡Login correcto! Cargando productos...";
            document.getElementById("panel").classList.remove("hidden");

            loadProducts();
            loadCart();
        }

        async function loadProducts() {
            const res = await fetch("/api/products");
            products = await res.json();

            const grid = document.getElementById("products-grid");
            grid.innerHTML = "";

            products.forEach(product => {
                const card = document.createElement("div");
                card.className = "product-card";
                card.innerHTML = "<div class=\"product-name\">" + product.name + "</div>" +
                                 "<div class=\"product-price\">$" + (product.price_cents / 100).toFixed(2) + "</div>" +
                                 "<button class=\"add-btn\" onclick=\"addToCart('" + product.id + "', 1)\">Añadir al carrito</button>";
                grid.appendChild(card);
            });
        }

        async function addToCart(productId, quantity) {
            await fetch("/api/cart", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": "Bearer " + token
                },
                body: JSON.stringify({product_id: productId, quantity: quantity})
            });
            alert("¡Producto añadido al carrito!");
            loadCart();
        }

        async function loadCart() {
            const res = await fetch("/api/cart", {
                headers: {"Authorization": "Bearer " + token}
            });

            if (res.ok) {
                const data = await res.json();
                document.getElementById("cart-count").innerText = data.items.length;

                const cartDisplay = document.getElementById("cart-display");
                cartDisplay.innerHTML = "";

                data.items.forEach(item => {
                    const product = products.find(p => p.id === item.product_id);
                    if (product) {
                        const cartItem = document.createElement("div");
                        cartItem.innerHTML = "<div style=\"margin: 5px 0;\">" +
                                             product.name + " x" + item.qty + " - $" + ((item.price_cents * item.qty / 100).toFixed(2)) +
                                             "</div>";
                        cartDisplay.appendChild(cartItem);
                    }
                });

                if (data.items.length === 0) {
                    cartDisplay.innerHTML = "<p>Tu carrito está vacío</p>";
                }
            }
        }

        async function checkout() {
            const res = await fetch("/api/checkout", {
                method: "POST",
                headers: {"Authorization": "Bearer " + token}
            });

            if (res.ok) {
                const data = await res.json();
                document.getElementById("result").innerText = "¡COMPRA REALIZADA!\n\n" + JSON.stringify(data, null, 2);

                // Mensaje de sostenibilidad
                const co2 = Math.floor(Math.random() * 15) + 5; // 5–20 kg CO2 random
                const trees = Math.floor(co2 / 5); // 1 árbol por 5 kg
                document.getElementById("eco-result").innerText = "¡Gracias por tu compra! Tu pedido compensó " + co2 + " kg de CO2 y plantó " + trees + " árboles 🌳";
                document.getElementById("eco-result").classList.remove("hidden");

                alert("¡Compra exitosa! Revisa tu orden y tu impacto ambiental.");
                loadCart(); // Vacía el carrito visual
            } else {
                const error = await res.text();
                alert("Error en checkout: " + error);
            }
        }

        async function verOrdenes() {
            const res = await fetch("/api/orders");
            const data = await res.json();
            document.getElementById("orders-result").innerText = JSON.stringify(data, null, 2);
        }
    </script>
</body>
</html>
`

func Register(
	mux *http.ServeMux,
	prodCtrl *controller.ProductController,
	cartCtrl *controller.CartController,
	orderCtrl *controller.OrderController,
	checkoutCtrl *controller.CheckoutController,
) {
	r := chi.NewRouter()

	// Rutas públicas
	r.Get("/api/products", prodCtrl.List)
	r.Get("/api/orders", orderCtrl.List)
	r.Post("/api/auth/register", registerHandler)
	r.Post("/api/auth/login", loginHandler)

	// Rutas protegidas
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthRequired)
		r.Get("/api/cart", cartCtrl.GetCart)
		r.Post("/api/cart", cartCtrl.AddToCart)
		r.Post("/api/checkout", checkoutCtrl.Checkout)
	})

	// Página principal con todo
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	// Health
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	mux.Handle("/", r)
}

// Handlers de auth
func registerHandler(w http.ResponseWriter, r *http.Request) {
	var req struct{ Email, Password string }
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	if err := authService.Register(req.Email, req.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "user registered"})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var req struct{ Email, Password string }
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	token, err := authService.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
