package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// CheckNickName 微信认证名称检测
type MiniProgramCheckNickNameRequest struct {
	Nickname string `json:"nick_name"` // 名称（昵称）
}

type MiniProgramCheckNickNameResponse struct {
	util.CommonError
}

func (miniprogram *MiniProgram) CheckNickName(req *MiniProgramCheckNickNameRequest) (*MiniProgramCheckNickNameResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/wxverify/checkwxverifynickname?access_token=%s", accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramCheckNickNameResponse{}
	err = util.DecodeWithError(response, resp, "CheckNickName")
	return resp, err
}
