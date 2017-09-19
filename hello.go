package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "project finder"
	app.Usage = "create a simple cli to find and get to all my projects"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello brave new world!")
		return nil
	}

	app.Run(os.Args)
}
