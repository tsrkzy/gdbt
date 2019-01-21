package main

import (
	"fmt"
	"github.com/lepra-tsr/gdbt/handler/list"
	"github.com/lepra-tsr/gdbt/handler/room"
	"github.com/lepra-tsr/gdbt/handler/post"
	"github.com/lepra-tsr/gdbt/handler/setup"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "gdbtApp"
	app.Usage = "idobata unofficial cli tool"
	app.Version = "0.0.1"

	// app.Flags = []cli.Flag{}
	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"setup"},
			Usage:   "load and store your organization and room information to ~/.gdbt",
			Action: func(c *cli.Context) error {
				if err := setup.Handler(); err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "room",
			Aliases: []string{"r"},
			Usage:   "choose current room to where you read/post.",
			Action: func(c *cli.Context) error {
				if err := room.Handler(); err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "show room messages.",
			Action: func(c *cli.Context) error {
				if err := list.Handler(); err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "post",
			Aliases: []string{"p"},
			Usage:   "post message to room.",
			Action: func(c *cli.Context) error {
				if err := post.Handler(c.Args().First()); err != nil {
					return err
				}
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
