package app

import (
	"fmt"
	"net/http"

	"github.com/anfimovoleh/ms-content-manager/config"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type API struct {
	cfg config.Config
}

func New(cfg config.Config) *API {
	return &API{
		cfg: cfg,
	}
}

func (a API) Start() error {
	var (
		log               = a.cfg.Log()
		httpConfiguration = a.cfg.HTTP()
	)

	serverHost := fmt.Sprintf("%s:%s", httpConfiguration.Host, httpConfiguration.Port)
	log.With(
		zap.String("addr", serverHost),
	).Info(fmt.Sprintf("starting api"))

	httpServer := http.Server{
		Addr:    serverHost,
		Handler: nil,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		return errors.Wrap(err, "failed to start http server")
	}

	return nil
}
