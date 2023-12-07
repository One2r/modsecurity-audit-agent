package alert

import (
	"time"
	"encoding/json"

	"go.uber.org/zap"
	"github.com/spf13/viper"
	resty "github.com/go-resty/resty/v2"

	"modsecurity-auditlog-agent/internal/constant"
	"modsecurity-auditlog-agent/internal/model/request"
)

func webhook(message string) {
	webhookType := viper.GetString("alert.webhook_type")
	switch webhookType {
	case constant.ALERT_WEBHOOK_TYPE_DINGTALK:
		dingtalk(message)
	case constant.ALERT_WEBHOOK_TYPE_WORKWX:
		workwx(message)
	default:
		return
	}
}

func dingtalk(message string) {
	client := resty.New()
	client.SetTimeout(15 * time.Second)

	webhookUrl := viper.GetString("alert.webhook_url")
	dingtalkMessage := request.DingtalkMessage{
		Msgtype: "markdown",
	}
	dingtalkMessage.Markdown.Title = "test"
	dingtalkMessage.Markdown.Text = message

	msg, err := json.Marshal(dingtalkMessage)
	if err != nil {
		zap.S().Warnf("json encode error:", err)
		return
	}

	_, respErr := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(msg).
		Post(webhookUrl)
	if respErr != nil {
		zap.S().Warnf("call dingtalk webhook error:", respErr)
		return
	}
}

func workwx(message string) {
	client := resty.New()
	client.SetTimeout(15 * time.Second)

	webhookUrl := viper.GetString("alert.webhook_url")
	workwxMessage := request.WorkwxMessage{
		Msgtype: "markdown",
	}
	workwxMessage.Markdown.Content = message

	msg, err := json.Marshal(workwxMessage)
	if err != nil {
		zap.S().Warnf("json encode error:", err)
		return
	}

	_, respErr := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(msg).
		Post(webhookUrl)
	if respErr != nil {
		zap.S().Warnf("call workwx webhook error:", respErr)
		return
	}
}
