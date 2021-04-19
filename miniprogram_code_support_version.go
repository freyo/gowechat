package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// SupportVersion 查询服务商的当月提审限额（quota）和加急次数
type MiniProgramSupportVersionRequest struct {
	//
}

type MiniProgramSupportVersionResponse struct {
	util.CommonError
	NowVersion int `json:"now_version"` // 当前版本
	UVInfo     struct {
		Items []MiniProgramUVInfoItem `json:"items"`
	} `json:"uv_info"` // 版本的用户占比列表
}

func (miniprogram *MiniProgram) SupportVersion() (*MiniProgramSupportVersionResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/wxopen/getweappsupportversion?access_token=%s", accessToken)
	response, err := util.HTTPPost(uri, "{}")
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramSupportVersionResponse{}
	err = util.DecodeWithError(response, resp, "SupportVersion")
	return resp, err
}

type MiniProgramUVInfoItem struct {
	Percentage float64 `json:"percentage"` // 百分比
	Version    string  `json:"version"`    // 基础库版本号
}
