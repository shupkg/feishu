package feishu

import (
	"net/url"

	"github.com/shupkg/feishu/core/consts"
	"github.com/shupkg/feishu/core/model/vo"
	"github.com/shupkg/feishu/core/util/http"
	"github.com/shupkg/feishu/core/util/json"
	"github.com/shupkg/feishu/core/util/log"
)

//获取用户所在的群列表 https://open.feishu.cn/document/ukTMukTMukTM/uQzMwUjL0MDM14CNzATN
func (t Tenant) GroupList(userAccessToken string, pageSize int, pageToken string) (*vo.GroupListRespVo, error) {
	queryParams := map[string]interface{}{}
	if pageSize > 0 {
		queryParams["page_size"] = pageSize
	}
	if pageToken != "" {
		queryParams["page_token"] = pageToken
	}
	respBody, err := http.Get(consts.ApiUserGroupLIst, queryParams, http.BuildTokenHeaderOptions(userAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GroupListRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取群成员列表 https://open.feishu.cn/document/ukTMukTMukTM/uUzMwUjL1MDM14SNzATN
func (t Tenant) ChatMembers(userAccessToken string, chatId string, pageSize int, pageToken string) (*vo.ChatMembersRespVo, error) {
	queryParams := map[string]interface{}{
		"chat_id": chatId,
	}
	if pageSize > 0 {
		queryParams["page_size"] = pageSize
	}
	if pageToken != "" {
		queryParams["page_token"] = pageToken
	}
	respBody, err := http.Get(consts.ApiChatMembers, queryParams, http.BuildTokenHeaderOptions(userAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.ChatMembersRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//搜索用户所在的群列表 https://open.feishu.cn/document/ukTMukTMukTM/uUTOyUjL1kjM14SN5ITN
func (t Tenant) ChatSearch(userAccessToken string, query string, pageSize int, pageToken string) (*vo.GroupListRespVo, error) {
	queryParams := map[string]interface{}{
		"query": url.QueryEscape(query),
	}
	if pageSize > 0 {
		queryParams["page_size"] = pageSize
	}
	if pageToken != "" {
		queryParams["page_token"] = pageToken
	}
	respBody, err := http.Get(consts.ApiChatSearch, queryParams, http.BuildTokenHeaderOptions(userAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GroupListRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}
