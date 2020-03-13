package main

import (
	"log"
	"time"

	ding_sdk_golang "github.com/chiachan163/ding-sdk-golang"
	ding_redis "github.com/chiachan163/ding-sdk-golang/redis"
	"github.com/xiaoenai/tp-micro/v6/model/redis"
)

func init() {
	config := &redis.Config{
		DeployType: "single",
		ForSingle: redis.SingleConfig{
			Addr: "127.0.0.1:6379",
		},
		ForCluster:      redis.ClusterConfig{},
		MaxRetries:      0,
		DialTimeout:     0,
		PoolSizePerNode: 0,
		IdleTimeout:     0,
	}
	ding_redis.Init(config, time.Hour*24)
}

func main() {
	const CORPID = "xxx"

	ding_sdk_golang.CallDingTalk(CORPID, ding_sdk_golang.GETSUITETOKENURL, func(corpId string, apiPath string) error {
		suiteAccessToken, err := ding_sdk_golang.GetSuiteToken("xxx", "xxx", "")
		if err != nil {
			log.Fatalf("ding_sdk_golang.GetSuiteToken error, err : %v", err)
		}
		log.Println("suite_access_token: ", suiteAccessToken)
		return nil
	})
}
