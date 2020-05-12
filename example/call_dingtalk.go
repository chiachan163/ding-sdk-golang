package main

import (
	"log"
	"time"

	"github.com/chiachan163/ding-sdk-golang/v2/dingtalk"

	"github.com/chiachan163/ding-sdk-golang/v2/arg"

	ding_sdk_golang "github.com/chiachan163/ding-sdk-golang/v2"
	ding_redis "github.com/chiachan163/ding-sdk-golang/v2/redis"
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

	ding_sdk_golang.CallDingTalk(CORPID, arg.GETCORPTOKENURLID, func(corpId string, apiPath int) (*ding_sdk_golang.RespResult, error) {
		result, err := dingtalk.GetSuiteToken("xxx", "xxx", "")
		if err != nil {
			log.Fatalf("ding_sdk_golang.GetSuiteToken error, err : %v", err)
		}
		log.Println("suite_access_token: ", result.SuiteAccessToken)

		return &ding_sdk_golang.RespResult{
			Errcode: result.Errcode,
			Errmsg:  result.Errmsg,
		}, nil
	})
}
