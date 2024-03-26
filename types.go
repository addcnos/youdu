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
)

type MessageText struct {
	Content string `json:"content"`
}

type MessageMedia struct {
	MediaID string `json:"media_id"`
}

type MessageFile struct {
	MediaID string `json:"media_id"`
	Name    string `json:"name"`
	Size    int64  `json:"size"`
}

type MessageMpNews struct {
	Title     string `json:"title"`
	MediaID   string `json:"media_id"`
	Content   string `json:"content"`
	Digest    string `json:"digest,omitempty"`
	ShowFront int    `json:"showFront,omitempty"`
}

type MessageLink struct {
	Title  string `json:"title"`
	URL    string `json:"url"`
	Action int    `json:"action,omitempty"`
}

type MessageExLink struct {
	Title   string `json:"title"`
	URL     string `json:"url"`
	MediaID string `json:"media_id"`
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
	URL      string `json:"url"`
	Tip      string `json:"tip"`
	Title    string `json:"title"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Duration int    `json:"duration"`
	Position int    `json:"position"`
	NoticeID string `json:"notice_id"`
	PopMode  int    `json:"pop_mode"`
}
