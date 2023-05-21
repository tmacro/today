package cli

import (
	"github.com/sirupsen/logrus"
	"github.com/tmacro/today/pkg/config"
)

type Context struct {
	Verbose bool
	Config  *config.TodayConfig
	Logger  *logrus.Logger
}

func NewContext(verbose bool, conf *config.TodayConfig) *Context {
	logger := logrus.New()
	if verbose {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}
	return &Context{
		Verbose: verbose,
		Config:  conf,
		Logger:  logger,
	}
}
