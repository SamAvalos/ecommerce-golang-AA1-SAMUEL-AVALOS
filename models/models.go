package models

import "time"

type Product struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	PriceCents int64  `json:"price_cents"`
	Stock      int    `json:"stock"`
	Active     bool   `json:"active"`
	CreatedAt  int64  `json:"created_at"`
}

type CartItem struct {
	ProductID  string `json:"product_id"`
	Qty        int    `json:"qty"`
	PriceCents int64  `json:"price_cents"`
}

type Cart struct {
	UserID string     `json:"user_id"`
	Items  []CartItem `json:"items"`
}

type OrderItem struct {
	ProductID  string `json:"product_id"`
	Qty        int    `json:"qty"`
	PriceCents int64  `json:"price_cents"`
}

type Order struct {
	ID         string      `json:"id"`
	UserID     string      `json:"user_id"`
	Items      []OrderItem `json:"items"`
	TotalCents int64       `json:"total_cents"`
	CreatedAt  time.Time   `json:"created_at"`
}
