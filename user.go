package youdu

import (
	"context"
	"net/http"
	"strconv"
)

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

type UserList struct {
	UserId     string       `json:"userId"`
	Name       string       `json:"name"`
	Gender     interface{}  `json:"gender"`
	Mobile     string       `json:"mobile"`
	Phone      string       `json:"phone"`
	Email      string       `json:"email"`
	Dept       []int        `json:"dept"`
	DeptDetail []DeptDetail `json:"deptDetail"`
}

type DeptUserListResponse struct {
	UserList []UserList `json:"userList"`
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