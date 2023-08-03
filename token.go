package youdu

import (
	"context"
	"strconv"
	"time"
)

type token struct {
	Token   string
	Expired time.Time
}

func (c *Client) GetToken(ctx context.Context) (string, error) {
	if c.token != nil && c.token.Expired.After(time.Now()) {
		return c.token.Token, nil
	}

	ciphertext, err := c.encryptor.Encrypt([]byte(strconv.Itoa(int(time.Now().Unix()))))
	if err != nil {
		return "", err
	}

	resp, err := c.getToken(ctx, tokenRequest{
		Buin:    c.config.Buin,
		AppId:   c.config.AppId,
		Encrypt: ciphertext,
	})
	if err != nil {
		return "", err
	}

	c.token = &token{
		Token:   resp.AccessToken,
		Expired: time.Now().Add(time.Duration(resp.ExpireIn)*time.Second - 10*time.Minute), // 提前10分钟过期
	}

	return resp.AccessToken, nil
}

type tokenRequest struct {
	Buin    int    `json:"buin"`
	AppId   string `json:"appId"`
	Encrypt string `json:"encrypt"`
}

type tokenResponse struct {
	AccessToken string `json:"accessToken"`
	ExpireIn    int    `json:"expireIn"`
}

func (c *Client) getToken(ctx context.Context, request tokenRequest) (response tokenResponse, err error) {
	req, err := c.newRequest(ctx, "POST", "/cgi/gettoken", withRequestBody(request))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseDecrypt())
	return
}
