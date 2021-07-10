package config

import (
	"fmt"
	"net/url"

	"go.uber.org/zap"

	"github.com/pkg/errors"

	"github.com/caarlos0/env"
)

type HTTP struct {
	Host string `env:"CM_HTTP_HOST" envDefault:"localhost"`
	Port string `env:"CM_HTTP_PORT" envDefault:"8080"`
}

func (h *HTTP) URL() (*url.URL, error) {
	resultURL, err := url.Parse(
		fmt.Sprintf("http://%s:%s", h.Host, h.Port),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse url")
	}

	return resultURL, nil
}

func (c *Impl) HTTP() *HTTP {
	if c.http != nil {
		return c.http
	}

	http := &HTTP{}
	if err := env.Parse(http); err != nil {
		c.Log().With(zap.Error(err)).
			Fatal("failed to initialize http configuration")
	}

	c.http = http

	c.Log().Info("initialized http configuration",
		zap.String("host", c.http.Host),
		zap.String("port", c.http.Port),
	)

	return c.http
}
