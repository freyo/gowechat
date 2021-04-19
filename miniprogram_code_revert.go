package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// CodeRevert 版本回退
type MiniProgramCodeRevertRequest struct {
	//
}

type MiniProgramCodeRevertResponse struct {
	util.CommonError
}

func (miniprogram *MiniProgram) CodeRevert() (*MiniProgramCodeRevertResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/revertcoderelease?access_token=%s", accessToken)
	response, err := util.HTTPGet(uri)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramCodeRevertResponse{}
	err = util.DecodeWithError(response, resp, "CodeRevert")
	return resp, err
}
