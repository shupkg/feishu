package feishu

import (
	"github.com/shupkg/feishu/core/consts"
	"github.com/shupkg/feishu/core/model/vo"
	"github.com/shupkg/feishu/core/util/http"
	"github.com/shupkg/feishu/core/util/json"
)

//搜索用户 https://bytedance.feishu.cn/docs/doccnizryz7NKuUmVfkRJWeZGVc
func (u User) SearchUser(query string, pageSize int, pageToken string) (*vo.SearchUserResp, error){
	queryParams := map[string]interface{}{
	}
	if query != ""{
		queryParams["query"] = query
	}
	if pageSize > 0{
		queryParams["page_size"] = pageSize
	}
	if pageToken != ""{
		queryParams["page_token"] = pageToken
	}
	respBody, err := http.Get(consts.ApiSearchUser, queryParams, http.BuildTokenHeaderOptions(u.UserAccessToken))
	if err != nil{
		return nil, err
	}
	respVo := &vo.SearchUserResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}
