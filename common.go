package ding_sdk_golang

import (
	"fmt"

	"github.com/chiachan163/ding-sdk-golang/redis"
)

type RespResult struct {
	Errcode int32  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func CallDingTalk(corpId string, apiPath string, fn func(corpId string, apiPath string) error) error {
	// 判断20s内本IP总调用钉钉API次数是否超过3000
	allTimes, err := redis.GetCallDingTalkAll()
	if err != nil {
		return fmt.Errorf("redis.GetCallDingTalkAll error: %v", err)
	}
	if allTimes > 3000 {
		return fmt.Errorf("调用钉钉平台接口总量超出限制")
	}
	return fn(corpId, apiPath)
}
