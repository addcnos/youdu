package youdu

import (
	"context"
	"net/http"
	"strconv"
)

type DeptList struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	ParentId int    `json:"parentId"`
	SortId   int    `json:"sortId"`
}

type DeptListResponse struct {
	DeptList []DeptList `json:"deptList"`
}

type AliasList struct {
	Id    int    `json:"id"`
	Alias string `json:"alias"`
}

type AliasListResponse struct {
	AliasList []AliasList `json:"aliasList"`
}

type AliaIdResponse struct {
	Id int `json:"id"`
}

func (c *Client) GetDeptList(ctx context.Context, id ...int) (response DeptListResponse, err error) {
	opts := []requestOption{
		withRequestAccessToken(),
		withRequestEncrypt(),
	}

	if len(id) > 0 {
		opts = append(opts, withRequestParamsKV("id", strconv.Itoa(id[0])))
	} else {
		opts = append(opts, withRequestParamsKV("id", "0"))
	}

	req, err := c.newRequest(ctx, http.MethodGet, "/cgi/dept/list", opts...)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseDecrypt())
	return
}

func (c *Client) GetDeptIdByAlias(ctx context.Context) (response AliasListResponse, err error) {
	opts := []requestOption{
		withRequestAccessToken(),
		withRequestEncrypt(),
	}

	req, err := c.newRequest(ctx, http.MethodGet, "/cgi/dept/getid", opts...)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseDecrypt())
	return
}

func (c *Client) GetDeptId(ctx context.Context, alias string) (response AliaIdResponse, err error) {
	opts := []requestOption{
		withRequestAccessToken(),
		withRequestEncrypt(),
		withRequestParamsKV("alias", alias),
	}

	req, err := c.newRequest(ctx, http.MethodGet, "/cgi/dept/getid", opts...)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseDecrypt())
	return
}
