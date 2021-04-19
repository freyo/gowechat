package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// AddCategory 添加类目
type MiniProgramAddCategoryRequest struct {
	Categories []MiniProgramCategory `json:"categories"` // 类目信息列表
}

type MiniProgramAddCategoryResponse struct {
	util.CommonError
}

func (miniprogram *MiniProgram) AddCategory(req *MiniProgramAddCategoryRequest) (*MiniProgramAddCategoryResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/wxopen/addcategory?access_token=%s", accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramAddCategoryResponse{}
	err = util.DecodeWithError(response, resp, "AddCategory")
	return resp, err
}

type MiniProgramCategory struct {
	First        int                              `json:"first"`      // 一级类目 ID
	Second       int                              `json:"second"`     // 二级类目 ID
	Certificates []MiniProgramCategoryCertificate `json:"certicates"` // 资质信息列表
}

type MiniProgramCategoryCertificate struct {
	Key   string `json:"key"`   // 资质名称
	Value string `json:"value"` // 资质图片
}
