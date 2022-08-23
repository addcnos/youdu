package youdu

import (
	"errors"
	"strconv"
)

const (
	deptListUrl  = "/cgi/dept/list"
	deptGetIdUrl = "/cgi/dept/getid"
)

type Dept struct {
	config *Config
}

type DeptItem struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	ParentId int    `json:"parentId"`
	SortId   int    `json:"sortId"`
}

func NewDept(config *Config) *Dept {
	return &Dept{
		config: config,
	}
}

// GetList 获取部门列表
func (d *Dept) GetList(depId int) ([]DeptItem, error) {
	accessToken, err := d.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return nil, err
	}

	resp, err := d.config.GetHttp().Get(deptListUrl, map[string]string{
		"id":          strconv.Itoa(depId),
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

	decrypt, err := d.config.GetEncryptor().Decrypt(jsonRet["encrypt"].(string))
	if err != nil {
		return nil, err
	}

	var v map[string][]DeptItem
	if err := decrypt.Unmarshal(&v); err != nil {
		return nil, err
	}

	return v["deptList"], nil
}

func (d *Dept) GetId(alias string) {
	// accessToken, err := d.config.GetAccessTokenProvider().GetAccessToken()
	// if err != nil {
	// 	return nil, err
	// }
	//
	// resp, err := d.config.GetHttp().Get(deptGetIdUrl, map[string]string{
	// 	"alias":       alias,
	// 	"accessToken": accessToken,
	// })
	//
	// if err != nil {
	// 	return nil, err
	// }
	//
	// if !resp.IsSuccess() {
	// 	return nil, errors.New("Response status code is " + strconv.Itoa(resp.StatusCode()))
	// }
	//
	// jsonRet, err := resp.Json()
	// if err != nil {
	// 	return nil, err
	// }
	//
	// decrypt, err := d.config.GetEncryptor().Decrypt(jsonRet["encrypt"].(string))
	// if err != nil {
	// 	return nil, err
	// }
	//
	// fmt.Println(decrypt)
	//
	// return nil, nil
}
