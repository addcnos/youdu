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
	MsgTypeSysMsg MsgType = "sysMsg"
	MsgTypePopMsg MsgType = "popMsg"
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

type MessageSysMessageToAll struct {
	OnliyOnline bool `json:"onliyOnline"`
}

type MessageSysMessageSysMsgMsg struct {
	Text MessageText `json:"text,omitempty"`
	Link MessageLink `json:"link,omitempty"`
}

type MessageSysMessageSysMsg struct {
	Title       string                       `json:"title"`
	PopDuration int                          `json:"popDuration,omitempty"`
	Msg         []MessageSysMessageSysMsgMsg `json:"msg"`
}

type MessagePopWindow struct {
	Title    string `json:"title"`
	Url      string `json:"url,omitempty"`
	Tip      string `json:"tip,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
	Duration int    `json:"duration,omitempty"`
	Position int    `json:"position,omitempty"`
	NoticeId string `json:"notice_id"`
	PopMode  int    `json:"pop_mode,omitempty"`
}
