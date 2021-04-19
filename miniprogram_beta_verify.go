package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// BetaVerify 试用小程序快速认证
type MiniProgramBetaVerifyRequest struct {
	VerifyInfo MiniProgramBetaVerifyInfo `json:"verify_info"` // 企业法人认证需要的信息
}

type MiniProgramBetaVerifyResponse struct {
	util.CommonError
}

func (miniprogram *MiniProgram) BetaVerify(req *MiniProgramBetaVerifyRequest) (*MiniProgramBetaVerifyResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/verifybetaweapp?access_token=%s", accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramBetaVerifyResponse{}
	err = util.DecodeWithError(response, resp, "BetaVerify")
	return resp, err
}

type MiniProgramBetaVerifyInfo struct {
	EnterpriseName     string `json:"enterprise_name"`      // 企业名（需与工商部门登记信息一致）
	Code               string `json:"code"`                 // 企业代码
	CodeType           string `json:"code_type"`            // 企业代码类型 1：统一社会信用代码（18 位） 2：组织机构代码（9 位 xxxxxxxx-x） 3：营业执照注册号(15 位)
	LegalPersonaWechat string `json:"legal_persona_wechat"` // 法人微信号
	LegalPersonaName   string `json:"legal_persona_name"`   // 法人姓名（绑定银行卡）
	LegalPersonaIDCard string `json:"legal_persona_idcard"` // 法人身份证号
	ComponentPhone     string `json:"component_phone"`      // 第三方联系电话
}
