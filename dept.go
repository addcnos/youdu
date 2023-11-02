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
