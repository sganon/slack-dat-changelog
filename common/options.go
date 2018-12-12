package common

import (
	log "github.com/sirupsen/logrus"
)

// ServerOptions describe available options to init server
type ServerOptions struct {
	EnableGitlab bool
	Host         string
	Port         string
}

func (so ServerOptions) log(level log.Level) {
	logFields := log.Fields{
		"GitlabEnabled": so.EnableGitlab,
	}
	const msg = "server options"
	switch level {
	case log.DebugLevel:
		log.WithFields(logFields).Debugln(msg)
	case log.InfoLevel:
		log.WithFields(logFields).Infoln(msg)
	}
}

// Info log server options with info level
func (so ServerOptions) Info() {
	so.log(log.InfoLevel)
}

// Debug log server options with debug level
func (so ServerOptions) Debug() {
	so.log(log.DebugLevel)
}
