package wxopen

import (
	"github.com/FengWuTech/commons/wechat/wxmodel"
	"time"
)

func GetWeixinAppAccessToken(appID int) string {
	app := wxmodel.GetWeixinApp(appID)
	if app == nil {
		return ""
	}

	cmt := wxmodel.GetWeixinComponent(*app.ComponentId)

	var updTime int64
	if app.AuthorizerAccessTokenUpdateTime == nil {
		updTime = 0
	} else {
		updTime = app.AuthorizerAccessTokenUpdateTime.Unix()
	}
	nowTime := time.Now().Unix()
	if nowTime-updTime < int64(*app.AuthorizerAccessTokenExpiresIn) { //过期
		return *app.AuthorizerAccessToken
	} else {
		wxOpen := WxOpenInit(*cmt.Type)
		res := wxOpen.UpdateAuthAccessToken(*app.AuthorizerAppid, *app.AuthorizerRefreshToken)
		if len(res.AuthorizerAccessToken) > 0 { //可以正常获取，更新到数据库
			var upd wxmodel.WeixinApp
			upd.AuthorizerAccessToken = &res.AuthorizerAccessToken
			upd.AuthorizerRefreshToken = &res.AuthorizerRefreshToken
			upd.AuthorizerAccessTokenExpiresIn = &res.ExpiresIn
			generateTime := time.Now()
			upd.AuthorizerAccessTokenUpdateTime = &generateTime
			wxmodel.UpdateWeixinApp(appID, upd)

			return res.AuthorizerAccessToken
		} else {
			return ""
		}
	}
}
