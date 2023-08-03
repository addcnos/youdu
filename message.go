package youdu

import "context"

type MsgType string

const (
	MsgTypeText   MsgType = "text"
	MsgTypeImage  MsgType = "image"
	MsgTypeFile   MsgType = "file"
	MsgTypeMpNews MsgType = "mpnews"
	MsgTypeLink   MsgType = "link"
	MsgTypeExLink MsgType = "exlink"
)

type MessageText struct {
	Content string `json:"content"`
}

type MessageMedia struct {
	MediaId string `json:"media_id"`
}

type MessageMpNews struct {
	Title     string `json:"title"`
	MediaId   string `json:"media_id"`
	Content   string `json:"content"`
	Digest    string `json:"digest,omitempty"`
	ShowFront int    `json:"showFront,omitempty"`
}

type MessageLink struct {
	Title  string `json:"title"`
	Url    string `json:"url"`
	Action int    `json:"action,omitempty"`
}

type MessageExLink struct {
	Title   string `json:"title"`
	Url     string `json:"url"`
	MediaId string `json:"media_id"`
	Digest  string `json:"digest,omitempty"`
}

type MessageRequest struct {
	ToUser  string          `json:"toUser"`
	ToDept  string          `json:"toDept"`
	MsgType MsgType         `json:"msgType"`
	Text    MessageText     `json:"text,omitempty"`
	Image   MessageMedia    `json:"image,omitempty"`
	File    MessageMedia    `json:"file,omitempty"`
	MpNews  []MessageMpNews `json:"mpnews,omitempty"`
	Link    MessageLink     `json:"link,omitempty"`
	ExLink  []MessageExLink `json:"exlink,omitempty"`
}

type MessageResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (c *Client) SendMessage(ctx context.Context, request MessageRequest) (response MessageResponse, err error) {
	req, err := c.newRequest(ctx, "POST", "/cgi/msg/send",
		withRequestBody(request), withRequestAccessToken(), withRequestEncrypt())
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}
