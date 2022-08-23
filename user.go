package youdu

import (
	"errors"
	"strconv"
)

const (
	userGetUrl         = "/cgi/user/get"
	userListUrl        = "/cgi/user/list"
	userSimpleListUrl  = "/cgi/user/simplelist"
	userEnableStateUrl = "/cgi/user/enable/state"
)

type User struct {
	config *Config
}

func NewUser(config *Config) *User {
	return &User{
		config: config,
	}
}

type UserInfo struct {
	Gid        int    `json:"gid"`
	UserId     string `json:"userId"`
	Name       string `json:"name"`
	Gender     int    `json:"gender"` // 性别。0表示男性，1表示女性
	Mobile     string `json:"mobile"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Dept       []int  `json:"dept"`
	DeptDetail []struct {
		DeptId   int    `json:"deptId"`
		DeptName string `json:"deptName"`
		Position string `json:"position"`
		Weight   int    `json:"weight"`
		SortId   int    `json:"sortId"`
	} `json:"deptDetail"`
	Attrs []interface{} `json:"attrs"`
}

// Get 获取用户信息
// see: https://youdu.im/doc/api/c01_00013.html#_6
func (u *User) Get(userId string) (UserInfo, error) {
	accessToken, err := u.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return UserInfo{}, err
	}

	resp, err := u.config.GetHttp().Get(userGetUrl, map[string]string{
		"userId":      userId,
		"accessToken": accessToken,
	})

	if err != nil {
		return UserInfo{}, err
	}

	if !resp.IsSuccess() {
		return UserInfo{}, errors.New("Response status code is " + strconv.Itoa(resp.StatusCode()))
	}

	jsonRet, err := resp.Json()
	if err != nil {
		return UserInfo{}, err
	}

	decrypt, err := u.config.GetEncryptor().Decrypt(jsonRet["encrypt"].(string))
	if err != nil {
		return UserInfo{}, err
	}

	var v UserInfo
	if err := decrypt.Unmarshal(&v); err != nil {
		return UserInfo{}, err
	}

	return v, nil
}

// List 获取部门用户详细信息
func (u *User) List(deptId int) ([]UserInfo, error) {
	accessToken, err := u.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return nil, err
	}

	resp, err := u.config.GetHttp().Get(userListUrl, map[string]string{
		"deptId":      strconv.Itoa(deptId),
		"accessToken": accessToken,
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

	decrypt, err := u.config.GetEncryptor().Decrypt(jsonRet["encrypt"].(string))
	if err != nil {
		return nil, err
	}

	var v map[string][]UserInfo
	if err := decrypt.Unmarshal(&v); err != nil {
		return nil, err
	}

	return v["userList"], nil
}

type SimpleUserInfo struct {
	Gid    int    `json:"gid"`
	UserId string `json:"userId"`
	Name   string `json:"name"`
	Gender int    `json:"gender"`
	Dept   []int  `json:"dept"`
}

// SimpleList 获取部门用户
func (u *User) SimpleList(deptId int) ([]SimpleUserInfo, error) {
	accessToken, err := u.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return nil, err
	}

	resp, err := u.config.GetHttp().Get(userSimpleListUrl, map[string]string{
		"deptId":      strconv.Itoa(deptId),
		"accessToken": accessToken,
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

	decrypt, err := u.config.GetEncryptor().Decrypt(jsonRet["encrypt"].(string))
	if err != nil {
		return nil, err
	}

	var v map[string][]SimpleUserInfo
	if err := decrypt.Unmarshal(&v); err != nil {
		return nil, err
	}

	return v["userList"], nil
}

// EnableState 查询用户激活状态
func (u *User) EnableState(userId string) (int, error) {
	accessToken, err := u.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return 0, err
	}

	resp, err := u.config.GetHttp().Get(userEnableStateUrl, map[string]string{
		"userId":      userId,
		"accessToken": accessToken,
	})

	if err != nil {
		return 0, err
	}

	if !resp.IsSuccess() {
		return 0, errors.New("Response status code is " + strconv.Itoa(resp.StatusCode()))
	}

	jsonRet, err := resp.Json()
	if err != nil {
		return 0, err
	}

	decrypt, err := u.config.GetEncryptor().Decrypt(jsonRet["encrypt"].(string))
	if err != nil {
		return 0, err
	}

	var v map[string]int
	if err := decrypt.Unmarshal(&v); err != nil {
		return 0, err
	}

	return v["enableState"], nil
}
