package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

type MiniProgramDomainAction string

const (
	MiniProgramDomainActionAdd    MiniProgramDomainAction = "add"    // 添加
	MiniProgramDomainActionDelete MiniProgramDomainAction = "delete" // 删除
	MiniProgramDomainActionSet    MiniProgramDomainAction = "set"    // 覆盖
	MiniProgramDomainActionGet    MiniProgramDomainAction = "get"    // 获取
)

// Domain 设置服务器域名
type MiniProgramDomainRequest struct {
	Action         MiniProgramDomainAction `json:"action"`                    // 操作类型
	RequestDomain  []string                `json:"requestdomain,omitempty"`   // request 合法域名；当 action 是 get 时不需要此字段
	SocketDomain   []string                `json:"wsrequestdomain,omitempty"` // socket 合法域名；当 action 是 get 时不需要此字段
	UploadDomain   []string                `json:"uploaddomain,omitempty"`    // uploadFile 合法域名；当 action 是 get 时不需要此字段
	DownloadDomain []string                `json:"downloaddomain,omitempty"`  // downloadFile 合法域名；当 action 是 get 时不需要此字段
}

type MiniProgramDomainResponse struct {
	util.CommonError
	RequestDomain  []string `json:"requestdomain"`   // request 合法域名；当 action 是 get 时不需要此字段
	SocketDomain   []string `json:"wsrequestdomain"` // socket 合法域名；当 action 是 get 时不需要此字段
	UploadDomain   []string `json:"uploaddomain"`    // uploadFile 合法域名；当 action 是 get 时不需要此字段
	DownloadDomain []string `json:"downloaddomain"`  // downloadFile 合法域名；当 action 是 get 时不需要此字段
}

func (miniprogram *MiniProgram) Domain(req *MiniProgramDomainRequest) (*MiniProgramDomainResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/modify_domain?access_token=%s", accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramDomainResponse{}
	err = util.DecodeWithError(response, resp, "Domain")
	return resp, err
}
