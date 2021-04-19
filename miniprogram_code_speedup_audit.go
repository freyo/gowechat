package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// SpeedUpAudit 加急审核申请
type MiniProgramSpeedUpAuditRequest struct {
	AuditID int `json:"auditid"` // 审核单ID
}

type MiniProgramSpeedUpAuditResponse struct {
	util.CommonError
}

func (miniprogram *MiniProgram) SpeedUpAudit(req *MiniProgramSpeedUpAuditRequest) (*MiniProgramSpeedUpAuditResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/speedupaudit?access_token=%s", accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramSpeedUpAuditResponse{}
	err = util.DecodeWithError(response, resp, "SpeedUpAudit")
	return resp, err
}
