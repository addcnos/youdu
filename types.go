package youdu

type MsgType string

const (
	MsgTypeText   MsgType = "text"
	MsgTypeImage  MsgType = "image"
	MsgTypeFile   MsgType = "file"
	MsgTypeMpNews MsgType = "mpnews"
	MsgTypeLink   MsgType = "link"
	MsgTypeExLink MsgType = "exlink"
	MsgTypeVoice  MsgType = "voice"
	MsgTypeVideo  MsgType = "video"
)

type MessageText struct {
	Content string `json:"content"`
}

type MessageMedia struct {
	MediaId string `json:"media_id"`
}

type MessageFile struct {
	MediaId string `json:"media_id"`
	Name    string `json:"name"`
	Size    int64  `json:"size"`
}
