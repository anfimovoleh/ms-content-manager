package middlewares

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
)

func BasicAuth(username, password string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			u, p, ok := r.BasicAuth()
			if !ok {
				fmt.Println("Error parsing basic auth")
				w.WriteHeader(401)
				return
			}
			if u != username {
				fmt.Printf("Username provided is correct: %s\n", u)
				w.WriteHeader(401)
				return
			}
			if p != password {
				fmt.Printf("Password provided is correct: %s\n", u)
				w.WriteHeader(401)
				return
			}

			next.ServeHTTP(ww, r)
		})
	}
}
