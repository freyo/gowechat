package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// CodeCommit 上传小程序代码
type MiniProgramCodeCommitRequest struct {
	TemplateID  int    `json:"template_id"`  // 代码库中的代码模板 ID
	ExtJSON     string `json:"ext_json"`     // 第三方自定义的配置
	UserVersion string `json:"user_version"` // 代码版本号，开发者可自定义（长度不要超过 64 个字符）
	UserDesc    string `json:"user_desc"`    // 代码描述，开发者可自定义
}

type MiniProgramCodeCommitResponse struct {
	util.CommonError
}

func (miniprogram *MiniProgram) CodeCommit(req *MiniProgramCodeCommitRequest) (*MiniProgramCodeCommitResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/commit?access_token=%s", accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramCodeCommitResponse{}
	err = util.DecodeWithError(response, resp, "CodeCommit")
	return resp, err
}

type MiniProgramExtJSON struct {
	ExtAppID       string                   `json:"extAppid"`
	Ext            map[string]string        `json:"ext"`
	ExtPages       map[string]interface{}   `json:"extPages"`
	Pages          []string                 `json:"pages"`
	Window         map[string]string        `json:"window"`
	NetworkTimeout map[string]interface{}   `json:"networkTimeout"`
	TabBar         map[string][]interface{} `json:"tabBar"`
	Plugins        []string                 `json:"plugins"`
}
