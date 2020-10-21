package feishu

import (
	"github.com/shupkg/feishu/core/consts"
	"github.com/shupkg/feishu/core/model/vo"
	"github.com/shupkg/feishu/core/util/http"
	"github.com/shupkg/feishu/core/util/json"
)

//发送消息卡片 https://open.feishu.cn/document/ukTMukTMukTM/uYTNwUjL2UDM14iN1ATN
func (t Tenant) SendMessage(msg vo.MsgVo) (*vo.MsgResp, error) {
	reqBody := json.ToJsonIgnoreError(msg)
	
	respBody, err := http.Post(consts.ApiRobotSendMessage, nil, reqBody, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		return nil, err
	}
	respVo := &vo.MsgResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//发送消息卡片 https://open.feishu.cn/document/ukTMukTMukTM/uYTNwUjL2UDM14iN1ATN
func (t Tenant) SendMessageBatch(msg vo.BatchMsgVo) (*vo.MsgResp, error) {
	reqBody := json.ToJsonIgnoreError(msg)
	respBody, err := http.Post(consts.ApiRobotSendBatchMessage, nil, reqBody, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		return nil, err
	}
	respVo := &vo.MsgResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}
