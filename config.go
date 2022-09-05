package youdu

import "strings"

type Config struct {
	Api    string
	Buin   int
	AppId  string
	AesKey string
	Path   string

	encryptor           *encryptor
	http                *Http
	accessTokenProvider *accessTokenProvider
}

func (c *Config) GetEncryptor() *encryptor {
	if c.encryptor == nil {
		c.encryptor = NewEncryptor(c)
	}

	return c.encryptor
}

func (c *Config) GetHttp() *Http {
	if c.http == nil {
		c.http = NewHttp(c)
	}

	return c.http
}

func (c *Config) GetAccessTokenProvider() *accessTokenProvider {
	if c.accessTokenProvider == nil {
		c.accessTokenProvider = NewAccessTokenProvider(c)
	}

	return c.accessTokenProvider
}

// GetPath 返回系统默认路径
func (c *Config) GetPath() string {
	if c.Path == "" {
		c.Path = c.GetDefaultPath()
	}

	return strings.TrimRight(c.Path, "/")
}

func (c *Config) GetDefaultPath() string {
	return "/tmp"
}
