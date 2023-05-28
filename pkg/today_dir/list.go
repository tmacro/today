package today_dir

import (
	"fmt"
	"io/ioutil"

	"github.com/tmacro/today/pkg/cli"
	"github.com/tmacro/today/pkg/utils"
)

type ListCommand struct {
}

func (c *ListCommand) Run(ctx *cli.Context) error {
	resolved, err := utils.ResolvePath(ctx.Config.Scratch.Directory)
	if err != nil {
		return err
	}
	files, err := ioutil.ReadDir(resolved)

	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() {
			fmt.Println(f.Name())
		}
	}
	return nil
}
