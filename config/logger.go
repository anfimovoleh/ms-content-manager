package config

import (
	"log"

	"github.com/caarlos0/env"
	"go.uber.org/zap"
)

// Log core structure responsible for logging in the application
type Log struct {
	Level      string `env:"CM_LOG_LEVEL" envDefault:"DEBUG"`
	EnableJSON bool   `env:"CM_LOG_ENABLE_JSON" envDefault:"false"`
}

func (c *Impl) Log() *zap.Logger {
	if c.log != nil {
		return c.log
	}

	var (
		l       Log
		err     error
		zCfg    = zap.NewProductionConfig()
		zLogger *zap.Logger
	)

	if err = env.Parse(&l); err != nil {
		log.Fatalf("failed to parse log object: %v \n", err)
	}

	switch l.Level {
	case zap.DebugLevel.CapitalString():
		zCfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case zap.ErrorLevel.CapitalString():
		zCfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case zap.InfoLevel.CapitalString():
		zCfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	default:
		log.Fatalf("unrecognized log level: %s; Available options: DEBUG, ERROR, INFO\n", l.Level)
	}

	switch l.EnableJSON {
	case true:
		zCfg.Encoding = "json"
	case false:
		zCfg.Encoding = "console"
	}

	zLogger, err = zCfg.Build()
	if err != nil {
		log.Fatalf("failed to build logger config: %v \n", err)
	}

	c.log = zLogger

	c.log.Info("initialized log configuration",
		zap.String("level", zCfg.Level.String()),
		zap.String("encoding", zCfg.Encoding),
	)

	return c.log
}
