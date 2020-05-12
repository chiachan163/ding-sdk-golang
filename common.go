package ding_sdk_golang

import (
	"fmt"
	"time"

	"github.com/chiachan163/ding-sdk-golang/v1/redis"
)

var ReTryTimes = 5

type RespResult struct {
	Errcode int32  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

//func CallDingTalk(corpId string, apiPath int, fn func(corpId string, apiPath int) (*RespResult, error)) error {
//	// 判断20s内本IP总调用钉钉API次数是否超过3000
//	allTimes, err := redis.GetCallDingTalkAll()
//	if err != nil {
//		return fmt.Errorf("redis.GetCallDingTalkAll error: %v", err)
//	}
//	if allTimes > 3000 {
//		return fmt.Errorf("调用钉钉平台接口总量超出限制")
//	}
//
//	res, rerr := fn(corpId, apiPath)
//	redis.InrcCallDingTalkAll()
//	redis.InrcCallDingTalkCorpIdApiPath(corpId, apiPath)
//	if res.Errcode == 0 {
//		return nil
//	}
//	for i := 0; i <= arg.TryTime; i++ {
//		switch res.Errcode {
//		case 90019, 90010, 90008, 90014:
//			time.Sleep(time.Second)
//			res, rerr = fn(corpId, apiPath)
//		default:
//			rerr = nil
//		}
//	}
//	rerr = fmt.Errorf(res.Errmsg)
//	return rerr
//}

func CallDingTalk(corpId string, suiteKey *string, apiPath int, fn func(corpId string, apiPath int) (*RespResult, error)) error {
	// 判断20s内本IP总调用钉钉API次数是否超过3000
	var tryTime = 0
	allTimes, err := redis.GetCallDingTalkAll()
	if err != nil {
		return fmt.Errorf("redis.GetCallDingTalkAll error: %v", err)
	}
	if allTimes >= 3000 {
		return fmt.Errorf("20s内调用钉钉平台接口总量超出限制")
	}

	apiPathTimes, err := redis.GetCallDingTalkApiPath(apiPath)
	if err != nil {
		return fmt.Errorf("redis.GetCallDingTalkAll error: %v", err)
	}
	if apiPathTimes >= 2000 {
		return fmt.Errorf("1min内调用单个接口的频率超过2000次")
	}

	corpIdApiPathTimes, err := redis.GetCallDingTalkCorpIdApiPath(corpId, apiPath)
	if err != nil {
		return fmt.Errorf("redis.GetCallDingTalkAll error: %v", err)
	}
	if corpIdApiPathTimes >= 1500 {
		return fmt.Errorf("1min内调用单个企业的单个接口的频率超过1500次")
	}
	if suiteKey != nil && *suiteKey != "" {
		corpIdSuiteKeyApiPathTimes, err := redis.GetCallDingTalkCorpIdSuiteKeyApiPath(corpId, *suiteKey, apiPath)
		if err != nil {
			return fmt.Errorf("redis.GetCallDingTalkAll error: %v", err)
		}
		if corpIdSuiteKeyApiPathTimes >= 1000 {
			return fmt.Errorf("套件%s在1min内调用单个企业的单个接口的频率超过1000次", *suiteKey)
		}
	}

	res, rerr := fn(corpId, apiPath)
	redis.InrcCallDingTalkAll()
	redis.InrcCallDingTalkApiPath(apiPath)
	redis.InrcCallDingTalkCorpIdApiPath(corpId, apiPath)
	redis.InrcCallDingTalkCorpIdSuiteKeyApiPath(corpId, *suiteKey, apiPath)
	if res.Errcode == 0 {
		return nil
	}
	for i := 0; i <= tryTime; i++ {
		switch res.Errcode {
		case 90019, 90010, 90008, 90014:
			time.Sleep(time.Second)
			res, rerr = fn(corpId, apiPath)
		default:
			rerr = nil
		}
	}
	rerr = fmt.Errorf(res.Errmsg)
	return rerr
}
