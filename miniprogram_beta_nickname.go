package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// BetaSetNickname 修改试用小程序名称
type MiniProgramBetaSetNicknameRequest struct {
	Name string `json:"name"` // 小程序名称，昵称半自动设定，强制后缀“的体验小程序”。且该参数会进行关键字检查，如果命中品牌关键字则会报错。 如遇到品牌大客户要用试用小程序，建议用户先换个名字，认证后再修改成品牌名。
}

type MiniProgramBetaSetNicknameResponse struct {
	util.CommonError
}

func (miniprogram *MiniProgram) BetaSetNickname(req *MiniProgramBetaSetNicknameRequest) (*MiniProgramBetaSetNicknameResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/setbetaweappnickname?access_token=%s", accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramBetaSetNicknameResponse{}
	err = util.DecodeWithError(response, resp, "BetaSetNickname")
	return resp, err
}
