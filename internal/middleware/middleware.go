package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/utils"
)

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

		next.ServeHTTP(w,r)
	})
}
