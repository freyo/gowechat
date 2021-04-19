package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// SetSignature 设置简介
type MiniProgramSetSignatureRequest struct {
	Signature string `json:"signature"` // 功能介绍（简介）
}

type MiniProgramSetSignatureResponse struct {
	util.CommonError
}

func (miniprogram *MiniProgram) SetSignature(req *MiniProgramSetSignatureRequest) (*MiniProgramSetSignatureResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/account/modifysignature?access_token=%s", accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramSetSignatureResponse{}
	err = util.DecodeWithError(response, resp, "SetSignature")
	return resp, err
}
