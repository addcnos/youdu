package youdu

import (
	"bytes"
	"encoding/json"
	"io"
	"net/url"
)

type Request struct {
	Buin    int    `json:"buin"`
	AppId   string `json:"appId"`
	Encrypt string `json:"encrypt"`
}

type Response struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Encrypt string `json:"encrypt,omitempty"`
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

func withRequestParams(params url.Values) requestOption {
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
