package cli

import "github.com/tmacro/today/pkg/config"

type SubCommand interface {
	Name() string
	Parse(args []string) error
	Run(*config.TodayConfig) error
}
