package redis

const (
	// 总调用次数 每20秒限制3000次
	CallDingTalkAll = "dingtalk:all:%d"
	// 企业调用接口次数 每分钟限制2000次
	CallDingTalkCorpIdApiPath = "dingtalk:%s:%d"
	// 企业调用单个接口的次数 每分钟限制1500次
	CallDingTalkApiPath = "dingtalk:%d"
	// 企业的套件调用单个接口的频率 每分钟限制1000次
	CallDingTalkCorpIdSuiteKeyApiPath = "dingtalk:%s:%s:%d"
)
