package middleware

import (
	"context"
	"ecommerce/internal/auth"
	"net/http"
	"strings"
)

func AuthRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		userID, err := auth.ValidateToken(tokenStr)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		// Guardamos el userID en el contexto
		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
