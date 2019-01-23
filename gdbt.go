package main

import (
	"fmt"
	"os"

	"github.com/lepra-tsr/gdbt/handler/draft"
	"github.com/lepra-tsr/gdbt/handler/list"
	"github.com/lepra-tsr/gdbt/handler/post"
	"github.com/lepra-tsr/gdbt/handler/room"
	"github.com/lepra-tsr/gdbt/handler/setup"
	"github.com/urfave/cli"
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
			Aliases: []string{"setup", "i"},
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
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "reload, r"},
				cli.BoolFlag{Name: "show, s"},
			},
			Action: func(c *cli.Context) error {
				reloadFlag := c.Bool("reload")
				showFlag := c.Bool("show")
				if err := room.Handler(reloadFlag, showFlag); err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "show room messages.",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "union, u"},
			},
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
			Flags: []cli.Flag{
				cli.StringFlag{Name: "message, m"},
				cli.BoolFlag{Name: "draft, d"},
			},
			Action: func(c *cli.Context) error {
				messageOption := c.String("message")
				draftFlag := c.Bool("draft")

				if err := post.Handler(messageOption, draftFlag); err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "draft",
			Aliases: []string{"d"},
			Usage:   "write draft or post with.",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "show, s"},
				cli.BoolFlag{Name: "post, p"},
			},
			Action: func(c *cli.Context) error {
				showFlag := c.Bool("show")
				postFlag := c.Bool("post")

				if err := draft.Handler(showFlag, postFlag); err != nil {
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
