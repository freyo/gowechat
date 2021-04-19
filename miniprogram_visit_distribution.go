package gowechat

import (
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

// 获取小程序访问分布数据
type MiniProgramVisitDistributionRequest struct {
	BeginDate string `json:"begin_date"` // 开始日期。格式为 yyyymmdd
	EndDate   string `json:"end_date"`   // 结束日期，限定查询 1 天数据，允许设置的最大值为昨日。格式为 yyyymmdd
}

type MiniProgramVisitDistributionResponse struct {
	util.CommonError
}

// 获取小程序访问分布数据
func (miniprogram *MiniProgram) VisitDistribution(req *MiniProgramVisitDistributionRequest) (*MiniProgramVisitDistributionResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("https://api.weixin.qq.com/datacube/getweanalysisappidvisitdistribution?access_token=%s", accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}
	resp := &MiniProgramVisitDistributionResponse{}
	err = util.DecodeWithError(response, resp, "VisitDistribution")
	return resp, err
}
