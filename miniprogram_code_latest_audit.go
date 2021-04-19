package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// LatestAudit 查询最新一次提交的审核状态
type MiniProgramLatestAuditRequest struct {
	//
}

type MiniProgramLatestAuditResponse struct {
	util.CommonError
	AuditID    int    `json:"auditid,omitempty"`    // 最新的审核 ID
	Status     int    `json:"status,omitempty"`     // 审核状态
	Reason     string `json:"reason,omitempty"`     // 当审核被拒绝时，返回的拒绝原因
	ScreenShot string `json:"ScreenShot,omitempty"` // 当审核被拒绝时，会返回审核失败的小程序截图示例。用 | 分隔的 media_id 的列表，可通过获取永久素材接口拉取截图内容
}

func (miniprogram *MiniProgram) LatestAudit() (*MiniProgramLatestAuditResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/get_latest_auditstatus?access_token=%s", accessToken)
	response, err := util.HTTPGet(uri)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramLatestAuditResponse{}
	err = util.DecodeWithError(response, resp, "LatestAudit")
	return resp, err
}
