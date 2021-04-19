package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// BetaRegister 创建试用小程序
type MiniProgramBetaRegisterRequest struct {
	Name   string `json:"name"`   // 小程序名称，昵称半自动设定，强制后缀“的体验小程序”。且该参数会进行关键字检查，如果命中品牌关键字则会报错。 如遇到品牌大客户要用试用小程序，建议用户先换个名字，认证后再修改成品牌名。
	OpenID string `json:"openid"` // 微信用户的openid，试用小程序创建成功后会默认将该用户设置为小程序管理员
}

type MiniProgramBetaRegisterResponse struct {
	util.CommonError
	UniqueID     string `json:"unique_id"`     // 该请求的唯一标识符，用于关联微信用户和后面产生的appid
	AuthorizeUrl string `json:"authorize_url"` // 用户授权确认url，需将该url发送给用户，用户进入授权页面完成授权方可创建小程序
}

func (miniprogram *MiniProgram) BetaRegister(req *MiniProgramBetaRegisterRequest) (*MiniProgramBetaRegisterResponse, error) {
	componentAK, err := miniprogram.OpenPlatformMiniProgram.GetComponent().GetComponentAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/component/fastregisterbetaweapp?access_token=%s", componentAK)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramBetaRegisterResponse{}
	err = util.DecodeWithError(response, resp, "BetaRegister")
	return resp, err
}
