package youdu

import (
	"context"
	"net/http"
)

type CreateGroupRequest struct {
	Name string `json:"name"`
}

type CreateGroupResponse struct {
	Id string `json:"id"`
}

type UpdateGroupRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GroupUpdateMemberRequest struct {
	Id       string   `json:"id"`
	UserList []string `json:"userList"`
}

type GroupInfoMember struct {
	Account string `json:"account"`
	Name    string `json:"name"`
	Mobile  string `json:"mobile"`
}

type GroupInfoResponse struct {
	Id      string            `json:"id"`
	Name    string            `json:"name"`
	Members []GroupInfoMember `json:"members"`
}

type GroupListItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GroupListResponse struct {
	GroupList []GroupListItem `json:"groupList"`
}

type IsGroupMemberResponse struct {
	Belong bool `json:"belong"`
}

func (c *Client) CreateGroup(ctx context.Context, request CreateGroupRequest) (response CreateGroupResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/cgi/group/create",
		withRequestBody(request), withRequestAccessToken(), withRequestEncrypt())
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseDecrypt())
	return
}

func (c *Client) DeleteGroup(ctx context.Context, groupId string) (response Response, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/cgi/group/delete",
		withRequestAccessToken(),
		withRequestParamsKV("id", groupId),
	)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

func (c *Client) UpdateGroup(ctx context.Context, request UpdateGroupRequest) (response Response, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/cgi/group/update",
		withRequestBody(request), withRequestAccessToken(), withRequestEncrypt())
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

func (c *Client) AddGroupMember(ctx context.Context, request GroupUpdateMemberRequest) (response Response, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/cgi/group/addmember",
		withRequestBody(request), withRequestAccessToken(), withRequestEncrypt())
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

func (c *Client) DelGroupMember(ctx context.Context, request GroupUpdateMemberRequest) (response Response, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/cgi/group/delmember",
		withRequestBody(request), withRequestAccessToken(), withRequestEncrypt())
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

func (c *Client) GetGroupInfo(ctx context.Context, groupId string) (response GroupInfoResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/cgi/group/info",
		withRequestAccessToken(),
		withRequestEncrypt(),
		withRequestParamsKV("id", groupId),
	)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseDecrypt())
	return
}

func (c *Client) GetGroupList(ctx context.Context, userId ...string) (response GroupListResponse, err error) {
	opts := []requestOption{
		withRequestAccessToken(),
		withRequestEncrypt(),
	}

	if len(userId) > 0 {
		opts = append(opts, withRequestParamsKV("userId", userId[0]))
	}

	req, err := c.newRequest(ctx, http.MethodGet, "/cgi/group/list", opts...)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseDecrypt())
	return
}

func (c *Client) IsGroupMember(ctx context.Context, groupId string, userId string) (response IsGroupMemberResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/cgi/group/ismember",
		withRequestAccessToken(),
		withRequestEncrypt(),
		withRequestParamsKV("id", groupId),
		withRequestParamsKV("userId", userId),
	)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseDecrypt())
	return
}
