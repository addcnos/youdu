package youdu

import (
	"encoding/json"
	"errors"
	"strconv"
)

const (
	groupCreateUrl    = "/cgi/group/create"
	groupDeleteUrl    = "/cgi/group/delete"
	groupUpdateUrl    = "/cgi/group/update"
	groupInfoUrl      = "/cgi/group/info"
	groupListUrl      = "/cgi/group/list"
	groupAddMemberUrl = "/cgi/group/addmember"
	groupDelMemberUrl = "/cgi/group/delmember"
	groupIsMemberUrl  = "/cgi/group/ismember"
)

type Group struct {
	config *Config
}

func NewGroup(config *Config) *Group {
	return &Group{
		config: config,
	}
}

// Create 创建一个群组
func (g *Group) Create(name string) (string, error) {
	accessToken, err := g.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return "", err
	}

	bodyJson, err := json.Marshal(map[string]interface{}{
		"name": name,
	})
	if err != nil {
		return "", err
	}

	encrypt, err := g.config.GetEncryptor().Encrypt(string(bodyJson))
	if err != nil {
		return "", err
	}

	resp, err := g.config.GetHttp().Post(groupCreateUrl+"?accessToken="+accessToken, map[string]interface{}{
		"appId":   g.config.AppId,
		"buin":    g.config.Buin,
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

	decrypt, err := g.config.GetEncryptor().Decrypt(jsonRet["encrypt"].(string))
	if err != nil {
		return "", err
	}

	var v map[string]string
	if err := decrypt.Unmarshal(&v); err != nil {
		return "", err
	}

	return v["id"], nil
}

func (g *Group) Delete(groupId string) (bool, error) {
	accessToken, err := g.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return false, err
	}

	bodyJson, err := json.Marshal(map[string]interface{}{
		"id": groupId,
	})
	if err != nil {
		return false, err
	}

	encrypt, err := g.config.GetEncryptor().Encrypt(string(bodyJson))
	if err != nil {
		return false, err
	}

	resp, err := g.config.GetHttp().Post(groupDeleteUrl+"?accessToken="+accessToken, map[string]interface{}{
		"appId":   g.config.AppId,
		"buin":    g.config.Buin,
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

func (g *Group) Update(groupId, groupName string) (bool, error) {
	accessToken, err := g.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return false, err
	}

	bodyJson, err := json.Marshal(map[string]interface{}{
		"id":   groupId,
		"name": groupName,
	})
	if err != nil {
		return false, err
	}

	encrypt, err := g.config.GetEncryptor().Encrypt(string(bodyJson))
	if err != nil {
		return false, err
	}

	resp, err := g.config.GetHttp().Post(groupUpdateUrl+"?accessToken="+accessToken, map[string]interface{}{
		"appId":   g.config.AppId,
		"buin":    g.config.Buin,
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

type GroupInfo struct {
	Id           string      `json:"id"`
	Name         string      `json:"name"`
	Admins       interface{} `json:"admins"`
	BelongDeptId int         `json:"belongDeptId"`
	IsDeptGroup  bool        `json:"isDeptGroup"`
	Master       int         `json:"master"`
	Members      []struct {
		Account string `json:"account"`
		Name    string `json:"name"`
		Mobile  string `json:"mobile"`
	} `json:"members"`
}

func (g *Group) Info(groupId string) (*GroupInfo, error) {
	accessToken, err := g.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return nil, err
	}

	resp, err := g.config.GetHttp().Get(groupInfoUrl+"?accessToken="+accessToken, map[string]string{
		"id": groupId,
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

	decrypt, err := g.config.GetEncryptor().Decrypt(jsonRet["encrypt"].(string))
	if err != nil {
		return nil, err
	}

	var v *GroupInfo
	if err = decrypt.Unmarshal(&v); err != nil {
		return nil, err
	}

	return v, nil
}

type GroupItem struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Version      int    `json:"version"`
	IsDeptGroup  bool   `json:"isDeptGroup"`
	BelongDeptId int    `json:"belongDeptId"`
}

func (g *Group) List(userId ...string) ([]GroupItem, error) {
	accessToken, err := g.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return nil, err
	}

	params := map[string]string{}
	if len(userId) > 0 {
		params["userId"] = userId[0]
	}

	resp, err := g.config.GetHttp().Get(groupListUrl+"?accessToken="+accessToken, params)

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

	decrypt, err := g.config.GetEncryptor().Decrypt(jsonRet["encrypt"].(string))
	if err != nil {
		return nil, err
	}

	var v map[string][]GroupItem
	if err := decrypt.Unmarshal(&v); err != nil {
		return nil, err
	}

	return v["groupList"], nil
}

func (g *Group) AddMember(groupId string, userId ...string) (bool, error) {
	accessToken, err := g.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return false, err
	}

	bodyJson, err := json.Marshal(map[string]interface{}{
		"id":       groupId,
		"userList": userId,
	})
	if err != nil {
		return false, err
	}

	encrypt, err := g.config.GetEncryptor().Encrypt(string(bodyJson))
	if err != nil {
		return false, err
	}

	resp, err := g.config.GetHttp().Post(groupAddMemberUrl+"?accessToken="+accessToken, map[string]interface{}{
		"appId":   g.config.AppId,
		"buin":    g.config.Buin,
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

func (g *Group) DelMember(groupId string, userId ...string) (bool, error) {
	accessToken, err := g.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return false, err
	}

	bodyJson, err := json.Marshal(map[string]interface{}{
		"id":       groupId,
		"userList": userId,
	})
	if err != nil {
		return false, err
	}

	encrypt, err := g.config.GetEncryptor().Encrypt(string(bodyJson))
	if err != nil {
		return false, err
	}

	resp, err := g.config.GetHttp().Post(groupDelMemberUrl+"?accessToken="+accessToken, map[string]interface{}{
		"appId":   g.config.AppId,
		"buin":    g.config.Buin,
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

func (g *Group) IsMember(groupId, userId string) (bool, error) {
	accessToken, err := g.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return false, err
	}

	resp, err := g.config.GetHttp().Get(groupIsMemberUrl+"?accessToken="+accessToken, map[string]string{
		"id":     groupId,
		"userId": userId,
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

	if jsonRet["errcode"] != 0 {
		return false, errors.New(jsonRet["errmsg"].(string))
	}

	decrypt, err := g.config.GetEncryptor().Decrypt(jsonRet["encrypt"].(string))
	if err != nil {
		return false, err
	}

	var v map[string]bool
	if err := decrypt.Unmarshal(&v); err != nil {
		return false, err
	}

	return v["belong"], nil
}
