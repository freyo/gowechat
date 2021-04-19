package gowechat

import (
	"github.com/silenceper/wechat/v2/credential"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniProgramConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	opContext "github.com/silenceper/wechat/v2/openplatform/context"
	openPlatformMiniProgram "github.com/silenceper/wechat/v2/openplatform/miniprogram"
	"github.com/sirupsen/logrus"
)

// MiniProgram 代小程序实现业务
type MiniProgram struct {
	AppID                   string
	OpenPlatformMiniProgram *openPlatformMiniProgram.MiniProgram
	MiniProgram             *miniprogram.MiniProgram
}

// NewMiniProgram
func NewMiniProgram(opCtx *opContext.Context, appID, refreshToken string) *MiniProgram {
	miniProgram := miniprogram.NewMiniProgram(&miniProgramConfig.Config{
		AppID:     opCtx.AppID,
		AppSecret: opCtx.AppSecret,
		Cache:     opCtx.Cache,
	})
	// 设置获取access_token的函数
	miniProgram.SetAccessTokenHandle(NewDefaultAuthrAccessToken(opCtx, appID, refreshToken))
	return &MiniProgram{
		AppID:                   appID,
		OpenPlatformMiniProgram: openPlatformMiniProgram.NewMiniProgram(opCtx, appID),
		MiniProgram:             miniProgram,
	}
}

// DefaultAuthrAccessToken 默认获取授权ak的方法
type DefaultAuthrAccessToken struct {
	opCtx        *opContext.Context
	appID        string
	refreshToken string
}

// NewDefaultAuthrAccessToken New
func NewDefaultAuthrAccessToken(opCtx *opContext.Context, appID, refreshToken string) credential.AccessTokenHandle {
	return &DefaultAuthrAccessToken{
		opCtx:        opCtx,
		appID:        appID,
		refreshToken: refreshToken,
	}
}

// GetAccessToken 获取ak
func (ak *DefaultAuthrAccessToken) GetAccessToken() (string, error) {
	token, err := ak.opCtx.GetAuthrAccessToken(ak.appID)
	logrus.Debugf("access_token=%s error=%v", token, err)
	if len(token) <= 0 || err != nil {
		authrAccessToken, err := ak.opCtx.RefreshAuthrToken(ak.appID, ak.refreshToken)
		if err != nil {
			return "", err
		}
		return authrAccessToken.AccessToken, nil
	}
	return token, err
}
