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

	app.Name = "gdbt"
	app.Usage = "idobata unofficial cli tool"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"setup", "i"},
			Usage:   "load your room information.",
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
			Usage:   "select room.",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "reload, r",
					Usage: "update room information.",
				},
				cli.BoolFlag{
					Name:  "show, s",
					Usage: "show room information.",
				},
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
			Usage:   "show timeline.",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name: "union, u",
					Usage:"show union(mixed) timeline",
				},
			},
			Action: func(c *cli.Context) error {
				unionFlag := c.Bool("union")
				if err := list.Handler(unionFlag); err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "post",
			Aliases: []string{"p"},
			Usage:   "post message.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "message, m",
					Usage: "post message.",
				},
				cli.BoolFlag{
					Name:  "draft, d",
					Usage: "post with draft file.",
				},
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
			Usage:   "write draft.",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "show, s",
					Usage: "show draft file.",
				},
				cli.BoolFlag{
					Name:  "post, p",
					Usage: "post with draft file.",
				},
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
