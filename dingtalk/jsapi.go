package dingtalk

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chiachan163/ding-sdk-golang/arg"

	"github.com/swxctx/ghttp"
)

func GetJsapiTicket(accessToken string) (_result *arg.GetJsapiTicketResult, err error) {
	var result arg.GetJsapiTicketResult
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(arg.GETJSAPITICKET+"?access_token=%s", accessToken),
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

	_result = &result
	return
}
