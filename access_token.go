package youdu

import (
	"errors"
	"strconv"
	"time"
)

const getTokenUrl = "/cgi/gettoken"

type accessToken struct {
	accessToken string
	expire      time.Time
}

type accessTokenProvider struct {
	config      *Config
	accessToken accessToken
}

func NewAccessTokenProvider(config *Config) *accessTokenProvider {
	return &accessTokenProvider{
		config: config,
	}
}

func (p *accessTokenProvider) GetAccessToken() (string, error) {
	// from memory
	if at := p.accessToken.GetAccessToken(); at != "" && p.accessToken.Expire().After(time.Now()) {
		return at, nil
	}

	// from cache
	// todo: support cache

	return p.Refresh()
}

func (p *accessTokenProvider) Refresh() (string, error) {
	// encrypt
	encrypt, err := p.config.GetEncryptor().Encrypt(strconv.Itoa(int(time.Now().Unix())))
	if err != nil {
		return "", err
	}

	resp, err := p.config.GetHttp().Post(getTokenUrl, map[string]interface{}{
		"appId":   p.config.AppId,
		"buin":    p.config.Buin,
		"encrypt": encrypt,
	})

	if err != nil {
		return "", err
	}

	if !resp.IsSuccess() {
		return "", errors.New("Response status code is " + strconv.Itoa(resp.StatusCode()))
	}

	jsonRet, err := resp.Json()
	if err != nil {
		return "", err
	}

	decrypt, err := p.config.GetEncryptor().Decrypt(jsonRet["encrypt"].(string))
	if err != nil {
		return "", err
	}

	var v struct {
		AccessToken string `json:"accessToken"`
		ExpireIn    int    `json:"expireIn"`
	}
	if err := decrypt.Unmarshal(&v); err != nil {
		return "", err
	}

	p.accessToken.accessToken = v.AccessToken
	p.accessToken.expire = time.Now().Add(time.Duration(v.ExpireIn-600) * time.Second) // 提前10分钟失效

	return p.accessToken.accessToken, nil
}

func (a accessToken) GetAccessToken() string {
	return a.accessToken
}

func (a accessToken) Expire() time.Time {
	return a.expire
}
