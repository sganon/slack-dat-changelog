package gitlab

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// RoutePrefix defines the prefix of all gitlab routes
const RoutePrefix = "/gitlab"

// Routes defines gitlab routes
func Routes() http.Handler {
	router := httprouter.New()
	router.POST(RoutePrefix+"/", handleWebHook)
	return router
}

func handleWebHook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Debugln("In Gitlab handler")

}
