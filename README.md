# Ecommerce API — In-Memory Store (Golang)

Una API de comercio electrónico desarrollada en **Go (Golang)** bajo una arquitectura modular clara y extensible.  
El sistema implementa productos, carrito de compras, órdenes y flujos completos de compra, siguiendo principios modernos de diseño de software y buenas prácticas del lenguaje Go.



## Características Principales

- Gestión de productos en memoria (sin base de datos).
- Carrito de compras por usuario.
- Proceso de checkout con validaciones reales.
- Registro de órdenes.
- Arquitectura modular:
  - `/models`
  - `/services`
  - `/routes`
- API REST completa y documentada.
- Implementación limpia siguiendo buenas prácticas de Go.



## Estructura del Proyecto
│── main.go

│
│── models/

│ ├── user.go

│ ├── product.go

│ └── cart.go

│
│── services/

│ ├── user_service.go

│ ├── product_service.go

│ └── cart_service.go

│
│── routes/

│ ├── user_routes.go

│ ├── product_routes.go

│ └── cart_routes.go

│
└── go.mod
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

## Endpoints Disponibles

### Usuarios
| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET    | `/api/users` | Listar usuarios |
| POST   | `/api/users` | Crear usuario |

---

### Productos
| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET    | `/api/products` | Listar productos |
| POST   | `/api/products` | Crear producto |
| GET    | `/api/products/{id}` | Obtener producto |
| DELETE | `/api/products/{id}` | Eliminar producto |

---

### Carrito
| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET    | `/api/cart/{userID}` | Ver carrito del usuario |
| POST   | `/api/cart/{userID}` | Agregar producto al carrito |

---

## Requisitos

- Go 1.21 o superior
- Postman / Thunder Client para pruebas
- Sistema operativo Windows, Linux o macOS

---

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


## Ejemplos de Requests

### Crear producto
POST http://localhost:8080/api/products

Body:
```json
{
  "id": "prod_1",
  "name": "Camiseta",
  "price": 25.50
}
```


### Agregar producto al carrito
POST http://localhost:8080/api/cart/user_1

### Ver productos

GET http://localhost:8080/api/products

## Extensiones Futuras

- Integración con base de datos real.

- Autenticación JWT.

- Módulo de pagos.

- Dashboard administrativo.

- Sistema de órdenes completo.

## Licencia
Este proyecto es de libre uso para fines educativos o profesionales. Puedes modificarlo y adaptarlo según tus necesidades.

