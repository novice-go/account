package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "db action"
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "action", Usage: "migrate, create, drop"},
	}
	app.Commands = []cli.Command{
		{Name: "create", Usage: "create db", Action: create},
		{Name: "drop", Usage: "drop db", Action: drop},
		{Name: "migrate", Usage: "migrate db", Action: migrate},
	}
	app.Action = cli.ShowAppHelp
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

func create (c *cli.Context) error {
	fmt.Println("create")
	return nil
}

func drop (c *cli.Context) error {
	fmt.Println("drop")
	return nil
}

func migrate (c *cli.Context) error {
	fmt.Println("migrate")
	return nil
}
