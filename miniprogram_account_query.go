package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
	"github.com/sirupsen/logrus"
)

// QueryAccount 获取基本信息
type MiniProgramQueryAccountRequest struct {
	//
}

type MiniProgramQueryAccountResponse struct {
	util.CommonError
	AppID             string                   `json:"appid"`           // 帐号 APPID
	AccountType       int                      `json:"account_type"`    // 帐号类型 1订阅号 2服务号 3小程序
	PrincipalType     int                      `json:"principal_type"`  // 主体类型 0个人 1企业 2媒体 3政府 4其他组织
	PrincipalName     string                   `json:"principal_name"`  // 主体名称
	RealNameStatus    int                      `json:"realname_status"` // 实名认证状态 1实名验证成功 2实名验证中 3实名验证失败
	WxVerifyInfo      MiniProgramVerifyInfo    `json:"wx_verify_info"`
	SignatureInfo     MiniProgramSignatureInfo `json:"signature_info"`
	HeadImageInfo     MiniProgramHeadImageInfo `json:"head_image_info"`
	NicknameInfo      MiniProgramNicknameInfo  `json:"nickname_info"`
	RegisteredCountry int                      `json:"registered_country"` // 注册国家 1017中国
}

func (miniprogram *MiniProgram) QueryAccount() (*MiniProgramQueryAccountResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	logrus.Debugf("access_token=%s error=%v", accessToken, err)
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/account/getaccountbasicinfo?access_token=%s", accessToken)
	response, err := util.HTTPGet(uri)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramQueryAccountResponse{}
	err = util.DecodeWithError(response, resp, "QueryAccount")
	return resp, err
}

type MiniProgramVerifyInfo struct {
	QualificationVerify   bool `json:"qualification_verify"`     // 是否资质认证，若是，拥有微信认证相关的权限。
	NamingVerify          bool `json:"naming_verify"`            // 是否名称认证
	AnnualReview          bool `json:"annual_review"`            // 是否需要年审（qualification_verify == true 时才有该字段）
	AnnualReviewBeginTime int  `json:"annual_review_begin_time"` // 年审开始时间，时间戳（qualification_verify == true 时才有该字段）
	AnnualReviewEndTime   int  `json:"annual_review_end_time"`   // 年审截止时间，时间戳（qualification_verify == true 时才有该字段）
}

type MiniProgramSignatureInfo struct {
	Signature       string `json:"signature"`         // 功能介绍
	ModifyUsedCount int    `json:"modify_used_count"` // 功能介绍已使用修改次数（本月）
	ModifyQuota     int    `json:"modify_quota"`      // 功能介绍修改次数总额度（本月）
}

type MiniProgramHeadImageInfo struct {
	HeadImageUrl    string `json:"head_image_url"`    // 头像 url
	ModifyUsedCount int    `json:"modify_used_count"` // 头像已使用修改次数（本月）
	ModifyQuota     int    `json:"modify_quota"`      // 头像修改次数总额度（本月）
}

type MiniProgramNicknameInfo struct {
	Nickname        string `json:"nickname"`          // 小程序名称
	ModifyUsedCount int    `json:"modify_used_count"` // 小程序名称已使用修改次数（本年）
	ModifyQuota     int    `json:"modify_quota"`      // 小程序名称修改次数总额度（本年）
}
