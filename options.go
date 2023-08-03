package youdu

import (
	"bytes"
	"encoding/json"
	"io"
)

type request struct {
	Buin    int    `json:"buin"`
	AppId   string `json:"appId"`
	Encrypt string `json:"encrypt"`
}

type response struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Encrypt string `json:"encrypt"`
}

type requestOptions struct {
	body            interface{}
	needEncrypt     bool
	needAccessToken bool
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

func newRequestOptions(opts ...requestOption) *requestOptions {
	args := &requestOptions{
		body: nil,
	}

	for _, opt := range opts {
		opt(args)
	}

	return args
}

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
