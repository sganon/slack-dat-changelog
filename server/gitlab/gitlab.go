package gitlab

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"

	"github.com/sganon/slack-dat-changelog/common"
	"github.com/sganon/slack-dat-changelog/gitlab"
	"github.com/sganon/slack-dat-changelog/slack"
)

// RoutePrefix defines the prefix of all gitlab routes
const RoutePrefix = "/gitlab"

type handler struct {
	slack  *slack.Client
	gitlab *gitlab.Client
}

// Routes defines gitlab routes
func Routes(slackURI, gitlabAccessToken string) http.Handler {
	slackClient := slack.New(slackURI, "tet-hooks")
	gitlabClient := gitlab.New(gitlabAccessToken)
	h := handler{slack: slackClient, gitlab: gitlabClient}
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

	changelog, err := h.gitlab.GetRawFile(body.Project.WebURL, "CHANGELOG.md")
	if err != nil {
		log.Error(fmt.Errorf("error in handleWebHook: %v", err))
		common.JSONResponse(w, http.StatusInternalServerError, common.BaseError{
			Message: "an unexpected error occured fetching changelog",
		})
	}
	// fmt.Println(string(changelog))

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
