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

var slackURIFlag = cli.StringFlag{
	Name:   "slackURI",
	Usage:  "Set slack webhook uri",
	EnvVar: "SLDC_SLACK_URI",
}

// Command  serve flags

var disableGitlabFlag = cli.BoolFlag{
	Name:  "disablegitlab",
	Usage: "Disable gitlab routes",
}

var gitlabAccessTokenFlag = cli.StringFlag{
	Name:   "gitlabtoken",
	Usage:  "Gitlab personal access token",
	EnvVar: "SLDC_GITLAB_TOKEN",
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

var gitlabTokensFlag = cli.StringSliceFlag{
	Name:   "gitlabtokens",
	Usage:  "Pass wb hooks token defined in gitlab",
	EnvVar: "SLDC_GITLAB_TOKENS",
}
