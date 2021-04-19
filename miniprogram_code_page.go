package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// CodePage 获取已上传的代码的页面列表
type MiniProgramCodePageRequest struct {
	//
}

type MiniProgramCodePageResponse struct {
	util.CommonError
	PageList []string `json:"page_list"` // 页面配置列表
}

func (miniprogram *MiniProgram) CodePage() (*MiniProgramCodePageResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/get_page?access_token=%s", accessToken)
	response, err := util.HTTPGet(uri)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramCodePageResponse{}
	err = util.DecodeWithError(response, resp, "CodePage")
	return resp, err
}
