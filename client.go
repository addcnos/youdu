package youdu

import "net/http"

type Config struct {
	Addr   string
	Buin   int
	AppID  string
	AesKey string
}

type Client struct {
	config *Config
	token  *token

	encryptor  *Encryptor
	httpClient *http.Client
}

func WithEncryptor(encryptor *Encryptor) func(c *Client) {
	return func(c *Client) {
		c.encryptor = encryptor
	}
}

func WithHTTPClient(client *http.Client) func(c *Client) {
	return func(c *Client) {
		c.httpClient = client
	}
}

type ClientOption func(c *Client)

func NewClient(config *Config, opts ...ClientOption) *Client {
	c := &Client{
		config: config,
	}

	for _, opt := range opts {
		opt(c)
	}

	if c.httpClient == nil {
		c.httpClient = http.DefaultClient
	}

	if c.encryptor == nil {
		c.encryptor = NewEncryptorWithConfig(config)
	}

	return c
}
