package youdu

import (
	"context"
	"net/http"
	"strconv"
)

type DeptItem struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ParentID int    `json:"parentId"`
	SortID   int    `json:"sortId"`
}

type DeptListResponse struct {
	DeptList []DeptItem `json:"deptList"`
}

type DeptAliasItem struct {
	ID    int    `json:"id"`
	Alias string `json:"alias"`
}

type DeptAliasListResponse struct {
	AliasList []DeptAliasItem `json:"aliasList"`
}

type DeptIDByAliasResponse struct {
	ID int `json:"id"`
}

type CreateDeptRequest struct {
	Name     string `json:"name"`
	Alias    string `json:"alias"`
	ID       int    `json:"id"`
	ParentID int    `json:"parentId"`
	SortID   int    `json:"sortId"`
}

type CreateDeptResponse struct {
	ID int `json:"id"`
}

type UpdateDeptRequest struct {
	Name     string `json:"name"`
	Alias    string `json:"alias"`
	ID       int    `json:"id"`
	ParentID int    `json:"parentId"`
	SortID   int    `json:"sortId"`
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

func (c *Client) GetDeptIDByAlias(ctx context.Context, alias string) (response DeptIDByAliasResponse, err error) {
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

func (c *Client) DeleteDept(ctx context.Context, deptID int) (response Response, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/cgi/dept/delete",
		withRequestAccessToken(),
		withRequestParamsKV("id", strconv.Itoa(deptID)),
	)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}
