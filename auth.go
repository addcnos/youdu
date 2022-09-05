package youdu

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

const (
	identifyUrl    = "/cgi/identify"
	userSetAuthUrl = "/cgi/user/setauth"
)

type Auth struct {
	config *Config
}

func NewAuth(config *Config) *Auth {
	return &Auth{
		config: config,
	}
}

type IdentifyResp struct {
	Buin   int `json:"buin"`
	Status struct {
		Code      int    `json:"code"`
		Message   string `json:"message"`
		CreatedAt string `json:"createdAt"`
	} `json:"status"`
	UserInfo struct {
		Gid        int    `json:"gid"`
		Account    string `json:"account"`
		ChsName    string `json:"chsName"`
		EngName    string `json:"engName"`
		Gender     int    `json:"gender"`
		OrgId      int    `json:"orgId"`
		Mobile     string `json:"mobile"`
		Phone      string `json:"phone"`
		Email      string `json:"email"`
		CustomAttr string `json:"customAttr"`
	} `json:"userInfo"`
}

// Identify 单点登录
func (a *Auth) Identify(token string) (i IdentifyResp, err error) {
	resp, err := a.config.GetHttp().Get(identifyUrl, map[string]string{
		"token": token,
	})
	if err != nil {
		return
	}

	if !resp.IsSuccess() {
		err = errors.New("Response status code is " + strconv.Itoa(resp.StatusCode()))
		return
	}

	if err = json.Unmarshal(resp.Body(), &i); err != nil {
		return
	}

	return
}

type SetAuthResp struct {
	FromUser   string `json:"fromUser"`
	CreateTime int    `json:"createTime"`
	PackageId  int    `json:"packageId"`
	MsgType    string `json:"msgType"`
	Passwd     string `json:"passwd"`
}

// SetAuth 第三方认证-设置认证信息
func (a *Auth) SetAuth(userId, password string) (bool, error) {
	accessToken, err := a.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return false, err
	}

	bodyJson, err := json.Marshal(map[string]interface{}{
		"userId":   userId,
		"authType": 2,
		"passwd":   fmt.Sprintf("%x", md5.Sum([]byte(password))),
	})
	if err != nil {
		return false, err
	}

	encrypt, err := a.config.GetEncryptor().Encrypt(string(bodyJson))
	if err != nil {
		return false, err
	}

	resp, err := a.config.GetHttp().Post(userSetAuthUrl+"?accessToken="+accessToken, map[string]interface{}{
		"appId":   a.config.AppId,
		"buin":    a.config.Buin,
		"encrypt": encrypt,
	})

	if err != nil {
		return false, err
	}

	if !resp.IsSuccess() {
		return false, errors.New("Response status code is " + strconv.Itoa(resp.StatusCode()))
	}

	jsonRet, err := resp.Json()
	if err != nil {
		return false, err
	}

	if jsonRet["errcode"].(float64) != 0 {
		return false, errors.New(jsonRet["errMsg"].(string))
	}

	return true, nil
}
