package today_dir

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/tmacro/today/pkg/cli"
	"github.com/tmacro/today/pkg/utils"
)

type ShowCommand struct {
	Date string `arg:"" optional:"" name:"date" help:"Date to show."`
}

func (c *ShowCommand) Run(ctx *cli.Context) error {
	var t time.Time
	if c.Date == "" {
		t = time.Now()
	} else {
		var err error
		t, err = utils.ParseDate(c.Date)
		if err != nil {
			return err
		}
	}

	date := t.Format(ctx.Config.DateFormat)
	dir := filepath.Join(ctx.Config.Scratch.Directory, date)
	resolved, err := utils.ResolvePath(dir)
	if err != nil {
		return err
	}

	fmt.Println(resolved)
	return nil
}
