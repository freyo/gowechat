package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// SetHeadImage 设置头像
type MiniProgramSetHeadImageRequest struct {
	HeadImageMediaID string `json:"head_img_media_id"` // 头像素材 media_id
	X1               string `json:"x1"`                // 裁剪框左上角 x 坐标（取值范围：[0, 1]）
	Y1               string `json:"y1"`                // 裁剪框左上角 y 坐标（取值范围：[0, 1]）
	X2               string `json:"x2"`                // 裁剪框右下角 x 坐标（取值范围：[0, 1]）
	Y2               string `json:"y2"`                // 裁剪框右下角 y 坐标（取值范围：[0, 1]）
}

type MiniProgramSetHeadImageResponse struct {
	util.CommonError
}

func (miniprogram *MiniProgram) SetHeadImage(req *MiniProgramSetHeadImageRequest) (*MiniProgramSetHeadImageResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/account/modifyheadimage?access_token=%s", accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramSetHeadImageResponse{}
	err = util.DecodeWithError(response, resp, "SetHeadImage")
	return resp, err
}
