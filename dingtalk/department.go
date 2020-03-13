package dingtalk

import (
	"fmt"
	"log"
	"net/http"

	ding_sdk_golang "github.com/chiachan163/ding-sdk-golang"

	"github.com/chiachan163/ding-sdk-golang/arg"

	"github.com/swxctx/ghttp"
)

type Department struct {
	// 部门id
	Id int64 `json:"id"`
	//部门名称
	Name string `json:"name"`
	// 父部门id，根部门为1
	Parentid int64 `json:"parentid"`
	// 当前部门在父部门下的所有子部门中的排序值
	Order int64 `json:"order"`
	// 是否同步创建一个关联此部门的企业群，true表示是，false表示不是
	CreateDeptGroup bool `json:"createDeptGroup"`
	// 当部门群已经创建后，是否有新人加入部门会自动加入该群，true表示是，false表示不是
	AutoAddUser bool `json:"autoAddUser"`
	// 是否隐藏部门，true表示隐藏，false表示显示
	DeptHiding bool `json:"deptHiding"`
	// 可以查看指定隐藏部门的其他部门列表，如果部门隐藏，则此值生效，取值为其他的部门id组成的的字符串，使用“\|”符号进行分割
	DeptPermits string `json:"deptPermits"`
	// 可以查看指定隐藏部门的其他人员列表，如果部门隐藏，则此值生效，取值为其他的人员userid组成的的字符串，使用“\|”符号进行分割
	UserPermits string `json:"userPermits"`
	// 是否本部门的员工仅可见员工自己，为true时，本部门员工默认只能看到员工自己
	OuterDept bool `json:"outerDept"`
	// 本部门的员工仅可见员工自己为true时，可以配置额外可见部门，值为部门id组成的的字符串，使用“\|”符号进行分割
	OuterPermitDepts string `json:"outerPermitDepts"`
	// 本部门的员工仅可见员工自己为true时，可以配置额外可见人员，值为userid组成的的字符串，使用“\|”符号进行分割
	OuterPermitUsers string `json:"outerPermitUsers"`
	// 企业群群主
	OrgDeptOwner string `json:"orgDeptOwner"`
	// 部门的主管列表，取值为由主管的userid组成的字符串，不同的userid使用“\|”符号进行分割
	DeptManagerUseridList string `json:"deptManagerUseridList"`
	// 部门标识字段，开发者可用该字段来唯一标识一个部门，并与钉钉外部通讯录里的部门做映射
	SourceIdentifier string `json:"sourceIdentifier"`
	// 部门群是否包含子部门
	GroupContainSubDept bool `json:"groupContainSubDept"`
	// 部门自定义字段
	Ext string `json:"ext"`
}

// 获取部门详情
func GetDepartment(accessToken string, id int64, lang *string) (department *Department, err error) {
	type Result struct {
		ding_sdk_golang.RespResult
		Department
	}
	url := fmt.Sprintf(arg.DEPARTMENTGET+"?access_token=%s&id=%d", accessToken, id)
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
	department = &result.Department
	return
}

type SimpleDepartment struct {
	// 部门id
	Id int64 `json:"id"`
	//部门名称
	Name string `json:"name"`
	// 父部门id，根部门为1
	Parentid int64 `json:"parentid"`
	// 是否同步创建一个关联此部门的企业群，true表示是，false表示不是
	CreateDeptGroup bool `json:"createDeptGroup"`
	// 当部门群已经创建后，是否有新人加入部门会自动加入该群，true表示是，false表示不是
	AutoAddUser bool `json:"autoAddUser"`
	// 部门自定义字段
	Ext string `json:"ext"`
}

// 获取部门列表
func DepartmentList(accessToken string, id *int64, lang *string, fetchChild *bool) (department []*SimpleDepartment, err error) {
	// ISV微应用固定传递false
	type Result struct {
		ding_sdk_golang.RespResult
		Department []*SimpleDepartment `json:"department"`
	}
	url := fmt.Sprintf(arg.DEPARTMENTLIST+"?access_token=%s", accessToken)
	if id != nil {
		url = fmt.Sprintf("%s&id=%d", url, *id)
	}
	if lang != nil {
		url = fmt.Sprintf("%s&lang=%s", url, *lang)
	}
	//if fetchChild != nil {
	//	url = fmt.Sprintf("%s&lang=%v", url, *fetchChild)
	//}
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
	department = result.Department
	return
}

// 获取子部门ID列表
func DepartmentListIds(accessToken string, id int64) (subDeptIdList []int64, err error) {
	// ISV微应用固定传递false
	type Result struct {
		ding_sdk_golang.RespResult
		SubDeptIdList []int64 `json:"sub_dept_id_list"`
	}
	var result Result
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(arg.DEPARTMENTLISTIDS+"?access_token=%s&id=%d", accessToken, id),
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
	subDeptIdList = result.SubDeptIdList
	return
}

// 查询指定用户的所有上级父部门路径
func DepartmentListParentDepts(accessToken string, userid string) (department [][]int64, err error) {
	type Result struct {
		ding_sdk_golang.RespResult
		Department [][]int64 `json:"department"`
	}
	var result Result
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(arg.DEPARTMENTLISTPARENTDEPTS+"?access_token=%s&userId=%s", accessToken, userid),
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
	department = result.Department
	return
}

// 查询部门的所有上级父部门路径
func DepartmentListParentDeptsByDept(accessToken string, id int64) (parentIds []int64, err error) {
	type Result struct {
		ding_sdk_golang.RespResult
		ParentIds []int64 `json:"parentIds"`
	}
	var result Result
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(arg.DEPARTMENTLISTPARENTDEPTSBYDEPT+"?access_token=%s&id=%d", accessToken, id),
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
	parentIds = result.ParentIds
	return
}
