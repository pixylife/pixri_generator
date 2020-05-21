package model

type GenResponse struct {
	GithubURL string ` json:"github"`
	WSURL string ` json:"websocket"`
	ClientRequestId string ` json:"client_id"`
}
