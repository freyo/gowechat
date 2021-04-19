package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// URLScheme
type MiniProgramURLSchemeRequest struct {
	Jump      MiniProgramURLSchemeJump `json:"jump_wxa"`              // 跳转到的目标小程序信息。
	Temporary bool                     `json:"is_expire"`             // 生成的scheme码类型，到期失效：true，永久有效：false。
	Timestamp int                      `json:"expire_time,omitempty"` // 到期失效的scheme码的失效时间，为Unix时间戳。生成的到期失效scheme码在该时间前有效。生成到期失效的scheme时必填。
}

type MiniProgramURLSchemeResponse struct {
	util.CommonError
	OpenLink string `json:"openlink"`
}

func (miniprogram *MiniProgram) URLScheme(req *MiniProgramURLSchemeRequest) (*MiniProgramURLSchemeResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/generatescheme?access_token=%s", accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramURLSchemeResponse{}
	err = util.DecodeWithError(response, resp, "URLScheme")
	return resp, err
}

type MiniProgramURLSchemeJump struct {
	Path  string `json:"path"`
	Query string `json:"query,omitempty"`
}
