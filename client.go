package youdu

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type Config struct {
	Addr   string
	Buin   int
	AppId  string
	AesKey string
}

type Client struct {
	config    *Config
	encryptor *Encryptor
	token     *token
}

func NewClient(config *Config) *Client {
	return &Client{
		config:    config,
		encryptor: NewEncryptorWithConfig(config),
	}
}

func (c *Client) newRequest(ctx context.Context, method string, path string, opts ...requestOption) (*http.Request, error) {
	opt := newRequestOptions(opts...)

	// body
	bodyReader, err := opt.bodyReader()
	if err != nil {
		return nil, err
	}

	//  access_token
	url := c.config.Addr + path
	if opt.needAccessToken {
		if token, err := c.GetToken(ctx); err != nil {
			return nil, err
		} else {
			url += "?access_token=" + token
		}
	}

	return http.NewRequestWithContext(ctx, method, url, bodyReader)
}

func (c *Client) sendRequest(req *http.Request, resp interface{}, opts ...responseOption) error {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return c.decodeResponse(res.Body, resp, opts...)
}

func (c *Client) decodeResponse(body io.Reader, resp interface{}, opts ...responseOption) error {
	opt := newResponseOptions(opts...)

	if !opt.needDecrypt {
		return json.NewDecoder(body).Decode(resp)
	}

	return c.decodeResponseWithDecrypt(body, resp, opts...)
}

func (c *Client) decodeResponseWithDecrypt(body io.Reader, resp interface{}, opts ...responseOption) error {
	var r response
	if err := json.NewDecoder(body).Decode(&r); err != nil {
		return err
	}

	if r.ErrCode != 0 {
		return newError(r.ErrCode, r.ErrMsg)
	}
	if r.Encrypt == "" {
		return newError(-1, "encrypt is empty")
	}

	rawData, err := c.encryptor.Decrypt(r.Encrypt)
	if err != nil {
		return err
	}

	if rawData.AppId != c.config.AppId {
		return newError(-1, "app_id not match")
	}

	if rawData.Data == nil {
		return newError(-1, "data is nil")
	}

	return json.Unmarshal(rawData.Data, resp)
}
