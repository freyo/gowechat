package gowechat

import (
	"fmt"
	"github.com/silenceper/wechat/v2/util"
	"net/url"
)

// 复用公众号主体快速注册小程序
// copyWxVerify	是否复用公众号的资质进行微信认证(1:申请复用资质进行微信认证 0:不申请)
// redirectURI 需和第三方平台在微信开放平台上面填写的登录授权的发起页域名一致
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Register_Mini_Programs/fast_registration_of_mini_program.html
func (o *OfficialAccount) RegisterMiniProgramPage(copyWxVerify int, redirectURI string) string {
	return fmt.Sprintf("https://mp.weixin.qq.com/cgi-bin/fastregisterauth?appid=%s&component_appid=%s&copy_wx_verify=%d&redirect_uri=%s",
		o.AppID, o.OpenPlatformOfficialAccount.GetContext().AppID, copyWxVerify, url.QueryEscape(redirectURI))
}

// RegisterMiniProgram 复用公众号主体快速注册小程序
type OfficialAccountRegisterMiniProgramRequest struct {
	Ticket string `json:"ticket"` // 公众号扫码授权的凭证
}

type OfficialAccountRegisterMiniProgramResponse struct {
	util.CommonError
	AppID             string `json:"appid"`              // 新创建小程序的appid
	AuthorizationCode string `json:"authorization_code"` // 新创建小程序的授权码
	IsWxVerifySucc    string `json:"is_wx_verify_succ"`  // 复用公众号微信认证小程序是否成功
	IsLinkSucc        string `json:"is_link_succ"`       // 小程序是否和公众号关联成功
}

func (o *OfficialAccount) RegisterMiniProgram(req *OfficialAccountRegisterMiniProgramRequest) (*OfficialAccountRegisterMiniProgramResponse, error) {
	accessToken, err := o.OpenPlatformOfficialAccount.OfficialAccount.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/account/fastregister?access_token=%s", accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	resp := &OfficialAccountRegisterMiniProgramResponse{}
	err = util.DecodeWithError(response, resp, "RegisterMiniProgram")
	return resp, err
}

// 换绑小程序管理员
// redirectURI 需和第三方平台在微信开放平台上面填写的登录授权的发起页域名一致
func (o *OfficialAccount) ChangeMiniProgramAdminPage(redirectURI string) string {
	return fmt.Sprintf("https://mp.weixin.qq.com/wxopen/componentrebindadmin?appid=%s&component_appid=%s&redirect_uri=%s",
		o.AppID, o.OpenPlatformOfficialAccount.GetContext().AppID, url.QueryEscape(redirectURI))
}

type OfficialAccountChangeMiniProgramAdminRequest struct {
	TaskID string `json:"taskid"` // 换绑管理员任务序列号
}

type OfficialAccountChangeMiniProgramAdminResponse struct {
	util.CommonError
}

func (o *OfficialAccount) ChangeMiniProgramAdmin(req *OfficialAccountChangeMiniProgramAdminRequest) (*OfficialAccountChangeMiniProgramAdminResponse, error) {
	accessToken, err := o.OpenPlatformOfficialAccount.OfficialAccount.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/account/componentrebindadmin?access_token=%s", accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	resp := &OfficialAccountChangeMiniProgramAdminResponse{}
	err = util.DecodeWithError(response, resp, "ChangeMiniProgramAdmin")
	return resp, err
}
