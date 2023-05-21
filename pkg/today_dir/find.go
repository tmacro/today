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

type FindCommand struct {
	Expression string `arg:"" name:"expression" help:"Regex expression to search for."`
	Date       string `arg:"" optional:"" name:"date" help:"Date to show."`
}

type FindTemplate struct {
	Directory  string
	Expression string
}

func (c *FindCommand) Run(ctx *cli.Context) error {

	var dir string
	if c.Date != "all" {
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
		dir = filepath.Join(ctx.Config.Scratch.Directory, date)
	} else {
		dir = ctx.Config.Scratch.Directory
	}

	resolved, err := utils.ResolvePath(dir)
	if err != nil {
		return err
	}

	tmpl, err := template.New("find").Parse(ctx.Config.Scratch.Find)
	if err != nil {
		return err
	}

	env := FindTemplate{
		Directory:  resolved,
		Expression: c.Expression,
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
