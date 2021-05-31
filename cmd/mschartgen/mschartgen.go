package main

import (
	"github.com/flaccid/mschartgen"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

var (
	orgName string
	rootMemberId string
)

func main() {
	app := &cli.App{
		Name:  "mschartgen",
		Usage: "generates an organisational chart from the microsoft graph api",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "log-level",
				Aliases: []string{"l"},
				Value:   "info",
				Usage:   "log level (debug|info|warn|error)",
			},
			&cli.StringFlag{
				Name:    "org-name",
				Aliases: []string{"o"},
				Usage:   "overrides the organisstion name",
			},
		},
		Before: func(c *cli.Context) error {
			if len(c.Args().Get(0)) > 0 {
				rootMemberId = c.Args().Get(0)
			}
			if len(os.Getenv("ROOT_MEMBER_ID")) > 1 {
				rootMemberId = os.Getenv("ROOT_MEMBER_ID")
			}
			if len(rootMemberId) < 1 {
				log.Fatalf("please provide the id of the root org member")
			}
			if len(c.String("org-name")) > 0 {
				orgName = c.String("org-name")
			}

			//log.SetOutput(os.Stdout)
			switch c.String("log-level") {
			case "debug":
				log.SetLevel(log.DebugLevel)
			case "info":
				log.SetLevel(log.InfoLevel)
			case "warn":
				log.SetLevel(log.WarnLevel)
			case "error":
				log.SetLevel(log.ErrorLevel)
			}
			return nil
		},
		Action: func(c *cli.Context) error {
			log.Info("starting mschartgen")
			mschartgen.Process(orgName, rootMemberId)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
