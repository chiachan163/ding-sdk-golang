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
	userInfo = &result.UserInfo
	return
}

type (
	User struct {
		Unionid         string  `json:"unionid"`
		Userid          string  `json:"userid"`
		Name            string  `json:"name"`
		Tel             string  `json:"tel"`
		WorkPlace       string  `json:"workPlace"`
		Remark          string  `json:"remark"`
		Mobile          string  `json:"mobile"`
		Email           string  `json:"email"`
		OrgEmail        string  `json:"orgEmail"`
		Active          bool    `json:"active"`
		OrderInDepts    string  `json:"orderInDepts"`
		IsAdmin         bool    `json:"isAdmin"`
		IsBoss          bool    `json:"isBoss"`
		IsLeaderInDepts string  `json:"isLeaderInDepts"`
		IsHide          bool    `json:"isHide"`
		Department      []int64 `json:"department"`
		Position        string  `json:"position"`
		Avatar          string  `json:"avatar"`
		HiredDate       int64   `json:"hiredDate"`
		Jobnumber       string  `json:"jobnumber"`
		IsSenior        bool    `json:"isSenior"`
		StateCode       string  `json:"stateCode"`
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
	user = &result.User
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
		Url:         fmt.Sprintf(USERGETUSERIDBYUNIONID, accessToken, unionid),
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

type AdminList struct {
	// 管理员角色，1:主管理员，2:子管理员
	SysLevel int32 `json:"sys_level"`
	// 员工id
	Userid string `json:"userid"`
}

// 根据unionid获取userid
func GetAdmin(accessToken string) (adminList []*AdminList, err error) {
	type Result struct {
		RespResult
		AdminList []*AdminList `json:"admin_list"`
	}
	var result Result
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(USERGETADMIN, accessToken),
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
	adminList = result.AdminList
	return
}

// 获取管理员通讯录权限范围
func GetAdminScope(accessToken string, userid string) (deptIds []int64, err error) {
	type Result struct {
		RespResult
		DeptIds []int64 `json:"dept_ids"`
	}
	var result Result
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(USERGETTADMINSCOPR, accessToken, userid),
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
	deptIds = result.DeptIds
	return
}

// 获取企业员工人数
func GetOrgUserCount(accessToken string, onlyActive int) (count int64, err error) {
	type Result struct {
		RespResult
		// 企业员工数量
		Count int64 `json:"count"`
	}
	var result Result
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(USERGETORGUSERCOUNT, accessToken, onlyActive),
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
	count = result.Count
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
type UserList struct {
	// 成员列表
	Userlist []*User `json:"userlist"`
	// 在分页查询时返回，true代表还有下一页更多数据
	HasMore bool `json:"hasMore"`
}

func ListByPage(accessToken string, departmentId int64, offset *int, size *int, order *string, lang *string) (userList *UserList, err error) {

	type Result struct {
		RespResult
		UserList
	}
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
	var result Result
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(USERLISTBYPAGE, accessToken, departmentId, _offset, _size, _order, _lang),
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
	userList = &result.UserList
	return
}

type SimpleUser struct {
	// 成员名称
	Name string `json:"name"`
	// 员工id
	Userid string `json:"userid"`
}

type SimpleUserList struct {
	// 成员列表
	Userlist []*SimpleUser `json:"userlist"`
	// 在分页查询时返回，true代表还有下一页更多数据
	HasMore bool `json:"hasMore"`
}

func Simplelist(accessToken string, departmentId int64, offset *int, size *int, order *string, lang *string) (userList *SimpleUserList, err error) {

	type Result struct {
		RespResult
		SimpleUserList
	}
	url := fmt.Sprintf(USERSIMPLELIST, accessToken, departmentId)
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
	var result Result
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
	if result.Errcode != 0 {
		err = fmt.Errorf("%s", result.Errmsg)
		return
	}
	log.Println(result)
	userList = &result.SimpleUserList
	return
}
