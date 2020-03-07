package main

import (
	"log"

	ding_sdk_golang "github.com/chiachan163/ding-sdk-golang"
)

func main() {
	const CORPID = "ding772ad001364d04a2bc961a6cb783455b"
	suiteAccessToken, err := ding_sdk_golang.GetSuiteToken("suitexiso0q2p5rxt3nr3", "kws8FIiNgFS2E0GD1CjQOfrD0uMJ9kO_bYs2JryR5GZU8Lm85xyMeX7YmHYjqAyZ", "")
	if err != nil {
		log.Fatalf("ding_sdk_golang.GetSuiteToken error, err : %v", err)
	}
	log.Println("suite_access_token: ", suiteAccessToken)
	accessToken, err := ding_sdk_golang.GetCorpToken(CORPID, suiteAccessToken)
	if err != nil {
		log.Fatalf("ding_sdk_golang.GetCorpToken error, err : %v", err)
	}
	log.Println("access_token: ", accessToken)

	// 免登授权码
	code := "f4196a7cd6e73908ab010ee5bc3b379a"
	userinfo, err := ding_sdk_golang.Getuserinfo(accessToken, code)
	if err != nil {
		log.Fatalf("ding_sdk_golang.GetCorpToken error, err : %v", err)
	}
	log.Println(userinfo)
	user, err := ding_sdk_golang.GetUser(accessToken, userinfo.Userid)
	if err != nil {
		log.Fatalf("ding_sdk_golang.GetCorpToken error, err : %v", err)
	}
	log.Println(user.Name)
}
