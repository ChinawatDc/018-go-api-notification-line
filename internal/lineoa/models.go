package lineoa

// ===== Webhook (รับจาก LINE) =====

type WebhookRequest struct {
	Destination string         `json:"destination"`
	Events      []WebhookEvent `json:"events"`
}

type WebhookEvent struct {
	Type       string `json:"type"`
	ReplyToken string `json:"replyToken"`
	Mode       string `json:"mode"`
	Timestamp  int64  `json:"timestamp"`
	Source     Source `json:"source"`
	Message    *Message `json:"message,omitempty"`
	Postback   *Postback `json:"postback,omitempty"`
}

type Source struct {
	Type    string `json:"type"` // user|group|room
	UserID  string `json:"userId,omitempty"`
	GroupID string `json:"groupId,omitempty"`
	RoomID  string `json:"roomId,omitempty"`
}

type Message struct {
	ID   string `json:"id"`
	Type string `json:"type"` // text, image, ...
	Text string `json:"text,omitempty"`
}

type Postback struct {
	Data string `json:"data"`
}

// ===== ส่งออกไปที่ LINE =====

type TextMessage struct {
	Type string `json:"type"` // "text"
	Text string `json:"text"`
}

type ReplyMessageRequest struct {
	ReplyToken string        `json:"replyToken"`
	Messages   []TextMessage `json:"messages"`
}

type PushMessageRequest struct {
	To       string        `json:"to"`
	Messages []TextMessage `json:"messages"`
}
