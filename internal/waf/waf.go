package waf

import (
	"github.com/spf13/viper"
	"modsecurity-auditlog-agent/internal/storage"
	"strings"
)

func AddIp(biz_type int, ip string) {
	var rdsKek = "waf:ip:whitelist"
	if biz_type == 2 {
		rdsKek = "waf:ip:blacklist"
	}

	arr := strings.Split(ip, ",")

	storage.SaveIpListToRedis(rdsKek, arr)
}

func DelIp(biz_type int, ip string) {
	var rdsKek = "waf:ip:whitelist"
	if biz_type == 2 {
		rdsKek = "waf:ip:blacklist"
	}

	arr := strings.Split(ip, ",")

	storage.DelIpListToRedis(rdsKek, arr)
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
