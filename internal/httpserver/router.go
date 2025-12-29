package httpserver

import "github.com/gin-gonic/gin"

func NewRouter(h *Handlers) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	api := r.Group("/api")
	{
		api.GET("/health", h.Health)

		// Webhook endpoint (LINE จะยิงเข้า)
		api.POST("/line/webhook", h.LineWebhook)

		// Admin API (ระบบคุณยิงเองเพื่อแจ้งเตือน)
		api.POST("/notify/admin", h.NotifyAdmin)
		api.POST("/notify/push", h.PushTo) // push ไป userId/groupId ใดๆ
	}

	return r
}
