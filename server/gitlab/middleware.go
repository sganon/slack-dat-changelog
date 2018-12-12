package gitlab

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/sganon/slack-dat-changelog/common"
)

// Middleware handles token verification for gitlab
type Middleware struct {
	routes http.Handler
	tokens []string
}

// NewMiddleware returns a pointer to a new GitlabMiddleware
func NewMiddleware(routes http.Handler, tokens []string) *Middleware {
	return &Middleware{
		routes: routes,
		tokens: tokens,
	}
}

func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debugln("In Gitlab middleware")
	reqToken := r.Header.Get("X-Gitlab-Token")
	if reqToken == "" && len(m.tokens) > 0 {
		res := common.BaseError{Message: "no token provided"}
		common.JSONResponse(w, http.StatusForbidden, res)
		return
	}
	hasValidToken := false
	for _, token := range m.tokens {
		if token == reqToken {
			hasValidToken = true
			break
		}
	}
	if !hasValidToken {
		res := common.BaseError{Message: "no valid token found"}
		common.JSONResponse(w, http.StatusForbidden, res)
		return
	}
	m.routes.ServeHTTP(w, r)
}
