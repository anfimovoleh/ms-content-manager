package main

import (
	app "github.com/anfimovoleh/ms-content-manager"
	"github.com/anfimovoleh/ms-content-manager/config"
	"go.uber.org/zap"
)

func main() {
	var (
		cfg = config.New()
		api = app.New(cfg)
	)

	err := api.Start()
	if err != nil {
		cfg.Log().With(zap.Error(err)).
			Fatal("failed to start api")
	}
}
