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
        .container { max-width: 1000px; margin: 0 auto; background: rgba(255,255,255,0.1); padding: 30px; border-radius: 15px; backdrop-filter: blur(10px); }
        h1, h2 { text-align: center; }
        .card { background: rgba(255,255,255,0.15); padding: 20px; border-radius: 12px; margin: 20px 0; }
        input, button { padding: 12px; margin: 8px; font-size: 16px; border-radius: 8px; border: none; }
        button { background: #00d1b2; color: white; cursor: pointer; width: 100%; }
        button:hover { background: #00b89c; }
        pre { background: rgba(0,0,0,0.5); padding: 15px; border-radius: 8px; overflow-x: auto; }
        .hidden { display: none; }
        a { color: #ffd700; text-decoration: underline; }
    </style>
</head>
<body>
    <div class="container">
        <h1>Tienda Ecommerce SamAv</h1>
        <p style="text-align:center; font-size:18px;">Proyecto 100% funcional • Clean Architecture • JWT</p>

        <div class="card">
            <h2>Productos disponibles</h2>
            <p>Haz clic aquí para ver todos los productos:</p>
            <p style="text-align:center;">
                <a href="/api/products" target="_blank" style="font-size:20px; color:yellow; font-weight:bold;">
                    http://localhost:8080/api/products
                </a>
            </p>
        </div>

        <div class="card">
            <h2>Login (fácil)</h2>
            <p>Regístrate o inicia sesión. Si no existes, se crea automáticamente.</p>
            <input type="email" id="email" placeholder="tu@email.com" value="juan@test.com">
            <input type="password" id="pass" placeholder="contraseña" value="123456">
            <button onclick="login()">Iniciar sesión / Registrarse</button>
            <p id="msg" style="margin-top:10px; font-weight:bold;"></p>
        </div>

        <div class="card hidden" id="panel">
            <h2>¡Bienvenido! Ahora puedes comprar</h2>
            <p>Token guardado. Usa estos botones:</p>
            <button onclick="verCarrito()">Ver mi carrito</button>
            <button onclick="addToCart()">Añadir 2 camisetas al carrito</button>
            <button onclick="checkout()" style="background:#e74c3c; margin-top:10px;">Hacer checkout (pagar)</button>
            <pre id="result"></pre>
        </div>

        <div class="card">
            <h2>¿Quieres ver órdenes?</h2>
            <p><a href="/api/orders" target="_blank">http://localhost:8080/api/orders</a></p>
        </div>
    </div>

    <script>
        let token = "";

        async function login() {
            const email = document.getElementById("email").value;
            const pass = document.getElementById("pass").value;

            // Primero intenta login
            let res = await fetch("/api/auth/login", {
                method: "POST",
                headers: {"Content-Type": "application/json"},
                body: JSON.stringify({email, password: pass})
            });

            // Si falla, intenta registro
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
                // Volver a intentar login
                res = await fetch("/api/auth/login", {
                    method: "POST",
                    headers: {"Content-Type": "application/json"},
                    body: JSON.stringify({email, password: pass})
                });
            }

            const data = await res.json();
            token = data.token;
            document.getElementById("msg").innerText = "¡Login correcto! Token guardado";
            document.getElementById("panel").classList.remove("hidden");
        }

        async function verCarrito() {
            const res = await fetch("/api/cart", {headers: {"Authorization": "Bearer " + token}});
            const data = await res.json();
            document.getElementById("result").innerText = JSON.stringify(data, null, 2);
        }

        async function addToCart() {
            await fetch("/api/cart", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": "Bearer " + token
                },
                body: JSON.stringify({product_id: "prod_1", quantity: 2})
            });
            alert("Añadidas 2 camisetas!");
            verCarrito();
        }

        async function checkout() {
            const res = await fetch("/api/checkout", {
                method: "POST",
                headers: {"Authorization": "Bearer " + token}
            });
            const data = await res.json();
            document.getElementById("result").innerText = "¡COMPRA REALIZADA!\n\n" + JSON.stringify(data, null, 2);
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
