package common

type UserID int

type ProjectID int

// GitlabTagBody correspond to the request body of a tag push in gitlab
type GitlabTagBody struct {
	ObjectKind  string    `json:"object_kind"`
	Before      string    `json:"before"`
	After       string    `json:"after"`
	Ref         string    `json:"ref"`
	CheckoutSHA string    `json:"checkout_sha"`
	UserID      UserID    `json:"user_id"`
	UserName    string    `json:"user_name"`
	UserAvatar  string    `json:"user_avatar"`
	ProjectID   ProjectID `json:"project_id"`
	Project     struct {
		ID         ProjectID `json:"id"`
		Name       string    `json:"name"`
		WebURL     string    `json:"web_url"`
		AvatarURL  string    `json:"avatar_url"`
		GitSSHURL  string    `json:"git_ssh_url"`
		GitHTTPURL string    `json:"git_http_url"`
		Namespace  string    `json:"namespace"`
	} `json:"project"`
	Repository struct {
		Name        string `json:"string"`
		URL         string `json:"url"`
		Description string `json:"description"`
	} `json:"repository"`
}
