package ding_sdk_golang

import (
	"fmt"
	"log"
	"net/http"

	"github.com/swxctx/ghttp"
)

func GetJsapiTicket(accessToken string) (ticket *string, err error) {
	type Result struct {
		RespResult
		Ticket string `json:"ticket"`
	}
	var result Result
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(GETJSAPITICKET, accessToken),
		Body:        nil,
		Method:      "GET",
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
	log.Println(result)
	ticket = &result.Ticket
	return
}
