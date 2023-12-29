package youdu

import (
	"context"
	"net/http"
	"strconv"
)

type CreateUserRequest struct {
	UserId      string      `json:"userId"`
	Name        string      `json:"name"`
	Gender      int         `json:"gender"`
	Mobile      string      `json:"mobile"`
	Phone       string      `json:"phone"`
	Email       string      `json:"email"`
	Dept        []int       `json:"dept"`
	EnableState EnableState `json:"enableState"`
	ShortCode   string      `json:"shortCode"`
}

type UpdateUserRequest struct {
	UserId    string `json:"userId"`
	Name      string `json:"name"`
	Gender    int    `json:"gender"`
	Mobile    string `json:"mobile"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Dept      []int  `json:"dept"`
	ShortCode string `json:"shortCode"`
}

type UpdateUserPositionRequest struct {
	UserId   string `json:"userId"`
	DeptId   int    `json:"deptId"`
	Position string `json:"position"`
	Weight   int    `json:"weight"`
	SortId   int    `json:"sortId"`
}

type BatchDeleteUserRequest struct {
	DelList []string `json:"delList"`
}

type DeptDetail struct {
	DeptId   int    `json:"deptId"`
	Position string `json:"position"`
	Weight   int    `json:"weight"`
	SortId   int    `json:"sortId"`
}

type UserResponse struct {
	UserId     string       `json:"userId"`
	Name       string       `json:"name"`
	Gender     int          `json:"gender"`
	Mobile     string       `json:"mobile"`
	Phone      string       `json:"phone"`
	Email      string       `json:"email"`
	Dept       []int        `json:"dept"`
	DeptDetail []DeptDetail `json:"deptDetail"`
}

type UserItem struct {
	UserId     string       `json:"userId"`
	Name       string       `json:"name"`
	Gender     int          `json:"gender"`
	Mobile     string       `json:"mobile,omitempty"`
	Phone      string       `json:"phone,omitempty"`
	Email      string       `json:"email,omitempty"`
	Dept       []int        `json:"dept"`
	DeptDetail []DeptDetail `json:"deptDetail,omitempty"`
}

type DeptUserListResponse struct {
	UserList []UserItem `json:"userList"`
}

type UserEnableStateResponse struct {
	EnableState int `json:"enableState"`
}

type UpdateUserEnableStateRequest struct {
	UserIdList  []string    `json:"userIdList"`
	EnableState EnableState `json:"enableState"`
}

func (c *Client) CreateUser(ctx context.Context, request CreateUserRequest) (response Response, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/cgi/user/create",
		withRequestAccessToken(),
		withRequestEncrypt(),
		withRequestBody(request),
	)

	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

func (c *Client) UpdateUser(ctx context.Context, request UpdateUserRequest) (response Response, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/cgi/user/update",
		withRequestAccessToken(),
		withRequestEncrypt(),
		withRequestBody(request),
	)

	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

func (c *Client) UpdateUserPosition(ctx context.Context, request UpdateUserPositionRequest) (response Response, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/cgi/user/positionupdate",
		withRequestAccessToken(),
		withRequestEncrypt(),
		withRequestBody(request),
	)

	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

func (c *Client) DeleteUser(ctx context.Context, userId string) (response Response, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/cgi/user/delete",
		withRequestAccessToken(),
		withRequestEncrypt(),
		withRequestParamsKV("userId", userId),
	)

	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

func (c *Client) BatchDeleteUser(ctx context.Context, request BatchDeleteUserRequest) (response Response, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/cgi/user/batchdelete",
		withRequestAccessToken(),
		withRequestEncrypt(),
		withRequestBody(request),
	)

	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

func (c *Client) GetUser(ctx context.Context, userId string) (response UserResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/cgi/user/get",
		withRequestAccessToken(),
		withRequestEncrypt(),
		withRequestParamsKV("userId", userId),
	)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseDecrypt())
	return
}

func (c *Client) GetDeptUserList(ctx context.Context, deptId int) (response DeptUserListResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/cgi/user/list",
		withRequestAccessToken(),
		withRequestEncrypt(),
		withRequestParamsKV("deptId", strconv.Itoa(deptId)),
	)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseDecrypt())
	return
}

func (c *Client) GetDeptUserSimpleList(ctx context.Context, deptId int) (response DeptUserListResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/cgi/user/simplelist",
		withRequestAccessToken(),
		withRequestEncrypt(),
		withRequestParamsKV("deptId", strconv.Itoa(deptId)),
	)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseDecrypt())
	return
}

func (c *Client) GetUserEnableState(ctx context.Context, userId string) (response UserEnableStateResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/cgi/user/enable/state",
		withRequestAccessToken(),
		withRequestEncrypt(),
		withRequestParamsKV("userId", userId),
	)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseDecrypt())
	return
}

func (c *Client) UpdateUserEnableState(ctx context.Context, request UpdateUserEnableStateRequest) (response Response, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/cgi/user/enable/stateupdate",
		withRequestAccessToken(),
		withRequestEncrypt(),
		withRequestBody(request),
	)

	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}
