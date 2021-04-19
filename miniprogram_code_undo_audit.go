package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// UndoAudit 小程序审核撤回
type MiniProgramUndoAuditRequest struct {
	//
}

type MiniProgramUndoAuditResponse struct {
	util.CommonError
}

func (miniprogram *MiniProgram) UndoAudit() (*MiniProgramUndoAuditResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/revertcoderelease?access_token=%s", accessToken)
	response, err := util.HTTPGet(uri)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramUndoAuditResponse{}
	err = util.DecodeWithError(response, resp, "UndoAudit")
	return resp, err
}
