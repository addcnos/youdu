package youdu

import (
	"github.com/go-resty/resty/v2"
)

type Http struct {
	config *Config
	client *resty.Client
}

func NewHttp(config *Config) *Http {
	return &Http{
		client: resty.New(),
		config: config,
	}
}

func (h *Http) Request(method, url string, params interface{}, fn ...func(*resty.Request)) (*Response, error) {
	var (
		req  *resty.Request
		resp *resty.Response
		err  error
	)

	req = h.client.R()

	if len(fn) > 0 {
		for _, f := range fn {
			f(req)
		}
	}

	if method == "POST" {
		resp, err = req.SetBody(params).Post(h.config.Api + url)
	} else {
		resp, err = req.SetQueryParams(params.(map[string]string)).Get(h.config.Api + url)
	}

	if err != nil {
		return nil, err
	}

	return NewResponse(resp), nil
}

func (h *Http) Get(url string, params map[string]string, fn ...func(*resty.Request)) (*Response, error) {
	return h.Request("GET", url, params, fn...)
}

func (h *Http) Post(url string, params interface{}, fn ...func(*resty.Request)) (*Response, error) {
	return h.Request("POST", url, params, fn...)
}
