# ding-sdk-golang
基于go语言适配的钉钉SDK(第三方企业应用)
## 使用方法
```
import (
	"log"
	"time"

	ding_sdk_golang "github.com/chiachan163/ding-sdk-golang"
	ding_redis "github.com/chiachan163/ding-sdk-golang/redis"
	"github.com/xiaoenai/tp-micro/v6/model/redis"
)

func init() {
	config := &redis.Config{
		DeployType: "single",
		ForSingle: redis.SingleConfig{
			Addr: "127.0.0.1:6379",
		},
		ForCluster:      redis.ClusterConfig{},
		MaxRetries:      0,
		DialTimeout:     0,
		PoolSizePerNode: 0,
		IdleTimeout:     0,
	}
	ding_redis.Init(config, time.Hour*24)
}

func main() {
	const CORPID = "xxx"

	ding_sdk_golang.CallDingTalk(CORPID, ding_sdk_golang.GETSUITETOKENURL, func(corpId string, apiPath string) error {
		suiteAccessToken, err := ding_sdk_golang.GetSuiteToken("xxx", "xxx", "")
		if err != nil {
			log.Fatalf("ding_sdk_golang.GetSuiteToken error, err : %v", err)
		}
		log.Println("suite_access_token: ", suiteAccessToken)
		return nil
	})
}
```
## 已适配内容
### 获取凭证
- 获取jsapi ticket 
    - /get_jsapi_ticket
- 获取企业授权凭证
    - /service/get_corp_token
- 获取第三方凭证 
    - /service/get_suite_token
### 应用授权
### 获取登录身份
- 获取用户userid 
    - /user/getuserinfo

### 通讯录管理-用户管理
- 获取管理员的微应用管理权限 
    - /user/can_access_microapp
- 根据部门id获取员工ID列表 
    - /user/getDeptMember
- 根据unionid获取userid
    - /user/getUseridByUnionid
- 获取管理员列表
    - /user/get_admin
- 获取管理员通讯录权限范围
    - /topapi/user/get_admin_scope
- 获取企业员工人数 
    - /user/get_org_user_count
- 获取部门用户详情 
    - /user/listbypage
- 获取部门用户 
    - /user/simplelist

### 通讯录管理-部门管理
- 获取部门详情
    - /department/get
- 获取部门列表
    - /department/list
- 获取子部门ID列表 
    - /department/list_ids
- 查询指定用户的所有上级父部门路径
    - /department/list_parent_depts
- 查询部门的所有上级父部门路径
    - /department/list_parent_depts_by_dept
    
### 通讯录管理-角色管理
### 通讯录管理-外部联系人
### 通讯录管理-家校通讯录
### 通讯录管理-权限范围
### 消息通知-工作通知消息
### 消息通知-普通消息
### 智能工作流-官方工作流
### 智能工作量-自有工作流
### 待办任务
### 考勤-排班
### 考勤-班次
### 考勤-假勤审批
### 考勤-考勤组管理
### 考勤-统计
### 文件存储-管理媒体文件
### 文件存储-管理钉盘文件
### 智能办公电话
### 业务事务回调
### 应用内购