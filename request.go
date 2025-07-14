package youdu

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

var ErrUnexpectedResponseCode = errors.New("youdu sdk: unexpected response code")

type requestType string

type NormalRequest struct {
	Buin    int    `json:"buin"`
	AppID   string `json:"appId"`
	Encrypt string `json:"encrypt"`
}

type SpecialRequest struct {
	AppID      string `json:"app_Id"`
	MsgEncrypt string `json:"msg_encrypt"`
}

const (
	NormalRequestType  requestType = "normal"
	SpecialRequestType requestType = "special"
	UploadRequestType  requestType = "upload"
)

type requestOptions struct {
	params          url.Values
	body            any
	needEncrypt     bool
	needAccessToken bool
	requestType     requestType
	contentType     string
}

func newRequestOptions(opts ...requestOption) *requestOptions {
	args := &requestOptions{
		body:        nil,
		params:      url.Values{},
		requestType: NormalRequestType,
	}

	for _, opt := range opts {
		opt(args)
	}

	return args
}

func (r *requestOptions) bodyReader(body any) (io.Reader, error) {
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

func withRequestBody(body any) requestOption {
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

func withRequestType(rt requestType) requestOption {
	return func(args *requestOptions) {
		args.requestType = rt
	}
}

func (c *Client) newRequest(
	ctx context.Context, method string, path string, opts ...requestOption,
) (req *http.Request, err error) {
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
		token, err := c.GetToken(ctx)
		if err != nil {
			return nil, err
		}
		opt.params.Add("accessToken", token)
	}

	req, err = http.NewRequestWithContext(ctx, method, urlPath+"?"+opt.params.Encode(), bodyReader)

	if opt.contentType != "" {
		req.Header.Set("Content-Type", opt.contentType)
	}

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

	switch opt.requestType {
	case SpecialRequestType:
		return opt.bodyReader(SpecialRequest{
			AppID:      c.config.AppID,
			MsgEncrypt: cipherText,
		})
	case NormalRequestType:
		return opt.bodyReader(NormalRequest{
			Buin:    c.config.Buin,
			AppID:   c.config.AppID,
			Encrypt: cipherText,
		})
	case UploadRequestType:
		bodyReader, err := c.uploadRequestBody(opt)
		if err != nil {
			return nil, err
		}
		return opt.bodyReader(bodyReader)
	default:
		return nil, errors.New("youdu sdk: unknown request type")
	}
}

func (c *Client) uploadRequestBody(opt *requestOptions) (any, error) {
	req := opt.body
	uploadReq, ok := req.(UploadMediaRequest)
	if !ok {
		return nil, fmt.Errorf("invalid request type %T", req)
	}

	body := bytes.NewBufferString("")
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("buin", fmt.Sprint(c.config.Buin)); err != nil {
		return nil, err
	}
	if err := writer.WriteField("appId", c.config.AppID); err != nil {
		return nil, err
	}

	meta := struct {
		Type string `json:"type"`
		Name string `json:"name"`
	}{uploadReq.FileType, uploadReq.FileName}

	metaBytes, err := json.Marshal(meta)
	if err != nil {
		return nil, err
	}

	encryptedMeta, err := c.encryptor.Encrypt(metaBytes)
	if err != nil {
		return nil, err
	}

	if err := writer.WriteField("encrypt", encryptedMeta); err != nil {
		return nil, err
	}

	filePart, err := writer.CreateFormFile("file", uploadReq.FileName)
	if err != nil {
		return nil, err
	}

	encryptedFile, err := c.encryptor.Encrypt(uploadReq.File)
	if err != nil {
		return nil, err
	}

	if _, err := filePart.Write([]byte(encryptedFile)); err != nil {
		return nil, err
	}
	opt.contentType = writer.FormDataContentType()

	if err := writer.Close(); err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) sendRequest(req *http.Request, resp any, opts ...responseOption) error {
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close() // nolint:errcheck

	if res.StatusCode != http.StatusOK {
		return ErrUnexpectedResponseCode
	}

	return c.decodeResponse(res.Body, resp, opts...)
}
