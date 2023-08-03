package youdu

import (
	"context"
	"net/http"
)

type SessionType string

var (
	SessionMultiType SessionType = "multi"
)

type CreateSessionRequest struct {
	Title   string      `json:"title"`
	Creator string      `json:"creator"`
	Type    SessionType `json:"type,omitempty"`
	Member  []string    `json:"member"`
}

type UpdateSessionRequest struct {
	SessionId string   `json:"sessionId"`
	OpUser    string   `json:"opUser"`
	Title     string   `json:"title,omitempty"`
	AddMember []string `json:"addMember"`
	DelMember []string `json:"delMember"`
}

type SessionResponse struct {
	SessionId string      `json:"sessionId"`
	Title     string      `json:"title"`
	Owner     string      `json:"owner"`
	Version   int         `json:"version"`
	Type      SessionType `json:"type"`
	Member    []string    `json:"member"`
}

func (c *Client) CreateSession(ctx context.Context, request CreateSessionRequest) (response SessionResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/cgi/session/create",
		withRequestBody(request), withRequestAccessToken(), withRequestEncrypt())
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseDecrypt())
	return
}

func (c *Client) GetSession(ctx context.Context, sessionId string) (response SessionResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/cgi/session/get",
		withRequestAccessToken(),
		withRequestEncrypt(),
		withRequestParamsKV("sessionId", sessionId),
	)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseDecrypt())
	return
}

func (c *Client) UpdateSession(ctx context.Context, request UpdateSessionRequest) (response SessionResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/cgi/session/update",
		withRequestBody(request), withRequestAccessToken(), withRequestEncrypt())
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseDecrypt())
	return
}
