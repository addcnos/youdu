package youdu

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/go-addcnos/youdu/message"
)

const (
	msgSendUrl   = "/cgi/msg/send"
	popWindowUrl = "/cgi/popwindow"
)

type Message struct {
	config *Config
}

func NewMessage(config *Config) *Message {
	return &Message{
		config: config,
	}
}

func (m *Message) Send(msg message.Message) error {
	accessToken, err := m.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return err
	}

	msgJson, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	encrypt, err := m.config.GetEncryptor().Encrypt(string(msgJson))
	if err != nil {
		return err
	}

	resp, err := m.config.GetHttp().Post(msgSendUrl+"?accessToken="+accessToken, map[string]interface{}{
		"appId":   m.config.AppId,
		"buin":    m.config.Buin,
		"encrypt": encrypt,
	})
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return errors.New("Response status code is " + strconv.Itoa(resp.StatusCode()))
	}

	jsonRet, err := resp.Json()
	if err != nil {
		return err
	}

	if jsonRet["errcode"].(float64) != 0 {
		return errors.New(jsonRet["errmsg"].(string))
	}

	return nil
}

func (m *Message) SendText(toUser, content string, toDept ...string) error {
	if len(toDept) == 0 {
		toDept = []string{""}
	}

	return m.Send(&message.TextMessage{
		ToUser:  toUser,
		ToDept:  toDept[0],
		MsgType: message.MsgTypeText,
		Text: &message.TextItem{
			Content: content,
		},
	})
}

func (m *Message) SendImage(toUser, mediaId string, toDept ...string) error {
	if len(toDept) == 0 {
		toDept = []string{""}
	}

	return m.Send(&message.ImageMessage{
		ToUser:  toUser,
		ToDept:  toDept[0],
		MsgType: message.MsgTypeImage,
		Image: &message.MediaItem{
			MediaId: mediaId,
		},
	})
}

func (m *Message) SendFile(toUser, mediaId string, toDept ...string) error {
	if len(toDept) == 0 {
		toDept = []string{""}
	}

	return m.Send(&message.FileMessage{
		ToUser:  toUser,
		ToDept:  toDept[0],
		MsgType: message.MsgTypeFile,
		File: &message.MediaItem{
			MediaId: mediaId,
		},
	})
}

func (m *Message) Popwindow(msg message.Message) error {
	if _, ok := msg.(*message.PopWindowMessage); !ok {
		return errors.New("message must be PopWindowMessage")
	}

	msgJson, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	encrypt, err := m.config.GetEncryptor().Encrypt(string(msgJson))
	if err != nil {
		return err
	}

	resp, err := m.config.GetHttp().Post(popWindowUrl, map[string]interface{}{
		"app_id":      m.config.AppId,
		"msg_encrypt": encrypt,
	})
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return errors.New("Response status code is " + strconv.Itoa(resp.StatusCode()))
	}

	jsonRet, err := resp.Json()
	if err != nil {
		return err
	}

	fmt.Println(jsonRet)

	if jsonRet["errcode"].(float64) != 0 {
		return errors.New(jsonRet["errmsg"].(string))
	}

	return nil
}
