package gowechat

import (
	opContext "github.com/silenceper/wechat/v2/openplatform/context"
	openPlatformOfficialAccount "github.com/silenceper/wechat/v2/openplatform/officialaccount"
)

type OfficialAccount struct {
	AppID                       string
	OpenPlatformOfficialAccount *openPlatformOfficialAccount.OfficialAccount
}

func NewOfficialAccount(opCtx *opContext.Context, appID string) *OfficialAccount {
	return &OfficialAccount{
		AppID:                       appID,
		OpenPlatformOfficialAccount: openPlatformOfficialAccount.NewOfficialAccount(opCtx, appID),
	}
}
