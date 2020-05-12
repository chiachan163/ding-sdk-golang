package dingtalk

import (
	"fmt"
	"net/http"

	"github.com/chiachan163/ding-sdk-golang/v2/arg"

	"github.com/swxctx/ghttp"
)

// 获取部门详情
func GetDepartment(accessToken string, id int64, lang *string) (_result *arg.GetDepartmentResult, err error) {

	url := fmt.Sprintf(arg.DEPARTMENTGET+"?access_token=%s&id=%d", accessToken, id)
	if lang != nil {
		url = fmt.Sprintf("%s&lang=%s", url, *lang)
	}
	var result arg.GetDepartmentResult
	resp, err := ghttp.Request{
		Url:         url,
		Body:        nil,
		Method:      "GET",
		ContentType: "application/json",
	}.Do()
	if err != nil {
		return
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

// 获取部门列表
func DepartmentList(accessToken string, id *int64, lang *string, fetchChild *bool) (_result *arg.DepartmentListResult, err error) {
	// ISV微应用固定传递false
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
	var result arg.DepartmentListResult
	resp, err := ghttp.Request{
		Url:         url,
		Body:        nil,
		Method:      "GET",
		ContentType: "application/json",
	}.Do()
	if err != nil {
		return
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

// 获取子部门ID列表
func DepartmentListIds(accessToken string, id int64) (_result *arg.DepartmentListIdsResult, err error) {
	// ISV微应用固定传递false
	var result arg.DepartmentListIdsResult
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(arg.DEPARTMENTLISTIDS+"?access_token=%s&id=%d", accessToken, id),
		Body:        nil,
		Method:      "GET",
		ContentType: "application/json",
	}.Do()
	if err != nil {
		return
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

// 查询指定用户的所有上级父部门路径
func DepartmentListParentDepts(accessToken string, userid string) (_result *arg.DepartmentListParentDeptsResult, err error) {
	var result arg.DepartmentListParentDeptsResult
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(arg.DEPARTMENTLISTPARENTDEPTS+"?access_token=%s&userId=%s", accessToken, userid),
		Body:        nil,
		Method:      "GET",
		ContentType: "application/json",
	}.Do()
	if err != nil {
		return
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

// 查询部门的所有上级父部门路径
func DepartmentListParentDeptsByDept(accessToken string, id int64) (_result *arg.DepartmentListParentDeptsByDeptResult, err error) {

	var result arg.DepartmentListParentDeptsByDeptResult
	resp, err := ghttp.Request{
		Url:         fmt.Sprintf(arg.DEPARTMENTLISTPARENTDEPTSBYDEPT+"?access_token=%s&id=%d", accessToken, id),
		Body:        nil,
		Method:      "GET",
		ContentType: "application/json",
	}.Do()
	if err != nil {
		return
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
