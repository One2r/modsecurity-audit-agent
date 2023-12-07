package constant

const (
	RDS_KEY_PREFIX = "waf:"
)

const (
	IP_TYPE_WHITELIST                int    = 1
	IP_TYPE_BLACKLIST                int    = 2
	RDS_KEY_IP_WHITELIST             string = RDS_KEY_PREFIX + "ip:whitelist"
	RDS_KEY_IP_BLACKLIST             string = RDS_KEY_PREFIX + "ip:blacklist"
	RDS_KEY_PREFIX_STATS_IP_SCAN_404 string = RDS_KEY_PREFIX + "stats:scan404:"
)

const (
	ALERT_TYPE_WEBHOOK             = "webhook"
	ALERT_WEBHOOK_TYPE_DINGTALK    = "dingtalk"
	ALERT_WEBHOOK_TYPE_WORKWX      = "workwx"
	RDS_KEY_PREFIX_LAST_ALERT_TIME = RDS_KEY_PREFIX + "alert:last:time"

	ALERT_MARKDOWN_MESSAGE_TPL = `
# 【警  告】疑似受到 web 攻击，请相关人员注意。  

### 上次攻击行为   
> 来源IP：{{.ClientIp}}   
> 服务：{{.Host}}   
> URI: {{.URI}}   
> 时间：{{.TimeStamp}}   
> 触发规则：{{.RuleId}}   
> 规则说明：{{.Match}}   

### 处理   
> HTTP 返回：{{.HTTPCode}}   
> 添加 IP 黑名单：{{.IsBanIp}}   
`
)
