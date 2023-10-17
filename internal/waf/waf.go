package waf

import (
	"github.com/spf13/viper"
	"modsecurity-auditlog-agent/internal/storage"
	"modsecurity-auditlog-agent/internal/model"
	"strings"
	"encoding/json"
)

const IP_TYPE_WHITELIST int = 1
const IP_TYPE_BLACKLIST int = 2

const RDS_KEY_IP_WHITELIST string = "waf:ip:whitelist"
const RDS_KEY_IP_BLACKLIST string = "waf:ip:blacklist"
const RDS_KEY_PREFIX_STATS_IP_SCAN_404 string = "waf:stats:scan404:"

func AddIp(ipType int, ip string) {
	var rdsKey string
	if ipType == IP_TYPE_BLACKLIST {
		rdsKey = RDS_KEY_IP_BLACKLIST
	} else {
		rdsKey = RDS_KEY_IP_WHITELIST
	}

	ipArr := strings.Split(ip, ",")

	storage.SaveIpListToRedis(rdsKey, ipArr)
}

func DelIp(ipType int, ip string) {
	var rdsKey string
	if ipType == IP_TYPE_BLACKLIST {
		rdsKey = RDS_KEY_IP_BLACKLIST
	} else {
		rdsKey = RDS_KEY_IP_WHITELIST
	}

	ipArr := strings.Split(ip, ",")

	storage.DelIpListFromRedis(rdsKey, ipArr)
}

func Auditlog(log []byte) {
	// save to es
	index := viper.GetString("elasticsearch.audit-log-index")
	storage.SaveToEs(log, index)

	var auditlog model.Auditlog
	if err := json.Unmarshal(log, &auditlog); err != nil {
		panic(err)
	}
	if auditlog.Transaction.Response.HTTPCode == 404 {
		banByScan404(auditlog)
	}
}
