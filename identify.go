package youdu

import (
	"context"
	"net/http"
)

type IdentifyStatus struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	CreatedAt string `json:"createdAt"`
}

type IdentifyUserInfo struct {
	Account string `json:"account"`
	ChsName string `json:"chsName"`
	Gender  int    `json:"gender"`
	Mobile  string `json:"mobile"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

type IdentifyResponse struct {
	Status   IdentifyStatus   `json:"status"`
	UserInfo IdentifyUserInfo `json:"userInfo"`
}

func (c *Client) Identify(ctx context.Context, token string) (response IdentifyResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/cgi/identify",
		withRequestParamsKV("token", token),
	)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}
