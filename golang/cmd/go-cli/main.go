package main

import (
	//"fmt"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

// 查看指令说明： cli -h
// 查看具体指令说明: cli test -h

func main() {
	app := &cli.App{
		Name:  "cli",
		Usage: "make an explosive entrace",
		Commands: []*cli.Command{ // 有Commands就不用action了
			{
				Name: "test",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "bool", // cli test -b 1
						Aliases: []string{"b"},
						Usage:   "is it happy?",
					},
					&cli.StringFlag{ // cli test -s happy
						Name:    "string",
						Aliases: []string{"s"},
						Usage:   "say hello!",
					},
				},
				Action: func(c *cli.Context) error {
					if c.Bool("bool") {
						log.Println("command: test, flag: bool")
						return nil
					}
					if s := c.String("string"); s != "" {
						log.Println("command: test, flag: ", s)
					}
					return nil
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
