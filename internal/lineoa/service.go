package lineoa

import (
	"context"
	"fmt"
	"strings"
)

type Service struct {
	client   *Client
	adminTo  string
}

type ServiceOptions struct {
	Client  *Client
	AdminTo string
}

func NewService(opt ServiceOptions) *Service {
	return &Service{client: opt.Client, adminTo: opt.AdminTo}
}

// HandleWebhookEvent: ตรงนี้คือที่คุณใส่ flow จริง
func (s *Service) HandleWebhookEvent(ctx context.Context, ev WebhookEvent) error {
	// event type: message, follow, postback, join, ...
	switch ev.Type {
	case "message":
		if ev.Message == nil {
			return nil
		}
		if ev.Message.Type == "text" {
			return s.handleText(ctx, ev)
		}
		return s.client.ReplyText(ctx, ev.ReplyToken, "รองรับเฉพาะข้อความตัวอักษรตอนนี้ครับ")
	case "follow":
		// ผู้ใช้ add friend
		return s.client.ReplyText(ctx, ev.ReplyToken, "สวัสดีครับ ✅ เพิ่มเพื่อนเรียบร้อย")
	case "join":
		return s.client.ReplyText(ctx, ev.ReplyToken, "ขอบคุณที่เชิญเข้ากลุ่มครับ")
	case "postback":
		return s.client.ReplyText(ctx, ev.ReplyToken, "รับ postback: "+safe(ev.Postback))
	default:
		// ignore events ที่ไม่ใช้
		return nil
	}
}

func (s *Service) handleText(ctx context.Context, ev WebhookEvent) error {
	msg := strings.TrimSpace(ev.Message.Text)

	// ตัวอย่างคำสั่ง
	switch {
	case strings.EqualFold(msg, "ping"):
		return s.client.ReplyText(ctx, ev.ReplyToken, "pong ✅")

	case strings.HasPrefix(strings.ToLower(msg), "sos"):
		// Reply user
		_ = s.client.ReplyText(ctx, ev.ReplyToken, "รับเรื่อง SOS แล้ว ✅ กำลังแจ้งผู้ดูแล")

		// Notify admin (push) ถ้ามี LINE_ADMIN_TO
		if s.adminTo != "" {
			src := describeSource(ev.Source)
			text := fmt.Sprintf("🚨 SOS แจ้งเตือน\nจาก: %s\nข้อความ: %s", src, msg)
			return s.client.PushText(ctx, s.adminTo, text)
		}
		return nil

	default:
		return s.client.ReplyText(ctx, ev.ReplyToken, "พิมพ์ 'ping' หรือ 'sos ...' เพื่อทดสอบครับ")
	}
}

func describeSource(s Source) string {
	if s.Type == "group" && s.GroupID != "" {
		return "group:" + s.GroupID
	}
	if s.Type == "room" && s.RoomID != "" {
		return "room:" + s.RoomID
	}
	if s.UserID != "" {
		return "user:" + s.UserID
	}
	return s.Type
}

func safe(p *Postback) string {
	if p == nil {
		return ""
	}
	return p.Data
}
