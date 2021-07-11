package server

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"github.com/anfimovoleh/ms-content-manager/server/middlewares"

	"github.com/go-chi/chi"
)

func Router() chi.Router {
	router := chi.NewRouter()
	privateAddressPool := middlewares.PrivateAddressPool()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"*", "GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"*", "Accept", "Authorization", "Content-Type", "X-CSRF-Asset", "x-auth"},
		ExposedHeaders:   []string{"*", "Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	router.Use(
		c.Handler,
		middleware.Recoverer,
	)

	router.Group(func(r chi.Router) {
		r.Use(middlewares.VerifyRemoteAddressIsPrivate(privateAddressPool))
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write(
				[]byte(`{"component": "ms-content-manager"}`),
			)
		})
	})

	return router
}
