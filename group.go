package youdu

import (
	"context"
	"net/http"
)

type CreateGroupRequest struct {
	Name string `json:"name"`
}

type CreateGroupResponse struct {
	ID string `json:"id"`
}

type UpdateGroupRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GroupUpdateMemberRequest struct {
	ID       string   `json:"id"`
	UserList []string `json:"userList"`
}

type GroupInfoMember struct {
	Account string `json:"account"`
	Name    string `json:"name"`
	Mobile  string `json:"mobile"`
}

type GroupInfoResponse struct {
	ID      string            `json:"id"`
	Name    string            `json:"name"`
	Members []GroupInfoMember `json:"members"`
}

type GroupItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GroupListResponse struct {
	GroupList []GroupItem `json:"groupList"`
}

type IsGroupMemberResponse struct {
	Belong bool `json:"belong"`
}

func (c *Client) CreateGroup(
	ctx context.Context,
	request CreateGroupRequest,
) (response CreateGroupResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/cgi/group/create",
		withRequestBody(request), withRequestAccessToken(), withRequestEncrypt())
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseDecrypt())
	return
}

func (c *Client) DeleteGroup(ctx context.Context, groupID string) (response Response, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/cgi/group/delete",
		withRequestAccessToken(),
		withRequestParamsKV("id", groupID),
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

func (c *Client) GetGroupInfo(ctx context.Context, groupID string) (response GroupInfoResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/cgi/group/info",
		withRequestAccessToken(),
		withRequestEncrypt(),
		withRequestParamsKV("id", groupID),
	)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseDecrypt())
	return
}

func (c *Client) GetGroupList(ctx context.Context, userID ...string) (response GroupListResponse, err error) {
	opts := []requestOption{
		withRequestAccessToken(),
		withRequestEncrypt(),
	}

	if len(userID) > 0 {
		opts = append(opts, withRequestParamsKV("userId", userID[0]))
	}

	req, err := c.newRequest(ctx, http.MethodGet, "/cgi/group/list", opts...)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseDecrypt())
	return
}

func (c *Client) IsGroupMember(
	ctx context.Context, groupID string, userID string,
) (response IsGroupMemberResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/cgi/group/ismember",
		withRequestAccessToken(),
		withRequestEncrypt(),
		withRequestParamsKV("id", groupID),
		withRequestParamsKV("userId", userID),
	)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseDecrypt())
	return
}
