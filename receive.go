package youdu

import (
	"encoding/json"
)

var (
	ErrReceiveEmptyEncrypt = newError(-1, "encrypt is empty")
	ErrReceiveEmptyData    = newError(-1, "data is empty")
)

type ReceiveRequest struct {
	ToBuin  int    `json:"toBuin"`
	ToApp   string `json:"toApp"`
	Encrypt string `json:"encrypt"`
}

type ReceiveMessage struct {
	FromUser   string       `json:"fromUser"`
	CreateTime int          `json:"createTime"`
	PackageId  string       `json:"packageId"`
	MsgType    MsgType      `json:"msgType"`
	Text       MessageText  `json:"text,omitempty"`
	Image      MessageMedia `json:"image,omitempty"`
	File       MessageFile  `json:"file,omitempty"`
}

type Receive struct {
	config    *Config
	encryptor *Encryptor
}

func NewReceive(config *Config) *Receive {
	return &Receive{
		config:    config,
		encryptor: NewEncryptorWithConfig(config),
	}
}

func (s *Receive) Decrypt(request ReceiveRequest) (message ReceiveMessage, err error) {
	if request.Encrypt == "" {
		err = ErrReceiveEmptyEncrypt
		return
	}

	rawData, err := s.encryptor.Decrypt(request.Encrypt)
	if err != nil {
		return
	}
	if rawData.Data == nil {
		err = ErrReceiveEmptyData
		return
	}

	err = json.Unmarshal(rawData.Data, &message)
	return
}
