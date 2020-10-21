package feishu

import (
	"github.com/shupkg/feishu/core/consts"
	"github.com/shupkg/feishu/core/model/vo"
	"github.com/shupkg/feishu/core/util/http"
	"github.com/shupkg/feishu/core/util/json"
)

//校验应用管理员 https://open.feishu.cn/document/ukTMukTMukTM/uITN1EjLyUTNx4iM1UTM
func (t Tenant) IsUserAdmin(openId string, employeeId string) (*vo.IsUserAdminResp, error){
	queryParams := map[string]interface{}{
	}
	if openId != ""{
		queryParams["open_id"] = openId
	}
	if employeeId != ""{
		queryParams["employee_id"] = employeeId
	}
	respBody, err := http.Get(consts.ApiIsUserAdmin, queryParams, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil{
		return nil, err
	}
	respVo := &vo.IsUserAdminResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}
