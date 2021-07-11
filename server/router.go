package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func Router() chi.Router {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(
			[]byte(fmt.Sprintf(`{"component": "ms-content-manager"}`)),
		)
	})

	return router
}
