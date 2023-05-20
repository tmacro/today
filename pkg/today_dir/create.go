package today_dir

import (
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/pflag"
	"github.com/tmacro/today/pkg/config"
)

type CreateCommand struct {
	fs *pflag.FlagSet
}

func NewCreateCommand() *CreateCommand {
	fs := pflag.NewFlagSet("create", pflag.ContinueOnError)
	return &CreateCommand{
		fs: fs,
	}
}

func (c *CreateCommand) Name() string {
	return "create"
}

func (c *CreateCommand) Parse(args []string) error {
	return c.fs.Parse(args)
}

func (c *CreateCommand) Run(conf *config.TodayConfig) error {
	now := time.Now()
	date := now.Format(conf.DateFormat)
	dir := filepath.Join(conf.Notes.Directory, date)
	return os.MkdirAll(dir, 0755)
}
