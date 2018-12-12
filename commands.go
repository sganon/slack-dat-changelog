package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/sganon/slack-dat-changelog/common"
	"github.com/sganon/slack-dat-changelog/server"
)

var serveCommand = cli.Command{
	Name:    "serve",
	Aliases: []string{"serve"},
	Usage:   "Launch server",
	Flags: []cli.Flag{
		disableGitlabFlag,
		serverHostFlag, serverPortFlag,
	},

	Action: func(c *cli.Context) (err error) {
		log.Debugln("Starting server..")
		so := common.ServerOptions{
			EnableGitlab: !c.Bool(disableGitlabFlag.GetName()),
			Host:         c.String(serverHostFlag.GetName()),
			Port:         c.String(serverPortFlag.GetName()),
		}
		so.Info()
		server := server.New(so)
		server.Run()
		return err
	},
}
