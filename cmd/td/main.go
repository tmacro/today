package main

import (
	"fmt"
	"os"

	"github.com/tmacro/today/pkg/cli"
	"github.com/tmacro/today/pkg/config"
	"github.com/tmacro/today/pkg/today_dir"
)

var cmds = []cli.SubCommand{
	today_dir.NewCreateCommand(),
	today_dir.NewShowCommand(),
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("You must pass a sub-command")
		os.Exit(1)
	}

	name := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == name {
			err := runCommand(cmd, args[1:])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			os.Exit(0)
		}
	}

	fmt.Printf("Unknown sub-command: %s\n", name)
	os.Exit(1)
}

func runCommand(cmd cli.SubCommand, args []string) error {
	conf, err := config.ReadConfig()
	if err != nil {
		return err
	}

	err = cmd.Parse(args)
	if err != nil {
		return err
	}

	return cmd.Run(conf)
}
