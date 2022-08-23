package youdu

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

type Response struct {
	restyResponse *resty.Response
	encryptor     *encryptor

	decryptResult *DecryptResult
}

func NewResponse(restyResponse *resty.Response) *Response {
	return &Response{
		restyResponse: restyResponse,
	}
}

func (r *Response) Body() []byte {
	return r.restyResponse.Body()
}

func (r *Response) String() string {
	return r.restyResponse.String()
}

func (r *Response) Json() (map[string]interface{}, error) {
	var v map[string]interface{}
	if err := json.Unmarshal(r.restyResponse.Body(), &v); err != nil {
		return nil, err
	}

	return v, nil
}

func (r *Response) StatusCode() int {
	return r.restyResponse.StatusCode()
}

func (r *Response) Header() map[string][]string {
	return r.restyResponse.Header()
}

func (r *Response) IsSuccess() bool {
	return r.StatusCode() == 200
}
