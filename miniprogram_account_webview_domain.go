package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// WebViewDomain 设置业务域名
type MiniProgramWebViewDomainRequest struct {
	Action        MiniProgramDomainAction `json:"action,omitempty"`        // 操作类型
	WebViewDomain []string                `json:"webviewdomain,omitempty"` // 小程序业务域名，当 action 参数是 get 时不需要此字段
}

type MiniProgramWebViewDomainResponse struct {
	util.CommonError
	WebViewDomain []string `json:"webviewdomain"` // request 合法域名；当 action 是 get 时不需要此字段
}

func (miniprogram *MiniProgram) WebViewDomain(req *MiniProgramWebViewDomainRequest) (*MiniProgramWebViewDomainResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/setwebviewdomain?access_token=%s", accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramWebViewDomainResponse{}
	err = util.DecodeWithError(response, resp, "WebViewDomain")
	return resp, err
}
