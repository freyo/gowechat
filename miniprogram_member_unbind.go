package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// UnbindMember 解除绑定体验者
type MiniProgramUnbindMemberRequest struct {
	WeChatID string `json:"wechatid,omitempty"` // 微信号
	UserStr  string `json:"userstr,omitempty"`  // 人员对应的唯一字符串， 可通过获取已绑定的体验者列表获取人员对应的字符串
}

type MiniProgramUnbindMemberResponse struct {
	util.CommonError
}

func (miniprogram *MiniProgram) UnbindMember(req *MiniProgramUnbindMemberRequest) (*MiniProgramUnbindMemberResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/unbind_tester?access_token=%s", accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramUnbindMemberResponse{}
	err = util.DecodeWithError(response, resp, "UnbindMember")
	return resp, err
}
