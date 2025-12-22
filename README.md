# Ecommerce API — In-Memory Store (Golang)

Una API de comercio electrónico desarrollada en **Go (Golang)** bajo una arquitectura modular clara y extensible.  
El sistema implementa productos, carrito de compras, órdenes y flujos completos de compra, siguiendo principios modernos de diseño de software y buenas prácticas del lenguaje Go.



## Características Principales

- **Arquitectura limpia** (Clean Architecture) con separación de capas
- **Autenticación JWT** (HS256) con middleware protector
- **Carrito por usuario** (map[string]Cart + RWMutex para concurrencia)
- **Checkout con reserva de stock** y rollback si falla
- **Interfaz web completa** (HTML + JS puro) sin necesidad de terminal
- **Mensaje de sostenibilidad** después de cada compra
- **Serialización/deserialización 100% con JSON**
- **Encapsulación estricta** (campos privados, paquetes internal)
- **Interfaces** para repositorios (fácil migrar a DB)
- **Manejo de errores profesional** con rollback
- **Comentarios detallados** en código complejo



## Estructura del Proyecto
ecommerce/
├── main.go
├── go.mod
├── models/
│   └── models.go
├── internal/
│   ├── repository/
│   │   ├── interfaces.go
│   │   └── inmemory.go
│   ├── service/
│   │   ├── product.go
│   │   ├── cart.go
│   │   ├── order.go
│   │   └── checkout.go
│   ├── controller/
│   │   ├── product.go
│   │   ├── cart.go
│   │   ├── order.go
│   │   └── checkout.go
│   ├── middleware/
│   │   └── auth.go
│   └── auth/
│       └── service.go
└── routes/
└── routes.go
## Descripción por Módulos

### 1. Models  
Definen las estructuras principales del sistema: usuarios, productos y carritos.  
Sirven como base para los servicios y representan los datos del negocio.

### 2. Services  
Implementan toda la lógica del negocio:
- Validaciones
- Reglas del sistema
- Manejo de datos en memoria
- Control del flujo de compra

Los controladores se apoyan totalmente en esta capa.

### 3. Routes  
Exponen los endpoints HTTP y conectan las solicitudes con sus respectivos servicios.  
Organizan la API y mantienen separado el manejo HTTP de la lógica del negocio.

### 4. main.go  
Punto de entrada del programa.  
Inicializa:
- Servicios
- Semilla de productos
- Rutas  
Y finalmente levanta el servidor en `:8080`.

---


## Endpoints API

| Método | Ruta                          | Descripción                                   | Auth |
|--------|-------------------------------|-----------------------------------------------|------|
| GET    | `/`                           | Página web principal con login y tienda       | No   |
| GET    | `/api/products`               | Listar productos                              | No   |
| POST   | `/api/auth/register`          | Registrar usuario                             | No   |
| POST   | `/api/auth/login`             | Login y obtener JWT                           | No   |
| GET    | `/api/cart`                   | Ver carrito del usuario                       | Sí   |
| POST   | `/api/cart`                   | Añadir producto al carrito                    | Sí   |
| POST   | `/api/checkout`               | Realizar pago y generar orden                 | Sí   |
| GET    | `/api/orders`                 | Listar órdenes del usuario                    | Sí   |
| GET    | `/health`                     | Health check                                  | No   |

## Requisitos Técnicos

- Go 1.21 o superior
- Dependencias: `github.com/go-chi/chi/v5`, `github.com/golang-jwt/jwt/v5`, `golang.org/x/crypto/bcrypt`

## Instalación y Ejecución

### 1. Clonar repositorio
git clone https://github.com/tu-usuario/ecommerce-go-api.git

cd ecommerce-go-api

### 2. Instalar dependencias
go mod tidy

### 3. Ejecutar el servidor
go run main.go

Servidor disponible en:
http://localhost:8080

### Visión del Futuro (2030-2035)

IA personalizada (recomendaciones con Grok/xAI)
Pagos Web3 (cripto y stablecoins)
Realidad Aumentada (prueba virtual de ropa)
Sostenibilidad integrada (compensación CO2 por compra)
Blockchain (trazabilidad total)
Edge Computing + WASM (latencia cero)

## Licencia
Este proyecto es de libre uso para fines educativos o profesionales. Puedes modificarlo y adaptarlo según tus necesidades.

