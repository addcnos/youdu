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

func (c *Client) newRequest(ctx context.Context, method string, path string, opts ...requestOption) (req *http.Request, err error) {
	var (
		opt = newRequestOptions(opts...)
		url = c.config.Addr + path
	)

	// body
	bodyReader, err := c.encodeRequestBody(opt)
	if err != nil {
		return nil, err
	}

	// access_token
	if opt.needAccessToken {
		if token, err := c.GetToken(ctx); err != nil {
			return nil, err
		} else {
			opt.params.Add("accessToken", token)
		}
	}

	req, err = http.NewRequestWithContext(ctx, method, url+"?"+opt.params.Encode(), bodyReader)
	return
}

func (c *Client) encodeRequestBody(opt *requestOptions) (io.Reader, error) {
	if opt.body == nil {
		return nil, nil
	}

	if !opt.needEncrypt {
		return opt.bodyReader(opt.body)
	}

	reqBytes, err := json.Marshal(opt.body)
	if err != nil {
		return nil, err
	}

	cipherText, err := c.encryptor.Encrypt(reqBytes)
	if err != nil {
		return nil, err
	}

	return opt.bodyReader(Request{
		Buin:    c.config.Buin,
		AppId:   c.config.AppId,
		Encrypt: cipherText,
	})

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
	var r Response
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

	if rawData.Data == nil {
		return newError(-1, "data is nil")
	}

	return json.Unmarshal(rawData.Data, resp)
}
