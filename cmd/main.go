package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"modsecurity-auditlog-agent/internal/waf"
	"net/http"
)

type Form struct {
	IpType int    `form:"ip_type" binding:"required"`
	Ip      string `form:"ip" binding:"required"`
}

func main() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error config file: %s", err)
	}

	if !viper.GetBool("app.debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.POST("/waf/audit-log", func(c *gin.Context) {
		reqBody, err := c.GetRawData()
		if err == nil && len(reqBody) != 0 {
			go waf.Auditlog(reqBody)
		}

		c.JSON(200, gin.H{
			"code": "0",
			"data": "",
		})
	})

	r.POST("/waf/404-log", func(c *gin.Context) {
		reqBody, err := c.GetRawData()
		if err == nil && len(reqBody) != 0 {
			go waf.Notfoundlog(reqBody)
		}

		c.JSON(200, gin.H{
			"code": "0",
			"data": "",
		})
	})

	r.POST("/waf/ip", func(c *gin.Context) {
		var form Form
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		waf.AddIp(form.IpType, form.Ip)

		c.JSON(200, gin.H{
			"code": "0",
			"data": "",
		})
	})

	r.DELETE("/waf/ip", func(c *gin.Context) {
		var form Form
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		waf.DelIp(form.IpType, form.Ip)

		c.JSON(200, gin.H{
			"code": "0",
			"data": "",
		})
	})

	r.Run(":" + viper.GetString("app.port"))
}
