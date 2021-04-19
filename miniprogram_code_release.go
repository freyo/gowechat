package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// CodeRelease 发布已通过审核的小程序
type MiniProgramCodeReleaseRequest struct {
	//
}

type MiniProgramCodeReleaseResponse struct {
	util.CommonError
}

func (miniprogram *MiniProgram) CodeRelease() (*MiniProgramCodeReleaseResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/release?access_token=%s", accessToken)
	response, err := util.HTTPPost(uri, "{}")
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramCodeReleaseResponse{}
	err = util.DecodeWithError(response, resp, "CodeRelease")
	return resp, err
}
