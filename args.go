package ding_sdk_golang

/*
 * 路径管理
 */
const DINGTALK = "https://oapi.dingtalk.com"

// 获取企业凭证
const GETCORPTOKENURL = DINGTALK + "/service/get_corp_token?suite_access_token=%s"

// 获取第三方凭证
const GETSUITETOKENURL = DINGTALK + "/service/get_suite_token"

// 获取用户userid
const USERGETUSERINFO = "https://oapi.dingtalk.com/user/getuserinfo?access_token=%s&code=%s"

// 获取用户详情
const USERGET = DINGTALK + "/user/get?access_token=%s&userid=%s"
