package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// AllCategory 获取可以设置的所有类目
type MiniProgramAllCategoryRequest struct {
	//
}

type MiniProgramAllCategoryResponse struct {
	util.CommonError
	CategoriesList struct {
		Categories []MiniProgramCategoryItem `json:"categories"`
	} `json:"categories_list"`
}

func (miniprogram *MiniProgram) AllCategory() (*MiniProgramAllCategoryResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/wxopen/getallcategories?access_token=%s", accessToken)
	response, err := util.HTTPGet(uri)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramAllCategoryResponse{}
	err = util.DecodeWithError(response, resp, "AllCategory")
	return resp, err
}

type MiniProgramCategoryItem struct {
	ID            int    `json:"id"`             // 类目 ID
	Name          string `json:"name"`           // 类目名称
	Level         int    `json:"level"`          // 类目层级
	Father        int    `json:"father"`         // 类目父级 ID
	Children      []int  `json:"children"`       // 子级类目 ID
	SensitiveType int    `json:"sensitive_type"` // 是否为敏感类目（1 为敏感类目，需要提供相应资质审核；0 为非敏感类目，无需审核）
	Qualify       struct {
		ExterList []struct {
			InnerList []MiniProgramCategoryInnerListItem `json:"inner_list"`
		} `json:"exter_list"`
	} `json:"qualify"` // 类目 ID
}

type MiniProgramCategoryInnerListItem struct {
	Name string `json:"name"` // 资质文件名称
	URL  string `json:"url"`  // 资质文件示例
}
