package arg

/*
 * 路径管理
 */
const DINGTALK = "https://oapi.dingtalk.com"

// 获取凭证
const (
	// 获取jsapi ticket
	GETJSAPITICKET = DINGTALK + "/get_jsapi_ticket"
	// 获取企业凭证
	GETCORPTOKENURL = DINGTALK + "/service/get_corp_token"
	// 获取第三方凭证
	GETSUITETOKENURL = DINGTALK + "/service/get_suite_token"
)

// 应用授权
// 获取登录身份
const (
	// 获取用户userid
	USERGETUSERINFO = "https://oapi.dingtalk.com/user/getuserinfo"
)

// 通讯录管理-用户管理
const (
	// 获取用户详情
	USERGET = DINGTALK + "/user/get"
	// 获取管理员的微应用管理权限
	USERCANACCESSMICROAPP = DINGTALK + "/user/can_access_microapp"
	// 根据部门id获取员工ID列表
	USERGETDEPTMEMBER = DINGTALK + "/user/getDeptMember"
	// 根据unionid获取userid
	USERGETUSERIDBYUNIONID = DINGTALK + "/user/getUseridByUnionid"
	// 获取管理员列表
	USERGETADMIN = DINGTALK + "/user/get_admin"
	// 获取管理员通讯录权限范围
	USERGETTADMINSCOPR = DINGTALK + "/topapi/user/get_admin_scope"
	// 获取企业员工人数
	USERGETORGUSERCOUNT = DINGTALK + "/user/get_org_user_count"
	// 获取部门用户详情
	USERLISTBYPAGE = DINGTALK + "/user/listbypage"
	// 获取部门用户
	USERSIMPLELIST = DINGTALK + "/user/simplelist"
)

// 通讯录管理-部门管理
const (
	// 获取部门详情
	DEPARTMENTGET = DINGTALK + "/department/get"
	// 获取部门列表
	DEPARTMENTLIST = DINGTALK + "/department/list"
	// 获取子部门ID列表
	DEPARTMENTLISTIDS = DINGTALK + "/department/list_ids"
	// 查询指定用户的所有上级父部门路径
	DEPARTMENTLISTPARENTDEPTS = DINGTALK + "/department/list_parent_depts"
	// 查询部门的所有上级父部门路径
	DEPARTMENTLISTPARENTDEPTSBYDEPT = DINGTALK + "/department/list_parent_depts_by_dept"
)

// 通讯录管理-角色管理
// 通讯录管理-外部联系人
// 通讯录管理-家校通讯录
// 通讯录管理-权限范围
// 消息通知-工作通知消息
// 消息通知-普通消息
// 智能工作流-官方工作流
// 智能工作量-自有工作流
// 待办任务
// 考勤-排班
// 考勤-班次
// 考勤-假勤审批
// 考勤-考勤组管理
// 考勤-统计
// 文件存储-管理媒体文件
// 文件存储-管理钉盘文件
// 智能办公电话
// 业务事务回调
// 应用内购
