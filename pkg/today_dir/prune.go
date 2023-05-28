package today_dir

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/tmacro/today/pkg/cli"
	"github.com/tmacro/today/pkg/utils"
)

type PruneCommand struct {
	DryRun bool `help:"Don't actually delete anything."`
}

func (c *PruneCommand) Run(ctx *cli.Context) error {
	resolved, err := utils.ResolvePath(ctx.Config.Scratch.Directory)
	if err != nil {
		return err
	}
	files, err := ioutil.ReadDir(resolved)

	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() && isScratchDir(f.Name()) {
			path := filepath.Join(resolved, f.Name())
			empty, err := isEmpty(path)
			if err != nil {
				return err
			}
			if empty {
				if !c.DryRun {
					fmt.Println(path)
					err := os.Remove(path)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}
	return nil
}

func isEmpty(dir string) (bool, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return false, err
	}
	return len(files) == 0, nil
}

func isScratchDir(dir string) bool {
	_, err := utils.ParseFullYear(dir)
	return err == nil
}
