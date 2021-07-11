package server

import (
	"net/http"

	"github.com/anfimovoleh/ms-content-manager/server/middlewares"

	"github.com/go-chi/chi"
)

func Router() chi.Router {
	router := chi.NewRouter()

	privateAddressPool := middlewares.PrivateAddressPool()

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
