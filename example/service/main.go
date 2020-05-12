package main

import (
	"fmt"
	"log"
	"time"

	ding_sdk_golang "github.com/chiachan163/ding-sdk-golang/v1"
	"github.com/chiachan163/ding-sdk-golang/v1/arg"
	"github.com/chiachan163/ding-sdk-golang/v1/dingtalk"

	ding_redis "github.com/chiachan163/ding-sdk-golang/v1/redis"
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
	//const CORPID = "xxx"
	//suiteAccessToken, err := dingtalk.GetSuiteToken("xxx", "xxx", "")
	//if err != nil {
	//	log.Fatalf("ding_sdk_golang.GetSuiteToken error, err : %v", err)
	//}
	//log.Println("suite_access_token: ", suiteAccessToken)
	//accessToken, err := dingtalk.GetCorpToken(CORPID, suiteAccessToken)
	//if err != nil {
	//	log.Fatalf("ding_sdk_golang.GetCorpToken error, err : %v", err)
	//}
	//log.Println("access_token: ", accessToken)
	//
	//// 免登授权码
	//code := "f4196a7cd6e73908ab010ee5bc3b379a"
	//userinfo, err := dingtalk.Getuserinfo(accessToken, code)
	//if err != nil {
	//	log.Fatalf("ding_sdk_golang.GetCorpToken error, err : %v", err)
	//}
	//log.Println(userinfo)
	//user, err := dingtalk.GetUser(accessToken, userinfo.Userid)
	//if err != nil {
	//	log.Fatalf("ding_sdk_golang.GetCorpToken error, err : %v", err)
	//}
	//log.Println(user.Name)
	const CORPID = "xxx"

	var authInfo arg.GetAuthInfoResult
	err := ding_sdk_golang.CallDingTalk(CORPID, arg.GETSUITETOKENURLID, func(corpId string, apiPath int) (*ding_sdk_golang.RespResult, error) {
		result, err := dingtalk.GetSuiteToken("suitetjww5szhquqyi5ey", "CeE8vanmB_3j1-tVz1immRUEmkA6KTlqieoroX4aQPyWYc32MXBkhDcw5k1T3HlB", "")
		if err != nil {
			log.Fatalf("ding_sdk_golang.GetSuiteToken error, err : %v", err)
		}
		log.Println("suite_access_token: ", result.SuiteAccessToken)

		err = ding_sdk_golang.CallDingTalk(CORPID, arg.SERVICEGETAUTHINFOID, func(corpId string, apiPath int) (*ding_sdk_golang.RespResult, error) {
			result2, err := dingtalk.GetAuthInfo("suitetjww5szhquqyi5ey", corpId, result.SuiteAccessToken)
			if err != nil {
				log.Fatalf("ding_sdk_golang.GetSuiteToken error, err : %v", err)
			}
			log.Println("authInfo: ", result)
			authInfo = *result2
			return &ding_sdk_golang.RespResult{
				Errcode: result.Errcode,
				Errmsg:  result.Errmsg,
			}, nil
		})

		return &ding_sdk_golang.RespResult{
			Errcode: result.Errcode,
			Errmsg:  result.Errmsg,
		}, nil

	})

	if err != nil {
		log.Fatal(err)
	}
	for _, _agent := range authInfo.AuthInfo.Agent {
		fmt.Println(_agent)
	}
}
