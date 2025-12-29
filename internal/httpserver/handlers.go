package httpserver

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ChinawatDc/018-go-api-notification-line/internal/lineoa"
)

type Handlers struct {
	ChannelSecret string
	LineSvc       *lineoa.Service
	LineClient    *lineoa.Client
	AdminTo       string
}

func NewHandlers(secret string, svc *lineoa.Service, client *lineoa.Client, adminTo string) *Handlers {
	return &Handlers{
		ChannelSecret: secret,
		LineSvc:       svc,
		LineClient:    client,
		AdminTo:       adminTo,
	}
}

func (h *Handlers) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (h *Handlers) LineWebhook(c *gin.Context) {
	// ต้องอ่าน raw body เพื่อ verify signature ก่อน (ห้ามแก้ไขก่อน verify) :contentReference[oaicite:5]{index=5}
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": err.Error()})
		return
	}

	sign := c.GetHeader("X-Line-Signature")
	if sign == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "missing X-Line-Signature"})
		return
	}

	if !lineoa.VerifySignature(h.ChannelSecret, body, sign) {
		c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "invalid signature"})
		return
	}

	var wh lineoa.WebhookRequest
	if err := json.Unmarshal(body, &wh); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "invalid json"})
		return
	}

	// handle events
	for _, ev := range wh.Events {
		_ = h.LineSvc.HandleWebhookEvent(c.Request.Context(), ev)
	}

	// LINE ต้องการ 200 OK
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

// ===== Admin APIs =====

type notifyAdminDTO struct {
	Message string `json:"message"`
}

func (h *Handlers) NotifyAdmin(c *gin.Context) {
	if h.AdminTo == "" {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "LINE_ADMIN_TO not set"})
		return
	}
	var dto notifyAdminDTO
	if err := c.ShouldBindJSON(&dto); err != nil || strings.TrimSpace(dto.Message) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "message is required"})
		return
	}
	if err := h.LineClient.PushText(c.Request.Context(), h.AdminTo, dto.Message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

type pushDTO struct {
	To      string `json:"to"`      // userId/groupId/roomId
	Message string `json:"message"` // text
}

func (h *Handlers) PushTo(c *gin.Context) {
	var dto pushDTO
	if err := c.ShouldBindJSON(&dto); err != nil || strings.TrimSpace(dto.To) == "" || strings.TrimSpace(dto.Message) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "to and message are required"})
		return
	}
	if err := h.LineClient.PushText(c.Request.Context(), dto.To, dto.Message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
}
