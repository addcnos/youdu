package youdu

import (
	"encoding/json"
	"io"
)

type Response struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Encrypt string `json:"encrypt,omitempty"`
}

type responseOptions struct {
	needDecrypt bool
}

type responseOption func(*responseOptions)

func newResponseOptions(opts ...responseOption) *responseOptions {
	args := &responseOptions{}

	for _, opt := range opts {
		opt(args)
	}

	return args
}

func withResponseDecrypt() responseOption {
	return func(args *responseOptions) {
		args.needDecrypt = true
	}
}

func (c *Client) decodeResponse(body io.Reader, resp interface{}, opts ...responseOption) error {
	opt := newResponseOptions(opts...)

	if !opt.needDecrypt {
		return json.NewDecoder(body).Decode(resp)
	}

	return c.decodeResponseWithDecrypt(body, resp, opts...)
}

func (c *Client) decodeResponseWithDecrypt(body io.Reader, resp interface{}, _ ...responseOption) error {
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
