package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// QueryMember 获取体验者列表
type MiniProgramQueryMemberRequest struct {
	Action string `json:"action"`
}

type MiniProgramQueryMemberResponse struct {
	util.CommonError
}

func (miniprogram *MiniProgram) QueryMember() (*MiniProgramQueryMemberResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	req := &MiniProgramQueryMemberRequest{Action: "get_experiencer"}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/memberauth?access_token=%s", accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramQueryMemberResponse{}
	err = util.DecodeWithError(response, resp, "QueryMember")
	return resp, err
}
