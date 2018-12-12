package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "slack-dat-changelog"
	app.Usage = "Server designed to receive webhooks from Gitlab in order to send changelog to Slack"
	app.Flags = []cli.Flag{
		logLevelFlag, debugFlag,
	}

	app.Before = func(c *cli.Context) (err error) {
		debugEnabled := c.Bool(debugFlag.GetName())
		rawLogLevel := c.String(logLevelFlag.GetName())
		if debugEnabled {
			log.SetLevel(log.DebugLevel)
		} else {
			logLevel, err := log.ParseLevel(rawLogLevel)
			if err != nil {
				return err
			}
			log.SetLevel(logLevel)
		}
		return err
	}

	app.Commands = []cli.Command{
		serveCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
