package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"modsecurity-auditlog-agent/internal/service/waf"
	"modsecurity-auditlog-agent/internal/model/request"
)

func AuditlogHandler(c *gin.Context) {
	reqBody, err := c.GetRawData()
	if err == nil && len(reqBody) != 0 {
		go waf.Auditlog(reqBody)
	}

	c.JSON(200, gin.H{
		"code": "0",
		"data": "",
	})
}

func AddIpHandler(c *gin.Context) {
	var form request.IpReq
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	waf.AddIp(form.IpType, form.Ip)

	c.JSON(200, gin.H{
		"code": "0",
		"data": "",
	})
}

func DelIpHandler(c *gin.Context) {
	var form request.IpReq
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	waf.DelIp(form.IpType, form.Ip)

	c.JSON(200, gin.H{
		"code": "0",
		"data": "",
	})
}
