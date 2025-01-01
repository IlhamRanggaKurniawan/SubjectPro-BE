package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/utils"
)

type Middleware func(http.Handler) http.Handler

func CreateStack(md ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(md) - 1; i >= 0; i-- {
			next = md[i](next)
		}
		return next
	}
}

func AuthMiddelware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		unprotectedRoutes := []string{
			"/v1/auth",
			"/v1/auth/login",
			"/v1/auth/token",
		}

		for _, route := range unprotectedRoutes {
			if r.URL.Path == route {
				next.ServeHTTP(w, r)
				return
			}
		}

		token := r.Header.Get("Authorization")

		if token == "" {
			utils.ErrorResponse(w, fmt.Errorf("token not found"), http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(token, "Bearer") {
			utils.ErrorResponse(w, fmt.Errorf("invalid token type"), http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(token, "Bearer ")

		err := utils.ValidateAccessToken(tokenString)

		if err != nil {
			utils.ErrorResponse(w, err, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		frontEndOrigin := os.Getenv("FRONT_END_ORIGIN")

		allowedOrigins := []string{
			frontEndOrigin,
			"",
		}

		origin := r.Header.Get("Origin")

		allowed := false

		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				w.Header().Set("Content-Type", "application/json")
				allowed = true
				break
			}
		}

		if !allowed {
			utils.ErrorResponse(w, fmt.Errorf("origin not allowed"), http.StatusForbidden)
			return
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
