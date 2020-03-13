package main

import (
	"log"

	"github.com/chiachan163/ding-sdk-golang/dingtalk"
)

func main() {
	const CORPID = "xxx"
	suiteAccessToken, err := dingtalk.GetSuiteToken("xxx", "xxx", "")
	if err != nil {
		log.Fatalf("ding_sdk_golang.GetSuiteToken error, err : %v", err)
	}
	log.Println("suite_access_token: ", suiteAccessToken)
	accessToken, err := dingtalk.GetCorpToken(CORPID, suiteAccessToken)
	if err != nil {
		log.Fatalf("ding_sdk_golang.GetCorpToken error, err : %v", err)
	}
	log.Println("access_token: ", accessToken)

	// 免登授权码
	code := "f4196a7cd6e73908ab010ee5bc3b379a"
	userinfo, err := dingtalk.Getuserinfo(accessToken, code)
	if err != nil {
		log.Fatalf("ding_sdk_golang.GetCorpToken error, err : %v", err)
	}
	log.Println(userinfo)
	user, err := dingtalk.GetUser(accessToken, userinfo.Userid)
	if err != nil {
		log.Fatalf("ding_sdk_golang.GetCorpToken error, err : %v", err)
	}
	log.Println(user.Name)
}
