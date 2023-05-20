package today_dir

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/spf13/pflag"
	"github.com/tmacro/today/pkg/config"
	"github.com/tmacro/today/pkg/utils"
)

type ShowCommand struct {
	fs   *pflag.FlagSet
	date string
}

func NewShowCommand() *ShowCommand {
	fs := pflag.NewFlagSet("show", pflag.ContinueOnError)
	return &ShowCommand{
		fs: fs,
	}
}

func (c *ShowCommand) Name() string {
	return "show"
}

func (c *ShowCommand) Parse(args []string) error {
	if len(args) > 0 {
		// try to parse all args as flags first
		err := c.fs.Parse(args)
		if err == nil {
			if c.fs.NArg() > 1 {
				return fmt.Errorf("too many arguments")
			}
			if c.fs.NArg() > 0 {
				c.date = c.fs.Arg(0)
			}
			return nil
		}

		// Try again without the last arg as it might be a negative number
		err = c.fs.Parse(args[:len(args)-1])
		if err != nil {
			return err
		}

		// last arg is date
		c.date = args[len(args)-1]
	}
	return nil
}

func (c *ShowCommand) Run(conf *config.TodayConfig) error {
	var t time.Time
	if c.date == "" {
		t = time.Now()
	} else {
		var err error
		t, err = utils.ParseDate(c.date)
		if err != nil {
			return err
		}
	}

	date := t.Format(conf.DateFormat)
	dir := filepath.Join(conf.Scratch.Directory, date)
	resolved, err := utils.ResolvePath(dir)
	if err != nil {
		return err
	}

	fmt.Println(resolved)
	return nil
}
