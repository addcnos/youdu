package session

type Session struct {
	SessionId  string   `json:"sessionId"`
	Type       string   `json:"type"`
	Owner      string   `json:"owner"`
	Title      string   `json:"title"`
	Version    int      `json:"version"`
	Member     []string `json:"member"`
	LastMsgId  int      `json:"lastMsgId"`
	ActiveTime int      `json:"activeTime"`
}
