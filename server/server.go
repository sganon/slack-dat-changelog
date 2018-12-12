package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"

	"github.com/sganon/slack-dat-changelog/common"
	"github.com/sganon/slack-dat-changelog/server/gitlab"
)

// Server handles routing
type Server struct {
	handler http.Handler
	addr    string
}

// New creates and returns a pointer to a new server
func New(so common.ServerOptions) *Server {
	var handler http.Handler
	var gitlabRouter http.Handler

	if so.EnableGitlab {
		gitlabRouter = gitlab.Routes(so.SlackURI)
		gitlabRouter = gitlab.NewMiddleware(gitlabRouter, so.GitlabTokens)
	}

	handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if strings.HasPrefix(path, gitlab.RoutePrefix) {
			log.Debugln("Gitlab route detected")
			gitlabRouter.ServeHTTP(w, r)
		}
	})

	handler = handlers.CustomLoggingHandler(os.Stdout, handler, logFormatter)

	return &Server{
		handler: handler,
		addr:    fmt.Sprintf("%s:%s", so.Host, so.Port),
	}
}

// Run launch server
func (s Server) Run() {
	log.WithFields(log.Fields{
		"addr": s.addr,
	}).Infoln("Listening on")
	if err := http.ListenAndServe(s.addr, handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(s.handler)); err != nil {
		log.Fatalln(err)
	}
}

func logFormatter(w io.Writer, p handlers.LogFormatterParams) {
	if p.StatusCode != http.StatusTemporaryRedirect {
		var requestLogger = log.New()
		requestLogger.Out = w
		logFields := log.Fields{
			"date":   time.Now(),
			"method": p.Request.Method,
			"path":   p.Request.URL.Path,
			"http_v": fmt.Sprintf("http %d.%d", p.Request.ProtoMajor, p.Request.ProtoMinor),
			"status": p.StatusCode,
			"agent":  p.Request.Header.Get("User-Agent"),
		}
		if p.StatusCode >= http.StatusBadRequest && p.StatusCode < http.StatusInternalServerError {
			requestLogger.WithFields(logFields).Warnln("")
		} else if p.StatusCode >= http.StatusInternalServerError {
			requestLogger.WithFields(logFields).Errorln("")
		} else {
			requestLogger.WithFields(logFields).Infoln("")
		}
	}
}
