package waf

import (
	"strings"
	"encoding/json"

	"github.com/spf13/viper"

	"modsecurity-auditlog-agent/internal/service/storage"
	"modsecurity-auditlog-agent/internal/model"
	"modsecurity-auditlog-agent/internal/constant"
)

func AddIp(ipType int, ip string) {
	var rdsKey string
	if ipType == constant.IP_TYPE_BLACKLIST {
		rdsKey = constant.RDS_KEY_IP_BLACKLIST
	} else {
		rdsKey = constant.RDS_KEY_IP_WHITELIST
	}

	ipArr := strings.Split(ip, ",")

	storage.SaveIpListToRedis(rdsKey, ipArr)
}

func DelIp(ipType int, ip string) {
	var rdsKey string
	if ipType == constant.IP_TYPE_BLACKLIST {
		rdsKey = constant.RDS_KEY_IP_BLACKLIST
	} else {
		rdsKey = constant.RDS_KEY_IP_WHITELIST
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
