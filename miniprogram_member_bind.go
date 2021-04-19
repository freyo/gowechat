package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// BindMember 绑定微信用户为体验者
type MiniProgramBindMemberRequest struct {
	WeChatID string `json:"wechatid"` // 微信号
}

type MiniProgramBindMemberResponse struct {
	util.CommonError
}

func (miniprogram *MiniProgram) BindMember(req *MiniProgramBindMemberRequest) (*MiniProgramBindMemberResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/bind_tester?access_token=%s", accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramBindMemberResponse{}
	err = util.DecodeWithError(response, resp, "DeleteCategory")
	return resp, err
}
