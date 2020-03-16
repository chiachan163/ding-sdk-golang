package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"

	"github.com/chiachan163/ding-sdk-golang/utils"
)

func GetCallDingTalkAllKey() (string, error) {
	ip, err := utils.GetIntranetIp()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(CallDingTalkAll, ip[len(ip)-1]), nil
}

// 调用总数自增
func InrcCallDingTalkAll() error {
	key, err := GetCallDingTalkAllKey()
	if err != nil {
		return fmt.Errorf("GetCallDingTalkAllKey error: %v", err)
	}
	if err := GetRedis().Incr(key).Err(); err != nil {
		return err
	}
	// 设置过期时间
	if err := GetRedis().Expire(key, time.Duration(20)*time.Second).Err(); err != nil {
		return err
	}
	return nil
}

func GetCallDingTalkAll() (int64, error) {
	key, err := GetCallDingTalkAllKey()
	if err != nil {
		return 0, fmt.Errorf("GetCallDingTalkAllKey error: %v", err)
	}
	_v, err := GetRedis().Get(key).Int64()
	if err != nil && err != redis.Nil {
		return 0, fmt.Errorf("GetCallDingTalkAll error: %v", err)
	}
	return _v, nil
}

// 调用总数自增
func InrcCallDingTalkCorpIdApiPath(corpId string, apiPath int) error {
	key := fmt.Sprintf(CallDingTalkCorpIdApiPath, corpId, apiPath)
	if err := GetRedis().Incr(key).Err(); err != nil {
		return err
	}
	// 设置过期时间
	if err := GetRedis().Expire(key, time.Duration(20)*time.Second).Err(); err != nil {
		return err
	}
	return nil
}

func GetCallDingTalkCorpIdApiPath(corpId string, apiPath int) (int64, error) {
	key := fmt.Sprintf(CallDingTalkCorpIdApiPath, corpId, apiPath)
	_v, err := GetRedis().Get(key).Int64()
	if err != nil && err != redis.Nil {
		return 0, fmt.Errorf("GetCallDingTalkAll error: %v", err)
	}
	return _v, nil
}
