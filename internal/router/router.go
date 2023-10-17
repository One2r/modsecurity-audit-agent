package router

import (
	"github.com/gin-gonic/gin"

	"modsecurity-auditlog-agent/internal/controller/api"
	"modsecurity-auditlog-agent/internal/middleware"
)

func SetRouter(r *gin.Engine) {

	r.POST("/waf/audit-log", middleware.ApiAuth(), api.AuditlogHandler)

	// ip router
	r.POST("/waf/ip", middleware.ApiAuth(), api.AddIpHandler)
	r.DELETE("/waf/ip", middleware.ApiAuth(), api.DelIpHandler)
}
