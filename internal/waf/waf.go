package waf

import (
	"github.com/spf13/viper"
	"modsecurity-auditlog-agent/internal/storage"
	"strings"
)

const IP_TYPE_WHITELIST int = 1
const IP_TYPE_BLACKLIST int = 2

func AddIp(ipType int, ip string) {
	rdsKey := "waf:ip:whitelist"
	if ipType == IP_TYPE_BLACKLIST {
		rdsKey = "waf:ip:blacklist"
	}

	ipArr := strings.Split(ip, ",")

	storage.SaveIpListToRedis(rdsKey, ipArr)
}

func DelIp(ipType int, ip string) {
	rdsKey := "waf:ip:whitelist"
	if ipType == IP_TYPE_BLACKLIST {
		rdsKey = "waf:ip:blacklist"
	}

	ipArr := strings.Split(ip, ",")

	storage.DelIpListToRedis(rdsKey, ipArr)
}

func Auditlog(log []byte) {
	// save to es
	index := viper.GetString("elasticsearch.audit-log-index")
	storage.SaveToEs(log, index)
}

func Notfoundlog(log []byte) {
	// save to es
	index := viper.GetString("elasticsearch.waf-404-log")
	storage.SaveToEs(log, index)
}
