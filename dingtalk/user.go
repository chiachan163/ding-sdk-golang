package dingtalk

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chiachan163/ding-sdk-golang/v2/arg"

	"github.com/swxctx/ghttp"
)

// 获取用户userid
// 通过免登授权码和access_token获取用户的userid。
// 请求参数：
// access_token 调用接口凭证
// code 免登授权码

func Getuserinfo(accessToken string, code string) (_result *arg.GetuserinfoResult, err error) {
	//type Body struct {
	//	// 调用接口凭证
	//	AccessToken string `json:"access_token"`
	//	// 免登授权码
	//	Code string `json:"code"`
	//}

	var result arg.GetuserinfoResult
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(arg.USERGETUSERINFO+"?access_token=%s&code=%s", accessToken, code),
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

// 获取用户详情
// 通过access_token和userid获取用户的信息
func GetUser(accessToken, userid string) (_result *arg.GetUserResult, err error) {

	var result arg.GetUserResult
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(arg.USERGET+"?access_token=%s&userid=%s", accessToken, userid),
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

// 获取管理员的微应用管理权限
func CanAccessMicroapp(accessToken string, appId string, userId string) (_result *arg.CanAccessMicroappResult, err error) {

	var result arg.CanAccessMicroappResult
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(arg.USERCANACCESSMICROAPP+"?access_token=%s&appId=%s&userId=%s", accessToken, appId, userId),
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

// 根据部门id获取员工ID列表
func GetDeptMember(accessToken string, deptId string) (_result *arg.GetDeptMemberResult, err error) {

	var result arg.GetDeptMemberResult
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(arg.USERGETDEPTMEMBER+"?access_token=%s&deptId=%s", accessToken, deptId),
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

// 根据unionid获取userid
func GetUseridByUnionid(accessToken string, unionid string) (_result *arg.GetUseridByUnionidResult, err error) {

	var result arg.GetUseridByUnionidResult
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(arg.USERGETUSERIDBYUNIONID+"?access_token=%s&unionid=%s", accessToken, unionid),
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

// 根据unionid获取userid
func GetAdmin(accessToken string) (_result *arg.GetAdminResult, err error) {

	var result arg.GetAdminResult
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(arg.USERGETADMIN+"?access_token=%s", accessToken),
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

// 获取管理员通讯录权限范围
func GetAdminScope(accessToken string, userid string) (_result *arg.GetAdminScopeResult, err error) {

	var result arg.GetAdminScopeResult
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(arg.USERGETTADMINSCOPR+"?access_token=%s&userid=%s", accessToken, userid),
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

// 获取企业员工人数
func GetOrgUserCount(accessToken string, onlyActive int) (_result *arg.GetOrgUserCountResult, err error) {

	var result arg.GetOrgUserCountResult
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(arg.USERGETORGUSERCOUNT+"?access_token=%s&onlyActive=%d", accessToken, onlyActive),
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

// 获取部门用户详情
// offset: 与size参数同时设置时才生效，此参数代表偏移量,偏移量从0开始
// size: 与offset参数同时设置时才生效，此参数代表分页大小，最大100
// order: 默认 是按自定义排序；[
//     entry_asc：代表按照进入部门的时间升序，
//     entry_desc：代表按照进入部门的时间降序，
//     modify_asc：代表按照部门信息修改时间升序，
//     modify_desc：代表按照部门信息修改时间降序，
//     custom：代表用户定义(未定义时按照拼音)排序
// ]

func ListByPage(accessToken string, departmentId int64, offset *int, size *int, order *string, lang *string) (_result *arg.ListByPageResult, err error) {

	_offset := 0
	_size := 5
	_order := "custom"
	_lang := "zh_CN"
	if offset != nil {
		_offset = *offset
	}
	if size != nil {
		_size = *size
	}
	if order != nil {
		_order = *order
	}
	if lang != nil {
		_lang = *lang
	}
	var result arg.ListByPageResult
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(arg.USERLISTBYPAGE+"?access_token=%s&department_id=%d&offset=%d&size=%d&orde=%s&lane=%s", accessToken, departmentId, _offset, _size, _order, _lang),
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

func Simplelist(accessToken string, departmentId int64, offset *int, size *int, order *string, lang *string) (_result *arg.SimplelistResult, err error) {

	url := fmt.Sprintf(arg.USERSIMPLELIST+"?access_token=%s&department_id=%d", accessToken, departmentId)
	if offset != nil {
		url = fmt.Sprintf("%s&offset=%d", url, *offset)
	}
	if size != nil {
		url = fmt.Sprintf("%s&size=%d", url, *size)
	}
	if order != nil {
		url = fmt.Sprintf("%s&order=%s", url, *order)
	}
	if lang != nil {
		url = fmt.Sprintf("%s&lang=%s", url, *lang)
	}
	var result arg.SimplelistResult
	resp, err := ghttp.Request{
		Url:         url,
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
