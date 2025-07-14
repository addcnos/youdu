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
	bodyDecrypt bool
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

func withResponseBodyDecrypt() responseOption {
	return func(args *responseOptions) {
		args.bodyDecrypt = true

	}
}

func (c *Client) decodeResponse(body io.Reader, resp any, opts ...responseOption) error {
	opt := newResponseOptions(opts...)

	if opt.bodyDecrypt {
		return c.decodeResponseWithBodyDecrypt(body, resp, opts...)
	}

	if !opt.needDecrypt {
		return json.NewDecoder(body).Decode(resp)
	}

	return c.decodeResponseWithDecrypt(body, resp, opts...)
}

func (c *Client) decodeResponseWithDecrypt(body io.Reader, resp any, _ ...responseOption) error {
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

func (c *Client) decodeResponseWithBodyDecrypt(body io.Reader, resp any, _ ...responseOption) error {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	rawData, err := c.encryptor.Decrypt(string(bodyBytes))
	if err != nil {
		return err
	}

	if rawData.Data == nil {
		return newError(-1, "decrypted data is nil")
	}
	if target, ok := resp.(*[]byte); ok {
		*target = rawData.Data
		return nil
	}
	return nil
}
