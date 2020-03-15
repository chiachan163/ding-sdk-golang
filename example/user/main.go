package main

import (
	"log"

	"github.com/chiachan163/ding-sdk-golang/dingtalk"
)

func main() {
	const CORPID = "xxx"
	suiteAccessToken, err := dingtalk.GetSuiteToken("xxx", "xxx", "")
	if err != nil {
		log.Fatalf("ding_sdk_golang.GetSuiteToken error, err : %v", err)
	}
	log.Println("suite_access_token: ", suiteAccessToken)
	accessToken, err := dingtalk.GetCorpToken(CORPID, suiteAccessToken.SuiteAccessToken)
	if err != nil {
		log.Fatalf("ding_sdk_golang.GetCorpToken error, err : %v", err)
	}
	log.Println("access_token: ", accessToken)

	//// 获取管理员的微应用管理权限
	//canAccess, err := ding_sdk_golang.CanAccessMicroapp(accessToken, "36750", "manager2872")
	//if err != nil {
	//	log.Fatalf("ding_sdk_golang.CanAccessMicroapp error, err : %v", err)
	//}
	//log.Println(*canAccess)

	//// 获取管理员的微应用管理权限
	//userIds, err := ding_sdk_golang.GetDeptMember(accessToken, "1")
	//if err != nil {
	//	log.Fatalf("ding_sdk_golang.GetDeptMember error, err : %v", err)
	//}
	//log.Println(userIds)

	//// 根据unionid获取userid
	//userid, err := ding_sdk_golang.GetUseridByUnionid(accessToken, "7h8Oa0JQhKe6OrjmfExtJQiEiE")
	//if err != nil {
	//	log.Fatalf("ding_sdk_golang.GetUseridByUnionid error, err : %v", err)
	//}
	//log.Println(userid)

	//// 获取管理员列表
	//adminList, err := ding_sdk_golang.GetAdmin(accessToken)
	//if err != nil {
	//	log.Fatalf("ding_sdk_golang.GetAdmin error, err : %v", err)
	//}
	//for _, admin := range adminList {
	//	log.Println(*admin)
	//}

	//// 获取管理员通讯录权限范围
	//deptIds, err := ding_sdk_golang.GetAdminScope(accessToken, "manager2872")
	//if err != nil {
	//	log.Fatalf("ding_sdk_golang.GetAdminScope error, err : %v", err)
	//}
	//log.Println(deptIds)

	//// 获取企业员工人数
	//count, err := ding_sdk_golang.GetOrgUserCount(accessToken, 0)
	//if err != nil {
	//	log.Fatalf("ding_sdk_golang.GetOrgUserCount error, err : %v", err)
	//}
	//log.Println(count)

	// 获取部门用户详情
	userList, err := dingtalk.ListByPage(accessToken.AccessToken, 1, nil, nil, nil, nil)
	if err != nil {
		log.Fatalf("ding_sdk_golang.ListByPage error, err : %v", err)
	}
	log.Println(userList.Userlist[0])

	////获取部门用户详情
	//userList, err := ding_sdk_golang.Simplelist(accessToken, 1, nil, nil, nil, nil)
	//if err != nil {
	//	log.Fatalf("ding_sdk_golang.ListByPage error, err : %v", err)
	//}
	//log.Println(userList.Userlist[0])
}
