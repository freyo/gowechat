package gowechat

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/officialaccount/material"
	"github.com/silenceper/wechat/v2/util"
)

// MediaUpload 临时素材上传
func (miniprogram *MiniProgram) MediaUpload(mediaType material.MediaType, filename string) (media material.Media, err error) {
	var accessToken string
	accessToken, err = miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s", accessToken, mediaType)
	var response []byte
	response, err = util.PostFile("media", filename, uri)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &media)
	if err != nil {
		return
	}
	if media.ErrCode != 0 {
		err = fmt.Errorf("MediaUpload error : errcode=%v , errmsg=%v", media.ErrCode, media.ErrMsg)
		return
	}
	return
}

// GetMediaURL 返回临时素材的下载地址供用户自己处理
// NOTICE: URL 不可公开，因为含access_token 需要立即另存文件
func (miniprogram *MiniProgram) GetMediaURL(mediaID string) (mediaURL string, err error) {
	var accessToken string
	accessToken, err = miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return
	}
	mediaURL = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s", accessToken, mediaID)
	return
}
