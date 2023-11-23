package storage

import (
	"bytes"
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"time"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	redis "github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func SaveToEs(data []byte) {

	esurl := viper.GetString("elasticsearch.url")
	if esurl == "" {
		return
	}

	cfg := elasticsearch.Config{
		Addresses: []string{
			esurl,
		},
		Username: viper.GetString("elasticsearch.username"),
		Password: viper.GetString("elasticsearch.password"),
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	index := viper.GetString("elasticsearch.audit-log-index")
	req := esapi.IndexRequest{
		Index: index,
		Body:  bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
}

func SaveIpListToRedis(key string, ipList []string) {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})

	for _, ip := range ipList {
		_, err := rdb.Do(ctx, "CF.ADDNX", key, ip).Bool()
		if err != nil {
			panic(err)
		}
	}
	defer rdb.Close()
}

func DelIpListFromRedis(key string, ipList []string) {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})

	for _, ip := range ipList {
		_, err := rdb.Do(ctx, "CF.DEL", key, ip).Bool()
		if err != nil {
			panic(err)
		}
	}

	defer rdb.Close()
}
