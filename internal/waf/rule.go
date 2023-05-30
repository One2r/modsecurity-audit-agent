package waf

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"modsecurity-auditlog-agent/internal/storage"
	"modsecurity-auditlog-agent/models"
)

func banByScan404(data []byte) {
	var log models.NotFoundLog
	if err := json.Unmarshal(data, &log); err != nil {
		panic(err)
	}

	rdsKey := RDS_KEY_PREFIX_STATS_IP_SCAN_404 + log.ClientIP
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
		ipArr := []string{log.ClientIP}
		storage.SaveIpListToRedis(RDS_KEY_IP_BLACKLIST, ipArr)
	}

	defer rdb.Close()
}
