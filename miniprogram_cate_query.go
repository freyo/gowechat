package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// QueryCategory 获取已设置的所有类目
type MiniProgramQueryCategoryRequest struct {
	//
}

type MiniProgramQueryCategoryResponse struct {
	util.CommonError
	Categories    []MiniProgramQueryCategoryItem `json:"categories"`     // 已设置的类目信息列表
	Limit         int                            `json:"limit"`          // 一个更改周期内可以添加类目的次数
	Quota         int                            `json:"quota"`          // 本更改周期内还可以添加类目的次数
	CategoryLimit int                            `json:"category_limit"` // 最多可以设置的类目数量
}

func (miniprogram *MiniProgram) QueryCategory() (*MiniProgramQueryCategoryResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/wxopen/getcategory?access_token=%s", accessToken)
	response, err := util.HTTPGet(uri)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramQueryCategoryResponse{}
	err = util.DecodeWithError(response, resp, "QueryCategory")
	return resp, err
}

type MiniProgramQueryCategoryItem struct {
	First       int    `json:"first"`        // 一级类目 ID
	FirstName   string `json:"first_name"`   // 一级类目名称
	Second      int    `json:"second"`       // 二级类目 ID
	SecondName  string `json:"second_name"`  // 二级类目名称
	AuditStatus int    `json:"audit_status"` // 审核状态（1 审核中 2 审核不通过 3 审核通过）
	AuditReason string `json:"audit_reason"` // 审核不通过的原因
}
