package youdu

import (
	"context"
	"net/http"
	"strconv"
)

type DeptItem struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	ParentId int    `json:"parentId"`
	SortId   int    `json:"sortId"`
}

type DeptListResponse struct {
	DeptList []DeptItem `json:"deptList"`
}

type DeptAliasItem struct {
	Id    int    `json:"id"`
	Alias string `json:"alias"`
}

type DeptAliasListResponse struct {
	AliasList []DeptAliasItem `json:"aliasList"`
}

type DeptIdByAliasResponse struct {
	Id int `json:"id"`
}

type CreateDeptRequest struct {
	Name     string `json:"name"`
	Alias    string `json:"alias"`
	Id       int    `json:"id"`
	ParentId int    `json:"parentId"`
	SortId   int    `json:"sortId"`
}

type CreateDeptResponse struct {
	Id int `json:"id"`
}

type UpdateDeptRequest struct {
	Name     string `json:"name"`
	Alias    string `json:"alias"`
	Id       int    `json:"id"`
	ParentId int    `json:"parentId"`
	SortId   int    `json:"sortId"`
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

func (c *Client) GetDeptAliasList(ctx context.Context) (response DeptAliasListResponse, err error) {
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

func (c *Client) GetDeptIdByAlias(ctx context.Context, alias string) (response DeptIdByAliasResponse, err error) {
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

func (c *Client) CreateDept(ctx context.Context, request CreateDeptRequest) (response CreateDeptResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/cgi/dept/create",
		withRequestBody(request),
		withRequestAccessToken(),
		withRequestEncrypt(),
	)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseDecrypt())
	return
}

func (c *Client) UpdateDept(ctx context.Context, request UpdateDeptRequest) (response Response, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/cgi/dept/update",
		withRequestBody(request),
		withRequestAccessToken(),
		withRequestEncrypt(),
	)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

func (c *Client) DeleteDept(ctx context.Context, deptId int) (response Response, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/cgi/dept/delete",
		withRequestAccessToken(),
		withRequestParamsKV("id", strconv.Itoa(deptId)),
	)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}
