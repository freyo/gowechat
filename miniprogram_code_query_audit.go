package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// QueryAudit 查询指定发布审核单的审核状态
type MiniProgramQueryAuditRequest struct {
	AuditID int `json:"auditid"` // 提交审核时获得的审核 id
}

type MiniProgramQueryAuditResponse struct {
	util.CommonError
	Status     int    `json:"status,omitempty"`
	Reason     string `json:"reason,omitempty"`
	Screenshot string `json:"screenshot,omitempty"`
}

func (miniprogram *MiniProgram) QueryAudit(req *MiniProgramQueryAuditRequest) (*MiniProgramQueryAuditResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/get_auditstatus?access_token=%s", accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramQueryAuditResponse{}
	err = util.DecodeWithError(response, resp, "QueryAudit")
	return resp, err
}
