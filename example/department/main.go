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
	accessToken, err := dingtalk.GetCorpToken(CORPID, suiteAccessToken)
	if err != nil {
		log.Fatalf("ding_sdk_golang.GetCorpToken error, err : %v", err)
	}
	log.Println("access_token: ", accessToken)

	//// 获取部门详情
	//department, err := ding_sdk_golang.GetDepartment(accessToken, 1, nil)
	//if err != nil {
	//	log.Fatalf("ding_sdk_golang.GetDepartment error, err : %v", err)
	//}
	//log.Println(*department)

	//// 获取部门列表
	//var id int64 = 1
	//department, err := ding_sdk_golang.DepartmentList(accessToken, &id, nil, nil)
	//if err != nil {
	//	log.Fatalf("ding_sdk_golang.GetDepartment error, err : %v", err)
	//}
	//log.Println(department[0])

	//// 获取子部门ID列表
	//subDeptIdList, err := ding_sdk_golang.DepartmentListIds(accessToken, 1)
	//if err != nil {
	//	log.Fatalf("ding_sdk_golang.GetDepartment error, err : %v", err)
	//}
	//log.Println(subDeptIdList)

	//// 查询指定用户的所有上级父部门路径
	//department, err := ding_sdk_golang.DepartmentListParentDepts(accessToken, "manager2872")
	//if err != nil {
	//	log.Fatalf("ding_sdk_golang.GetDepartment error, err : %v", err)
	//}
	//log.Println(department)

	// 查询部门的所有上级父部门路径
	parentIds, err := dingtalk.DepartmentListParentDeptsByDept(accessToken, 280467091)
	if err != nil {
		log.Fatalf("ding_sdk_golang.GetDepartment error, err : %v", err)
	}
	log.Println(parentIds)

	//log.Println("ticket: ", fmt.Sprintf("%p", unsafe.Pointer(ticket)))
}
