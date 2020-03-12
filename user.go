package ding_sdk_golang

import (
	"fmt"
	"log"
	"net/http"

	"github.com/swxctx/ghttp"
)

// 获取用户userid
// 通过免登授权码和access_token获取用户的userid。
// 请求参数：
// access_token 调用接口凭证
// code 免登授权码

type UserInfo struct {
	// 员工姓名
	Name string `json:"name"`
	// 员工在当前企业内的唯一标识，也称staffId
	Userid string `json:"userid"`
	// 是否是管理员，true：是，false：不是
	IsSys bool `json:"is_sys"`
	// 级别，1：主管理员，2：子管理员，100：老板，0：其他（如普通员工）
	SysLevel int32 `json:"sys_level"`
}

func Getuserinfo(accessToken string, code string) (userInfo *UserInfo, err error) {
	//type Body struct {
	//	// 调用接口凭证
	//	AccessToken string `json:"access_token"`
	//	// 免登授权码
	//	Code string `json:"code"`
	//}
	type Result struct {
		RespResult
		UserInfo
	}
	var result Result
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(USERGETUSERINFO, accessToken, code),
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
	userInfo = &UserInfo{
		Name:     result.Name,
		Userid:   result.Userid,
		IsSys:    result.IsSys,
		SysLevel: result.SysLevel,
	}
	return
}

type (
	User struct {
		Name            string  `json:"name"`
		Unionid         string  `json:"unionid"`
		Userid          string  `json:"userid"`
		IsLeaderInDepts string  `json:"isLeaderInDepts"`
		IsBoss          bool    `json:"isBoss"`
		HiredDate       int64   `json:"hiredDate"`
		IsSenior        bool    `json:"isSenior"`
		Department      []int32 `json:"department"`
		OrderInDepts    string  `json:"orderInDepts"`
		Active          bool    `json:"active"`
		Avatar          string  `json:"avatar"`
		IsAdmin         bool    `json:"isAdmin"`
		IsHide          bool    `json:"isHide"`
		Jobnumber       string  `json:"jobnumber"`
		Position        string  `json:"position"`
		Roles           []*Role `json:"roles"`
	}
	IsLeaderInDepts struct {
		// 部门ID
		Key int32
		// 是否为主管
		Value bool
	}
	Role struct {
		// 角色id
		Id int64
		// 角色名称
		Name string
		// 角色组名称
		GroupName string
	}
)

// 获取用户详情
// 通过access_token和userid获取用户的信息
func GetUser(accessToken, userid string) (user *User, err error) {
	type Result struct {
		RespResult
		User
	}
	var result Result
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(USERGET, accessToken, userid),
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
	user = &User{
		Name:            result.Name,
		Unionid:         result.Unionid,
		Userid:          result.Userid,
		IsLeaderInDepts: result.IsLeaderInDepts,
		IsBoss:          result.IsBoss,
		HiredDate:       result.HiredDate,
		IsSenior:        result.IsSenior,
		Department:      result.Department,
		OrderInDepts:    result.OrderInDepts,
		Active:          result.Active,
		Avatar:          result.Avatar,
		IsAdmin:         result.IsAdmin,
		IsHide:          result.IsHide,
		Jobnumber:       result.Jobnumber,
		Position:        result.Position,
		Roles:           result.Roles,
	}
	return
}

// 获取管理员的微应用管理权限
func CanAccessMicroapp(accessToken string, appId string, userId string) (canAccess *bool, err error) {
	type Result struct {
		RespResult
		CanAccess bool `json:"canAccess"`
	}
	var result Result
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(USERCANACCESSMICROAPP, accessToken, appId, userId),
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
	canAccess = &result.CanAccess
	return
}

// 根据部门id获取员工ID列表
func GetDeptMember(accessToken string, deptId string) (userIds []string, err error) {
	type Result struct {
		RespResult
		UserIds []string `json:"userIds"`
	}
	var result Result
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(USERGETDEPTMEMBER, accessToken, deptId),
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
	userIds = result.UserIds
	return
}

type Userid struct {
	// 联系类型，0表示企业内部员工，1表示企业外部联系人
	ContactType int32 `json:"contactType"`
	// 员工id
	Userid string `json:"userid"`
}

// 根据unionid获取userid
func GetUseridByUnionid(accessToken string, unionid string) (userid *Userid, err error) {
	type Result struct {
		RespResult
		Userid
	}
	var result Result
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(GETUSERIDBYUNIONID, accessToken, unionid),
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
	userid = &result.Userid
	return
}
