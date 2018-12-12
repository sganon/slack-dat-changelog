package gitlab

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"

	"github.com/sganon/slack-dat-changelog/common"
	"github.com/sganon/slack-dat-changelog/slack"
)

// RoutePrefix defines the prefix of all gitlab routes
const RoutePrefix = "/gitlab"

type handler struct {
	slack *slack.Client
}

// Routes defines gitlab routes
func Routes(slackURI string) http.Handler {
	client := slack.New(slackURI, "tet-hooks")
	h := handler{slack: client}
	router := httprouter.New()
	router.POST(RoutePrefix+"/", h.handleWebHook)
	return router
}

func (h handler) handleWebHook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
	h.slack.SendMessage(slack.Payload{
		Attachments: []slack.Attachment{
			{
				Fallback: fmt.Sprintf("New release of project %d", body.ProjectID),
				Pretext:  fmt.Sprintf("New release of project %d", body.ProjectID),
				Color:    "good",
				Fields: []slack.Field{
					{
						Title: "Added",
						Value: "First changelog",
					},
				},
			},
		},
	})
}
