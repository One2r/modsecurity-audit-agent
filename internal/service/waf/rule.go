package waf

import (
	"context"

	redis "github.com/redis/go-redis/v9"
	"github.com/spf13/viper"

	"modsecurity-auditlog-agent/internal/service/storage"
	"modsecurity-auditlog-agent/internal/model"
	"modsecurity-auditlog-agent/internal/constant"
)

func banByScan404(auditlog model.Auditlog) {
	rdsKey := constant.RDS_KEY_PREFIX_STATS_IP_SCAN_404 + auditlog.Transaction.ClientIP
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})

	incrBy := rdb.IncrBy(ctx, rdsKey, 1)
	if incrBy.Err() != nil {
		panic(incrBy.Err())
	}

	if incrBy.Val() > viper.GetInt64("waf.ban-ip-if-scan-404-url") {
		ipArr := []string{auditlog.Transaction.ClientIP}
		storage.SaveIpListToRedis(constant.RDS_KEY_IP_BLACKLIST, ipArr)
	}

	defer rdb.Close()
}
