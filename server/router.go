package server

import (
	"net/http"

	"github.com/anfimovoleh/ms-content-manager/server/handlers"

	"github.com/anfimovoleh/ms-content-manager/config"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"github.com/anfimovoleh/ms-content-manager/server/middlewares"

	"github.com/go-chi/chi"
)

func Router(cfg config.Config) chi.Router {
	router := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"*", "GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"*", "Accept", "Authorization", "Content-Type", "X-CSRF-Asset", "x-auth"},
		ExposedHeaders:   []string{"*", "Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	router.Use(
		middlewares.Logger(cfg.Log(), cfg.HTTP().ReqDurThreshold),
		c.Handler,
		middleware.Recoverer,
	)

	router.Mount("/admin", adminRouter(cfg))

	return router
}

// A completely separate router for administrator routes
func adminRouter(cfg config.Config) http.Handler {
	r := chi.NewRouter()
	r.Use(
		middlewares.VerifyRemoteAddressIsPrivate(middlewares.PrivateAddressPool()),
		middlewares.BasicAuth("admin", "admin"),
	)

	r.Route("/upload", func(r chi.Router) {
		r.Post("/file", handlers.NewUploadDocumentHandler(cfg.Log()).Handle)
	})

	return r
}
