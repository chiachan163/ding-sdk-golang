package main

import (
	"log"

	ding_sdk_golang "github.com/chiachan163/ding-sdk-golang"
)

func main() {
	const CORPID = "xxx"
	suiteAccessToken, err := ding_sdk_golang.GetSuiteToken("xxx", "xxx", "")
	if err != nil {
		log.Fatalf("ding_sdk_golang.GetSuiteToken error, err : %v", err)
	}
	log.Println("suite_access_token: ", suiteAccessToken)
	accessToken, err := ding_sdk_golang.GetCorpToken(CORPID, suiteAccessToken)
	if err != nil {
		log.Fatalf("ding_sdk_golang.GetCorpToken error, err : %v", err)
	}
	log.Println("access_token: ", accessToken)

	// 获取jsapi ticket
	ticket, err := ding_sdk_golang.GetJsapiTicket(accessToken)
	if err != nil {
		log.Fatalf("ding_sdk_golang.GetCorpToken error, err : %v", err)
	}
	log.Println(*ticket)
	//log.Println("ticket: ", fmt.Sprintf("%p", unsafe.Pointer(ticket)))
}
