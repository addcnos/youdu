package youdu

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/addcnos/youdu/session"
)

const (
	sessionCreateUrl = "/cgi/session/create"
	sessionGetUrl    = "/cgi/session/get"
	sessionUpdateUrl = "/cgi/session/update"
	sessionSendUrl   = "/cgi/session/send"
)

type Session struct {
	config *Config
}

func NewSession(config *Config) *Session {
	return &Session{
		config: config,
	}
}

// Create 创建一个会话
// members 第一个默认为创建者
func (s *Session) Create(title string, members []string) (*session.Session, error) {
	if len(members) < 3 {
		return nil, errors.New("members must be at least 3")
	}

	accessToken, err := s.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return nil, err
	}

	bodyJson, err := json.Marshal(map[string]interface{}{
		"title":   title,
		"creator": members[0],
		"member":  members,
		"type":    "multi",
	})
	if err != nil {
		return nil, err
	}

	encrypt, err := s.config.GetEncryptor().Encrypt(string(bodyJson))
	if err != nil {
		return nil, err
	}

	resp, err := s.config.GetHttp().Post(sessionCreateUrl+"?accessToken="+accessToken, map[string]interface{}{
		"appId":   s.config.AppId,
		"buin":    s.config.Buin,
		"encrypt": encrypt,
	})
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, errors.New("Response status code is " + strconv.Itoa(resp.StatusCode()))
	}

	jsonRet, err := resp.Json()
	if err != nil {
		return nil, err
	}

	decrypt, err := s.config.GetEncryptor().Decrypt(jsonRet["encrypt"].(string))
	if err != nil {
		return nil, err
	}

	var v *session.Session
	if err := decrypt.Unmarshal(&v); err != nil {
		return nil, err
	}

	return v, nil
}

// Get 获取会话信息
func (s *Session) Get(sessionId string) (*session.Session, error) {
	accessToken, err := s.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return nil, err
	}

	resp, err := s.config.GetHttp().Get(sessionGetUrl+"?accessToken="+accessToken, map[string]string{
		"sessionId": sessionId,
	})
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, errors.New("Response status code is " + strconv.Itoa(resp.StatusCode()))
	}

	jsonRet, err := resp.Json()
	if err != nil {
		return nil, err
	}

	decrypt, err := s.config.GetEncryptor().Decrypt(jsonRet["encrypt"].(string))
	if err != nil {
		return nil, err
	}

	var v *session.Session
	if err := decrypt.Unmarshal(&v); err != nil {
		return nil, err
	}

	return v, nil
}

// Update 更新会话信息
func (s *Session) Update(sessionId, opUser, title string, addMembers, delMembers []string) (*session.Session, error) {
	accessToken, err := s.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return nil, err
	}

	bodyJson, err := json.Marshal(map[string]interface{}{
		"sessionId": sessionId,
		"opUser":    opUser,
		"title":     title,
		"addMember": addMembers,
		"delMember": delMembers,
	})
	if err != nil {
		return nil, err
	}

	encrypt, err := s.config.GetEncryptor().Encrypt(string(bodyJson))
	if err != nil {
		return nil, err
	}

	resp, err := s.config.GetHttp().Post(sessionUpdateUrl+"?accessToken="+accessToken, map[string]interface{}{
		"appId":   s.config.AppId,
		"buin":    s.config.Buin,
		"encrypt": encrypt,
	})
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, errors.New("Response status code is " + strconv.Itoa(resp.StatusCode()))
	}

	jsonRet, err := resp.Json()
	if err != nil {
		return nil, err
	}

	decrypt, err := s.config.GetEncryptor().Decrypt(jsonRet["encrypt"].(string))
	if err != nil {
		return nil, err
	}

	fmt.Println(decrypt)

	var v *session.Session
	if err := decrypt.Unmarshal(&v); err != nil {
		return nil, err
	}

	return v, nil
}

// Send 发送消息
func (s *Session) Send(message session.Message) error {
	accessToken, err := s.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return err
	}

	bodyJson, err := json.Marshal(message)
	if err != nil {
		return err
	}

	encrypt, err := s.config.GetEncryptor().Encrypt(string(bodyJson))
	if err != nil {
		return err
	}

	resp, err := s.config.GetHttp().Post(sessionSendUrl+"?accessToken="+accessToken, map[string]interface{}{
		"appId":   s.config.AppId,
		"buin":    s.config.Buin,
		"encrypt": encrypt,
	})
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return errors.New("Response status code is " + strconv.Itoa(resp.StatusCode()))
	}

	jsonRet, err := resp.Json()
	if err != nil {
		return err
	}

	if jsonRet["errcode"].(float64) != 0 {
		return errors.New(jsonRet["errmsg"].(string))
	}

	return nil
}
