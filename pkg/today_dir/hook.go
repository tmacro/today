package today_dir

import "github.com/tmacro/today/pkg/cli"

type HookCommand struct {
	Shell string `help:"Shell to install hook for."`
}

func (c *HookCommand) Run(ctx *cli.Context) error {
	return nil
}
