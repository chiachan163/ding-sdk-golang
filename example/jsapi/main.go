package main

import (
	"log"
	"time"

	"github.com/chiachan163/ding-sdk-golang/v2/dingtalk"

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
	suiteAccessToken, err := dingtalk.GetSuiteToken("xxx", "xxx", "")
	if err != nil {
		log.Fatalf("ding_sdk_golang.GetSuiteToken error, err : %v", err)
	}
	log.Println("suite_access_token: ", suiteAccessToken)
	accessToken, err := dingtalk.GetCorpToken(CORPID, suiteAccessToken.SuiteAccessToken)
	if err != nil {
		log.Fatalf("ding_sdk_golang.GetCorpToken error, err : %v", err)
	}
	log.Println("access_token: ", accessToken)

	// 获取jsapi ticket
	ticket, err := dingtalk.GetJsapiTicket(accessToken.AccessToken)
	if err != nil {
		log.Fatalf("ding_sdk_golang.GetCorpToken error, err : %v", err)
	}
	log.Println(*ticket)
	//log.Println("ticket: ", fmt.Sprintf("%p", unsafe.Pointer(ticket)))
}
