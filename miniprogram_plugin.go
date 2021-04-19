package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

type MiniProgramPluginAction string

const (
	MiniProgramPluginActionApply  MiniProgramPluginAction = "apply"  // 申请使用插件
	MiniProgramPluginActionList   MiniProgramPluginAction = "list"   // 查询已添加的插件列表
	MiniProgramPluginActionUnbind MiniProgramPluginAction = "unbind" // 删除已添加的插件
	MiniProgramPluginActionUpdate MiniProgramPluginAction = "update" // 快速更新插件版本号
)

// Plugin 插件
type MiniProgramPluginRequest struct {
	Action      MiniProgramPluginAction `json:"action"`
	PluginAppID string                  `json:"plugin_appid,omitempty"`
	UserVersion string                  `json:"user_version,omitempty"`
}

type MiniProgramPluginResponse struct {
	util.CommonError
	PluginList []MiniProgramPluginItem `json:"plugin_list"`
}

func (miniprogram *MiniProgram) Plugin(req *MiniProgramPluginRequest) (*MiniProgramPluginResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/plugin?access_token=%s", accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramPluginResponse{}
	err = util.DecodeWithError(response, resp, "Plugin")
	return resp, err
}

type MiniProgramPluginItem struct {
	AppID      string `json:"appid"`
	Status     string `json:"status"`
	Nickname   string `json:"nickname"`
	HeadImgURL string `json:"headimgurl"`
}
