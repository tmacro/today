package main

import (
	"github.com/alecthomas/kong"
	"github.com/tmacro/today/pkg/cli"
	"github.com/tmacro/today/pkg/config"
	"github.com/tmacro/today/pkg/today_dir"
)

var _cli struct {
	Verbose bool `flag:"verbose" help:"verbose output"`
	Dir     struct {
		Create today_dir.CreateCommand `cmd:"" help:"Create today's directory."`
		Show   today_dir.ShowCommand   `cmd:"" help:"Print the path to today's directory."`
		View   today_dir.ViewCommand   `cmd:"" help:"View today's directory."`
		Search today_dir.SearchCommand `cmd:"" help:"Search files in today's directory."`
		Find   today_dir.FindCommand   `cmd:"" help:"Find files in today's directory."`
		List   today_dir.ListCommand   `cmd:"" help:"List all scratch directories."`
		Prune  today_dir.PruneCommand  `cmd:"" help:"Delete empty scratch directories."`
		Hook   today_dir.HookCommand   `cmd:"" help:"Install shell hook."`
	} `cmd:"" help:"Manage today's directory."`
}

func main() {
	ctx := kong.Parse(&_cli)

	conf, err := config.ReadConfig()
	if err != nil {
		ctx.FatalIfErrorf(err)
	}

	err = ctx.Run(cli.NewContext(_cli.Verbose, conf))
	ctx.FatalIfErrorf(err)
}
