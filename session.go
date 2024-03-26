package youdu

import (
	"context"
	"net/http"
)

type SessionType string

var SessionMultiType SessionType = "multi"

type CreateSessionRequest struct {
	Title   string      `json:"title"`
	Creator string      `json:"creator"`
	Type    SessionType `json:"type,omitempty"`
	Member  []string    `json:"member"`
}

type UpdateSessionRequest struct {
	SessionID string   `json:"sessionId"`
	OpUser    string   `json:"opUser"`
	Title     string   `json:"title,omitempty"`
	AddMember []string `json:"addMember"`
	DelMember []string `json:"delMember"`
}

type SessionResponse struct {
	SessionID string      `json:"sessionId"`
	Title     string      `json:"title"`
	Owner     string      `json:"owner"`
	Version   int         `json:"version"`
	Type      SessionType `json:"type"`
	Member    []string    `json:"member"`
}

func (c *Client) CreateSession(
	ctx context.Context, request CreateSessionRequest,
) (response SessionResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/cgi/session/create",
		withRequestBody(request), withRequestAccessToken(), withRequestEncrypt())
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseDecrypt())
	return
}

func (c *Client) GetSession(ctx context.Context, sessionID string) (response SessionResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/cgi/session/get",
		withRequestAccessToken(),
		withRequestEncrypt(),
		withRequestParamsKV("sessionId", sessionID),
	)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseDecrypt())
	return
}

func (c *Client) UpdateSession(
	ctx context.Context, request UpdateSessionRequest,
) (response SessionResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/cgi/session/update",
		withRequestBody(request), withRequestAccessToken(), withRequestEncrypt())
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseDecrypt())
	return
}

type InterfaceSessionMessageRequest interface{}

var (
	_ InterfaceSessionMessageRequest = SessionMessageResponse{}
	_ InterfaceSessionMessageRequest = TextSessionMessageRequest{}
	_ InterfaceSessionMessageRequest = ImageSessionMessageRequest{}
	_ InterfaceSessionMessageRequest = FileSessionMessageRequest{}
	_ InterfaceSessionMessageRequest = VoiceSessionMessageRequest{}
	_ InterfaceSessionMessageRequest = VideoSessionMessageRequest{}
)

type SessionMessageResponse struct {
	SessionID string       `json:"sessionId,omitempty"`
	Receiver  string       `json:"receiver,omitempty"`
	Sender    string       `json:"sender"`
	MsgType   MsgType      `json:"msgType"`
	Text      MessageText  `json:"text,omitempty"`
	Image     MessageMedia `json:"image,omitempty"`
	File      MessageFile  `json:"file,omitempty"`
	Voice     MessageMedia `json:"voice,omitempty"`
	Video     MessageMedia `json:"video,omitempty"`
}

type TextSessionMessageRequest struct {
	SessionID string      `json:"sessionId,omitempty"`
	Receiver  string      `json:"receiver,omitempty"`
	Sender    string      `json:"sender"`
	MsgType   MsgType     `json:"msgType"`
	Text      MessageText `json:"text"`
}

type ImageSessionMessageRequest struct {
	SessionID string       `json:"sessionId,omitempty"`
	Receiver  string       `json:"receiver,omitempty"`
	Sender    string       `json:"sender"`
	MsgType   MsgType      `json:"msgType"`
	Image     MessageMedia `json:"image"`
}

type FileSessionMessageRequest struct {
	SessionID string      `json:"sessionId,omitempty"`
	Receiver  string      `json:"receiver,omitempty"`
	Sender    string      `json:"sender"`
	MsgType   MsgType     `json:"msgType"`
	File      MessageFile `json:"file"`
}

type VoiceSessionMessageRequest struct {
	SessionID string       `json:"sessionId,omitempty"`
	Receiver  string       `json:"receiver,omitempty"`
	Sender    string       `json:"sender"`
	MsgType   MsgType      `json:"msgType"`
	Voice     MessageMedia `json:"voice"`
}

type VideoSessionMessageRequest struct {
	SessionID string       `json:"sessionId,omitempty"`
	Receiver  string       `json:"receiver,omitempty"`
	Sender    string       `json:"sender"`
	MsgType   MsgType      `json:"msgType"`
	Video     MessageMedia `json:"video"`
}

func (c *Client) SendSessionMessage(
	ctx context.Context,
	request InterfaceSessionMessageRequest,
) (response Response, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/cgi/session/send",
		withRequestBody(request), withRequestAccessToken(), withRequestEncrypt())
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

func (c *Client) SendTextSessionMessage(
	ctx context.Context, request TextSessionMessageRequest,
) (response Response, err error) {
	request.MsgType = MsgTypeText
	return c.SendSessionMessage(ctx, request)
}

func (c *Client) SendImageSessionMessage(
	ctx context.Context, request ImageSessionMessageRequest,
) (response Response, err error) {
	request.MsgType = MsgTypeImage
	return c.SendSessionMessage(ctx, request)
}

func (c *Client) SendFileSessionMessage(
	ctx context.Context, request FileSessionMessageRequest,
) (response Response, err error) {
	request.MsgType = MsgTypeFile
	return c.SendSessionMessage(ctx, request)
}

func (c *Client) SendVoiceSessionMessage(
	ctx context.Context, request VoiceSessionMessageRequest,
) (response Response, err error) {
	request.MsgType = MsgTypeVoice
	return c.SendSessionMessage(ctx, request)
}

func (c *Client) SendVideoSessionMessage(
	ctx context.Context, request VideoSessionMessageRequest,
) (response Response, err error) {
	request.MsgType = MsgTypeVideo
	return c.SendSessionMessage(ctx, request)
}
