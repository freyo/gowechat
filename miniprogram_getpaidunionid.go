package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const miniProgramgetpaidunionidURL = "https://api.weixin.qq.com/wxa/getpaidunionid"

// MiniProgramGetPaidUnionidRequest 小程序 unionid
type MiniProgramGetPaidUnionidRequest struct {
	AppID         string `json:"appid" form:"appid" binding:"required"`   // appid
	Openid        string `json:"openid" form:"openid" binding:"required"` // 是	支付用户唯一标识
	TransactionID string `json:"transaction_id"  form:"transaction_id"`   // 否	微信订单号
	MchID         string `json:"mch_id"  form:"mch_id"`                   // 否	商户号，和商户订单号配合使用
	OutTradeNo    string `json:"out_trade_no"  form:"out_trade_no"`       // 否	商户订单号，和商户号配合使用
}

// MiniProgramGetPaidUnionidResponse 小程序 unionid
type MiniProgramGetPaidUnionidResponse struct {
	util.CommonError
	Unionid string `json:"unionid"`
}

// GetPaidUnionid 获取支付后的用户 unionid
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/User_Management.html
func (miniprogram *MiniProgram) GetPaidUnionid(req *MiniProgramGetPaidUnionidRequest) (*MiniProgramGetPaidUnionidResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	var url string

	// 如果交易单号不为空，我们使用 微信订单号 这种方式获取用户 unionid
	if req.TransactionID != "" {
		url = fmt.Sprintf("%s?access_token=%s&openid=%s&transaction_id=%s", miniProgramgetpaidunionidURL, accessToken, req.Openid, req.TransactionID)
	} else {
		url = fmt.Sprintf(
			"%s?access_token=%s&openid=%s&mch_id=%s&out_trade_no=%s",
			miniProgramgetpaidunionidURL, accessToken, req.Openid, req.MchID, req.OutTradeNo,
		)
	}

	response, err := util.HTTPGet(url)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramGetPaidUnionidResponse{}
	err = util.DecodeWithError(response, resp, "GetPaidUnionid")
	return resp, err
}
