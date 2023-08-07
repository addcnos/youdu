package youdu

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	Buin    int    `json:"buin"`
	AppId   string `json:"appId"`
	Encrypt string `json:"encrypt"`
}

type requestOptions struct {
	params          url.Values
	body            interface{}
	needEncrypt     bool
	needAccessToken bool
}

func newRequestOptions(opts ...requestOption) *requestOptions {
	args := &requestOptions{
		body:   nil,
		params: url.Values{},
	}

	for _, opt := range opts {
		opt(args)
	}

	return args
}

func (r *requestOptions) bodyReader(body interface{}) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}

	if v, ok := body.(io.Reader); ok {
		return v, nil
	}

	reqBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(reqBytes), nil
}

type requestOption func(*requestOptions)

func withRequestBody(body interface{}) requestOption {
	return func(args *requestOptions) {
		args.body = body
	}
}

func withRequestEncrypt() requestOption {
	return func(args *requestOptions) {
		args.needEncrypt = true
	}
}

func withRequestParams(params url.Values) requestOption { //nolint:unused
	return func(args *requestOptions) {
		args.params = params
	}
}

func withRequestParamsKV(key, value string) requestOption {
	return func(args *requestOptions) {
		args.params.Add(key, value)
	}
}

func withRequestAccessToken() requestOption {
	return func(args *requestOptions) {
		args.needAccessToken = true
	}
}

func (c *Client) newRequest(ctx context.Context, method string, path string, opts ...requestOption) (req *http.Request, err error) {
	var (
		opt     = newRequestOptions(opts...)
		urlPath = c.config.Addr + path
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

	req, err = http.NewRequestWithContext(ctx, method, urlPath+"?"+opt.params.Encode(), bodyReader)
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
