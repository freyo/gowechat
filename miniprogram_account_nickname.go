package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// SetNickname 设置名称
type MiniProgramSetNicknameRequest struct {
	Nickname string `json:"nick_name"`         // 昵称，不支持包含“小程序”关键字的昵称
	IDCard   string `json:"id_card,omitempty"` // 个人号必填 身份证照片 mediaid
	License  string `json:"license,omitempty"` // 组织号必填	组织机构代码证或营业执照 mediaid
	// 其他证明材料 mediaid
	NamingOtherStuff1 string `json:"naming_other_stuff_1,omitempty"`
	NamingOtherStuff2 string `json:"naming_other_stuff_2,omitempty"`
	NamingOtherStuff3 string `json:"naming_other_stuff_3,omitempty"`
	NamingOtherStuff4 string `json:"naming_other_stuff_4,omitempty"`
	NamingOtherStuff5 string `json:"naming_other_stuff_5,omitempty"`
}

type MiniProgramSetNicknameResponse struct {
	util.CommonError
}

func (miniprogram *MiniProgram) SetNickname(req *MiniProgramSetNicknameRequest) (*MiniProgramSetNicknameResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/setnickname?access_token=%s", accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramSetNicknameResponse{}
	err = util.DecodeWithError(response, resp, "SetNickname")
	return resp, err
}
