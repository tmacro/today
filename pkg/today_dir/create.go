package today_dir

import (
	"os"
	"path/filepath"
	"time"

	"github.com/tmacro/today/pkg/config"
	"github.com/tmacro/today/pkg/utils"
)

type CreateCommand struct {
	Date string `arg:"" optional:"" name:"date" help:"Date to show."`
}

func (c *CreateCommand) Run(conf *config.TodayConfig) error {
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

	date := t.Format(conf.DateFormat)
	dir := filepath.Join(conf.Notes.Directory, date)
	return os.MkdirAll(dir, 0755)
}
