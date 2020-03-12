package ding_sdk_golang

import (
	"fmt"
	"log"
	"net/http"

	"github.com/swxctx/ghttp"
)

func GetCorpToken(authCorpid string, suiteAccessToken string) (access_token string, err error) {
	type Body struct {
		// 授权企业corpId
		AuthCorpid string `json:"auth_corpid"`
	}
	type Result struct {
		Errcode int32  `json:"errcode"`
		Errmsg  string `json:"errmsg"`
		// 授权方（企业）corp_access_token
		AccessToken string `json:"access_token"`
		// access_token超时时间
		ExpiresIn int32 `json:"expires_in"`
	}
	var result Result
	//timestamp := time.Now().Unix()
	//signature := sign.DingSign(timestamp, suiteTicket, suiteSecret)
	//
	body := &Body{
		AuthCorpid: authCorpid,
	}
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(GETCORPTOKENURL, suiteAccessToken),
		Body:        body,
		Method:      "POST",
		ContentType: "application/json",
	}.Do()
	if err != nil {
		log.Panicln(err)
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Request err, status -> %d", resp.StatusCode)
		return
	}
	err = resp.Body.FromToJson(&result)
	if err != nil {
		return
	}
	if result.Errcode != 0 {
		err = fmt.Errorf("%s", result.Errmsg)
		return
	}
	access_token = result.AccessToken
	return
}

func GetSuiteToken(suiteKey string, suiteSecret string, suiteTicket string) (suiteAccessToken string, err error) {
	type Body struct {
		// 钉钉推送的suiteTicket。测试应用可以随意填写。
		SuiteTicket string `json:"suite_ticket"`
		// 第三方应用suite_secret
		SuiteSecret string `json:"suite_secret"`
		// 第三方应用suite_key
		SuiteKey string `json:"suite_key"`
	}
	type Result struct {
		RespResult
		// 授权方（企业）corp_access_token
		SuiteAccessToken string `json:"suite_access_token"`
	}
	var result Result
	body := &Body{
		SuiteTicket: suiteTicket,
		SuiteSecret: suiteSecret,
		SuiteKey:    suiteKey,
	}
	resp, err := ghttp.Request{
		Url:         GETSUITETOKENURL,
		Body:        body,
		Method:      "POST",
		ContentType: "application/json",
	}.Do()
	if err != nil {
		log.Panicln(err)
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Request err, status -> %d", resp.StatusCode)
		return
	}
	err = resp.Body.FromToJson(&result)
	if err != nil {
		return
	}
	if result.Errcode != 0 {
		err = fmt.Errorf("%s", result.Errmsg)
		return
	}
	suiteAccessToken = result.SuiteAccessToken
	return
}
