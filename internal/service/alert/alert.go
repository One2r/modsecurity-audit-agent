package alert

import (
	"time"
	"context"
	"bytes"
	"net/http"
	"text/template"

	"go.uber.org/zap"
	"github.com/spf13/viper"
	redis "github.com/redis/go-redis/v9"

	"modsecurity-auditlog-agent/internal/constant"
	"modsecurity-auditlog-agent/internal/model"
	"modsecurity-auditlog-agent/internal/model/request"
)

func Alert(log model.Auditlog) {
	rdsKey := constant.RDS_KEY_PREFIX_LAST_ALERT_TIME
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
	v, _ := rdb.Exists(ctx, rdsKey).Result()

	if v == 0 {
		alertType := viper.GetString("alert.type")
		if alertType == constant.ALERT_TYPE_WEBHOOK {
			vars := request.MdTplMessage{
				ClientIp:  log.Transaction.ClientIP,
				Host:      log.Transaction.Request.Headers.Host,
				URI:       log.Transaction.Request.URI,
				TimeStamp: log.Transaction.TimeStamp,
				HTTPCode:  log.Transaction.Response.HTTPCode,
				IsBanIp:   "是",
			}
			if len(log.Transaction.Messages) > 0 {
				vars.RuleId = log.Transaction.Messages[0].Details.RuleID
				vars.Match = log.Transaction.Messages[0].Details.Match
			} else {
				if log.Transaction.Response.HTTPCode == http.StatusNotFound {
					vars.RuleId = "-"
					vars.Match = "非法 URL 扫描"
				}
			}

			tmpl, err := template.New("md").Parse(constant.ALERT_MARKDOWN_MESSAGE_TPL)
			if err != nil {
				zap.S().Warnf("read webhook alert tpl error:", err)
				return
			}
			var message bytes.Buffer
			err = tmpl.Execute(&message, vars)
			if err != nil {
				zap.S().Warnf("execute webhook alert tpl error:", err)
				return
			}

			webhook(message.String())
		}

		expire, _ := time.ParseDuration(viper.GetString("alert.time_intervals"))
		rdb.SetEx(ctx, rdsKey, time.Now(), expire)
	}
}
