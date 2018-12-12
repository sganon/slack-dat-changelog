package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// App flags

var logLevelFlag = cli.StringFlag{
	Name:  "loglevel",
	Value: log.InfoLevel.String(),
	Usage: "Set your logging level (debug|info|warn|error)",
}

var debugFlag = cli.BoolFlag{
	Name:  "debug",
	Usage: "Set logging level to debug",
}

// Command  serve flags

var disableGitlabFlag = cli.BoolFlag{
	Name:  "disablegitlab",
	Usage: "Disable gitlab routes",
}

var serverHostFlag = cli.StringFlag{
	Name:  "host",
	Value: "0.0.0.0",
	Usage: "Set server host",
}

var serverPortFlag = cli.StringFlag{
	Name:  "port",
	Value: "8080",
	Usage: "Set server port",
}
