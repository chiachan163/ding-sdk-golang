package dingtalk

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chiachan163/ding-sdk-golang/v2/arg"

	"github.com/swxctx/ghttp"
)

func GetCorpToken(authCorpid string, suiteAccessToken string) (_result *arg.GetCorpTokenResult, err error) {
	type Body struct {
		// 授权企业corpId
		AuthCorpid string `json:"auth_corpid"`
	}

	var result arg.GetCorpTokenResult
	//timestamp := time.Now().Unix()
	//signature := sign.DingSign(timestamp, suiteTicket, suiteSecret)
	//
	body := &Body{
		AuthCorpid: authCorpid,
	}
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(arg.GETCORPTOKENURL+"?suite_access_token=%s", suiteAccessToken),
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
	_result = &result
	return
}

func GetSuiteToken(suiteKey string, suiteSecret string, suiteTicket string) (_result *arg.GetSuiteTokenResult, err error) {
	type Body struct {
		// 钉钉推送的suiteTicket。测试应用可以随意填写。
		SuiteTicket string `json:"suite_ticket"`
		// 第三方应用suite_secret
		SuiteSecret string `json:"suite_secret"`
		// 第三方应用suite_key
		SuiteKey string `json:"suite_key"`
	}

	var result arg.GetSuiteTokenResult
	body := &Body{
		SuiteTicket: suiteTicket,
		SuiteSecret: suiteSecret,
		SuiteKey:    suiteKey,
	}
	resp, err := ghttp.Request{
		Url:         arg.GETSUITETOKENURL,
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

	_result = &result
	return
}

func GetAuthInfo(suiteKey string, authCorpid string, suiteAccessToken string) (_result *arg.GetAuthInfoResult, err error) {
	type Body struct {
		// 授权方ID
		AuthCorpid string `json:"auth_corpid"`
		// 第三方应用suite_key
		SuiteKey string `json:"suite_key"`
	}

	var result arg.GetAuthInfoResult
	body := &Body{
		AuthCorpid: authCorpid,
		SuiteKey:   suiteKey,
	}
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf("%s?suite_access_token=%s", arg.SERVICEGETAUTHINFO, suiteAccessToken),
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

	_result = &result
	return
}
