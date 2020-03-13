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

	// 获取jsapi ticket
	ticket, err := dingtalk.GetJsapiTicket(accessToken)
	if err != nil {
		log.Fatalf("ding_sdk_golang.GetCorpToken error, err : %v", err)
	}
	log.Println(*ticket)
	//log.Println("ticket: ", fmt.Sprintf("%p", unsafe.Pointer(ticket)))
}
