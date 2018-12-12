package gitlab

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"

	"github.com/sganon/slack-dat-changelog/common"
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
	var body common.GitlabTagBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		log.Error(fmt.Errorf("error in handleWebHook: %v", err))
		common.JSONResponse(w, http.StatusInternalServerError, common.BaseError{
			Message: "unexpected error decoding request body",
		})
	}
	logFields := log.Fields{
		"project_id":   body.ProjectID,
		"ref":          body.Ref,
		"checkout_sha": body.CheckoutSHA,
		"user_name":    body.UserName,
	}
	log.WithFields(logFields).Debugln("Processing tag hook")
}
