package arg

import ding_sdk_golang "github.com/chiachan163/ding-sdk-golang"

type GetCorpTokenResult struct {
	ding_sdk_golang.RespResult
	// 授权方（企业）corp_access_token
	AccessToken string `json:"access_token"`
	// access_token超时时间
	ExpiresIn int32 `json:"expires_in"`
}

type GetSuiteTokenResult struct {
	ding_sdk_golang.RespResult
	// 授权方（企业）corp_access_token
	SuiteAccessToken string `json:"suite_access_token"`
}

type GetDepartmentResult struct {
	ding_sdk_golang.RespResult
	Department
}

type DepartmentListResult struct {
	ding_sdk_golang.RespResult
	Department []*SimpleDepartment `json:"department"`
}

type DepartmentListIdsResult struct {
	ding_sdk_golang.RespResult
	SubDeptIdList []int64 `json:"sub_dept_id_list"`
}

type DepartmentListParentDeptsResult struct {
	ding_sdk_golang.RespResult
	Department [][]int64 `json:"department"`
}

type DepartmentListParentDeptsByDeptResult struct {
	ding_sdk_golang.RespResult
	ParentIds []int64 `json:"parentIds"`
}

type GetJsapiTicketResult struct {
	ding_sdk_golang.RespResult
	Ticket string `json:"ticket"`
}

type GetuserinfoResult struct {
	ding_sdk_golang.RespResult
	UserInfo
}

type GetUserResult struct {
	ding_sdk_golang.RespResult
	User
}

type CanAccessMicroappResult struct {
	ding_sdk_golang.RespResult
	CanAccess bool `json:"canAccess"`
}

type GetDeptMemberResult struct {
	ding_sdk_golang.RespResult
	UserIds []string `json:"userIds"`
}

type GetUseridByUnionidResult struct {
	ding_sdk_golang.RespResult
	Userid
}

type GetAdminResult struct {
	ding_sdk_golang.RespResult
	AdminList []*AdminList `json:"admin_list"`
}

type GetAdminScopeResult struct {
	ding_sdk_golang.RespResult
	DeptIds []int64 `json:"dept_ids"`
}

type GetOrgUserCountResult struct {
	ding_sdk_golang.RespResult
	// 企业员工数量
	Count int64 `json:"count"`
}

type ListByPageResult struct {
	ding_sdk_golang.RespResult
	UserList
}

type SimplelistResult struct {
	ding_sdk_golang.RespResult
	SimpleUserList
}

type GetAuthInfoResult struct {
	ding_sdk_golang.RespResult
	AuthUserInfo    *AuthUserInfo    `json:"auth_user_info"`
	AuthCorpInfo    *AuthCorpInfo    `json:"auth_corp_info"`
	ChannelAuthInfo *ChannelAuthInfo `json:"channel_auth_info"`
	AuthInfo        *AuthInfo        `json:"auth_info"`
	AuthMarketInfo  *AuthMarketInfo  `json:"auth_market_info"`
}
