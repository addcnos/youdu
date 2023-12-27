package youdu

import (
	"context"
	"net/http"
)

type InterfaceMessageRequest interface{}

var (
	_ InterfaceMessageRequest = MessageRequest{}
	_ InterfaceMessageRequest = TextMessageRequest{}
	_ InterfaceMessageRequest = ImageMessageRequest{}
	_ InterfaceMessageRequest = FileMessageRequest{}
	_ InterfaceMessageRequest = MpNewsMessageRequest{}
	_ InterfaceMessageRequest = LinkMessageRequest{}
	_ InterfaceMessageRequest = ExLinkMessageRequest{}
	_ InterfaceMessageRequest = MessageSysMessageRequest{}
)

type MessageRequest struct {
	// General
	ToUser  string  `json:"toUser"`
	ToDept  string  `json:"toDept"`
	MsgType MsgType `json:"msgType"`

	// Text, Image, File, MpNews, Link, ExLink
	Text   MessageText     `json:"text,omitempty"`
	Image  MessageMedia    `json:"image,omitempty"`
	File   MessageMedia    `json:"file,omitempty"`
	MpNews []MessageMpNews `json:"mpnews,omitempty"`
	Link   MessageLink     `json:"link,omitempty"`
	ExLink []MessageExLink `json:"exlink,omitempty"`

	// SysMsg
	ToAll  MessageSysMessageToAll  `json:"toAll,omitempty"`
	SysMsg MessageSysMessageSysMsg `json:"sysMsg"`
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

type MessageSysMessageRequest struct {
	ToUser  string                  `json:"toUser,omitempty"`
	ToDept  string                  `json:"toDept,omitempty"`
	ToAll   MessageSysMessageToAll  `json:"toAll,omitempty"`
	MsgType MsgType                 `json:"msgType"`
	SysMsg  MessageSysMessageSysMsg `json:"sysMsg"`
}

type PopWindowMessageRequest struct {
	ToUser    string           `json:"toUser"`
	ToDept    string           `json:"toDept"`
	PopWindow MessagePopWindow `json:"popWindow"`
}

func (c *Client) SendMessage(ctx context.Context, request InterfaceMessageRequest) (response Response, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/cgi/msg/send",
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

func (c *Client) SendSysMessage(ctx context.Context, request MessageSysMessageRequest) (response Response, err error) {
	request.MsgType = MsgTypeSysMsg
	return c.SendMessage(ctx, request)
}

func (c *Client) SendPopWindowMessage(ctx context.Context, request PopWindowMessageRequest) (response Response, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/cgi/popwindow",
		withRequestBody(request),
		withRequestEncrypt(),
		withRequestType(SpecialRequestType),
	)

	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}
