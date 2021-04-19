package gowechat

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/silenceper/wechat/v2/util"
)

// CodePreview 获取体验版二维码
type MiniProgramCodePreviewRequest struct {
	Path string `json:"path"` // 指定二维码扫码后直接进入指定页面并可同时带上参数）
}

type MiniProgramCodePreviewResponse struct {
	util.CommonError
	QRCode []byte `json:"-"`
}

func (miniprogram *MiniProgram) CodePreview(req *MiniProgramCodePreviewRequest) (*MiniProgramCodePreviewResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/get_qrcode?access_token=%s", accessToken)
	response, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	contentType := response.Header.Get("Content-Type")
	resp := &MiniProgramCodePreviewResponse{}
	if strings.HasPrefix(contentType, "application/json") {
		err = util.DecodeWithError(responseData, resp, "CodePreview")
	} else if contentType == "image/jpeg" {
		resp.QRCode = responseData
	} else {
		err = fmt.Errorf("fetchCode error : unknown response content type - %v", contentType)
	}

	return resp, err
}
