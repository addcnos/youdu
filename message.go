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

type InterfaceMessageRequest interface{}

var (
	_ InterfaceMessageRequest = MessageRequest{}
	_ InterfaceMessageRequest = TextMessageRequest{}
	_ InterfaceMessageRequest = ImageMessageRequest{}
	_ InterfaceMessageRequest = FileMessageRequest{}
	_ InterfaceMessageRequest = MpNewsMessageRequest{}
	_ InterfaceMessageRequest = LinkMessageRequest{}
	_ InterfaceMessageRequest = ExLinkMessageRequest{}
)

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

type TextMessageRequest struct {
	ToUser  string      `json:"toUser"`
	ToDept  string      `json:"toDept"`
	MsgType MsgType     `json:"msgType"`
	Text    MessageText `json:"text"`
}

type ImageMessageRequest struct {
	ToUser  string       `json:"toUser"`
	ToDept  string       `json:"toDept"`
	MsgType MsgType      `json:"msgType"`
	Image   MessageMedia `json:"image"`
}

type FileMessageRequest struct {
	ToUser  string       `json:"toUser"`
	ToDept  string       `json:"toDept"`
	MsgType MsgType      `json:"msgType"`
	File    MessageMedia `json:"file"`
}

type MpNewsMessageRequest struct {
	ToUser  string          `json:"toUser"`
	ToDept  string          `json:"toDept"`
	MsgType MsgType         `json:"msgType"`
	MpNews  []MessageMpNews `json:"mpnews"`
}

type LinkMessageRequest struct {
	ToUser  string      `json:"toUser"`
	ToDept  string      `json:"toDept"`
	MsgType MsgType     `json:"msgType"`
	Link    MessageLink `json:"link"`
}

type ExLinkMessageRequest struct {
	ToUser  string          `json:"toUser"`
	ToDept  string          `json:"toDept"`
	MsgType MsgType         `json:"msgType"`
	ExLink  []MessageExLink `json:"exlink"`
}

type MessageResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (c *Client) SendMessage(ctx context.Context, request InterfaceMessageRequest) (response Response, err error) {
	req, err := c.newRequest(ctx, "POST", "/cgi/msg/send",
		withRequestBody(request), withRequestAccessToken(), withRequestEncrypt())
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

func (c *Client) SendTextMessage(ctx context.Context, request TextMessageRequest) (response Response, err error) {
	request.MsgType = MsgTypeText
	return c.SendMessage(ctx, request)
}

func (c *Client) SendImageMessage(ctx context.Context, request ImageMessageRequest) (response Response, err error) {
	request.MsgType = MsgTypeImage
	return c.SendMessage(ctx, request)
}

func (c *Client) SendFileMessage(ctx context.Context, request FileMessageRequest) (response Response, err error) {
	request.MsgType = MsgTypeFile
	return c.SendMessage(ctx, request)
}

func (c *Client) SendMpNewsMessage(ctx context.Context, request MpNewsMessageRequest) (response Response, err error) {
	request.MsgType = MsgTypeMpNews
	return c.SendMessage(ctx, request)
}

func (c *Client) SendLinkMessage(ctx context.Context, request LinkMessageRequest) (response Response, err error) {
	request.MsgType = MsgTypeLink
	return c.SendMessage(ctx, request)
}

func (c *Client) SendExLinkMessage(ctx context.Context, request ExLinkMessageRequest) (response Response, err error) {
	request.MsgType = MsgTypeExLink
	return c.SendMessage(ctx, request)
}
