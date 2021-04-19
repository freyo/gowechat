package gowechat

import (
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

//获取小程序访问分布数据
type MiniProgramVisitPageRequest struct {
	BeginDate string `json:"begin_date"` //开始日期。格式为 yyyymmdd
	EndDate   string `json:"end_date"`   //结束日期，限定查询 1 天数据，允许设置的最大值为昨日。格式为 yyyymmdd
}

type MiniProgramVisitPageResponse struct {
	util.CommonError
}

//获取小程序访问分布数据
func (miniprogram *MiniProgram) VisitPage(req *MiniProgramVisitPageRequest) (*MiniProgramVisitPageResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("https://api.weixin.qq.com/datacube/getweanalysisappidvisitpage?access_token=%s", accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramVisitPageResponse{}
	err = util.DecodeWithError(response, resp, "VisitDistribution")
	return resp, err
}
