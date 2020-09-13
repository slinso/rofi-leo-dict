package main

import (
	"errors"
	"log"
	"os"

	rld "github.com/seambiz/rofi-leo-dict"
	"github.com/urfave/cli/v2"
)

func main() {
	app :=
		cli.App{
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "from",
					Usage:       "dict language",
					Value:       "en",
					DefaultText: "english",
				},
				&cli.StringFlag{
					Name:        "to",
					Usage:       "dict language",
					Value:       "de",
					DefaultText: "german",
				},
			},
			Action: func(c *cli.Context) error {
				if c.Args().Len() != 1 {
					return errors.New("no argument passed")
				}

				data, err := rld.ScrapeLeo(c.String("from"), c.String("to"), c.Args().First())
				if err != nil {
					panic(err)
				}

				rld.CreateTable(c.String("from"), c.String("to"), data)

				return nil
			},
		}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
