package config

import "go.uber.org/zap"

type Config interface {
	HTTP() *HTTP
	Log() *zap.Logger
}

type Impl struct {
	// internal objects
	http *HTTP
	log  *zap.Logger
}

func New() Config {
	return &Impl{}
}
