package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// DeleteCategory 删除类目
type MiniProgramDeleteCategoryRequest struct {
	First  int `json:"first"`  // 一级类目 ID
	Second int `json:"second"` // 二级类目 ID
}

type MiniProgramDeleteCategoryResponse struct {
	util.CommonError
}

func (miniprogram *MiniProgram) DeleteCategory(req *MiniProgramDeleteCategoryRequest) (*MiniProgramDeleteCategoryResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/wxopen/deletecategory?access_token=%s", accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramDeleteCategoryResponse{}
	err = util.DecodeWithError(response, resp, "DeleteCategory")
	return resp, err
}
