# gowechat

微信开放平台 Go SDK

```shell
go get -u github.com/freyo/gowechat
```

```go
package main

import (
	"fmt"
	"github.com/freyo/gowechat"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/openplatform/config"
)

func main() {
	openPlatform := wechat.NewWechat().GetOpenPlatform(&config.Config{
		AppID:          "openplatform-appid",
		AppSecret:      "openplatform-secret",
		Token:          "openplatform-token",
		EncodingAESKey: "openplatform-aeskey",
		Cache: cache.NewRedis(&cache.RedisOpts{
			Host:     "redis-host",
			Password: "redis-pwd",
		}),
	})

	mp := gowechat.NewMiniProgram(openPlatform.Context, "authr-appid", "authr-refresh-token")
	resp, err := mp.CheckNickName(&gowechat.MiniProgramCheckNickNameRequest{
		Nickname: "nick-name",
	})
	fmt.Println(resp, err)

	oc := gowechat.NewOfficialAccount(openPlatform.Context, "authr-appid")
	path := oc.RegisterMiniProgramPage(1, "https://")
	fmt.Println(path)
}
```