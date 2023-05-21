package today_dir

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
	"time"

	"github.com/tmacro/today/pkg/cli"
	"github.com/tmacro/today/pkg/utils"
)

type ViewCommand struct {
	Date string `arg:"" optional:"" name:"date" help:"Date to show."`
}

type ViewTemplate struct {
	Directory string
	Date      string
}

func (c *ViewCommand) Run(ctx *cli.Context) error {
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

	tmpl, err := template.New("view").Parse(ctx.Config.Scratch.Viewer)
	if err != nil {
		return err
	}

	env := ViewTemplate{
		Directory: resolved,
		Date:      date,
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, env)
	if err != nil {
		return err
	}

	cmd := buf.String()
	ctx.Logger.Debug(cmd)
	proc := exec.Command("sh", "-c", cmd)
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stderr
	return proc.Run()
}
