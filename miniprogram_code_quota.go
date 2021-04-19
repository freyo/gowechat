package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// QueryQuota 查询服务商的当月提审限额（quota）和加急次数
type MiniProgramQueryQuotaRequest struct {
	//
}

type MiniProgramQueryQuotaResponse struct {
	util.CommonError
	Rest         int `json:"rest"`          // quota剩余值
	Limit        int `json:"limit"`         // 当月分配quota
	SpeedUpRest  int `json:"speedup_rest"`  // 剩余加急次数
	SpeedUpLimit int `json:"speedup_limit"` // 当月分配加急次数
}

func (miniprogram *MiniProgram) QueryQuota() (*MiniProgramQueryQuotaResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/queryquota?access_token=%s", accessToken)
	response, err := util.HTTPGet(uri)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramQueryQuotaResponse{}
	err = util.DecodeWithError(response, resp, "QueryQuota")
	return resp, err
}
