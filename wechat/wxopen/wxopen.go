package wxopen

import (
	"fmt"
	"github.com/FengWuTech/commons/wechat/wxmodel"
	"github.com/jmniu/weixin-third-dev/wxopenapi"
	"time"
)

func RefreshComponentAccessToken(cmtID int) string {
	nowTime := time.Now()
	cmt := wxmodel.GetWeixinComponent(cmtID)
	if cmt == nil {
		fmt.Printf("component accesstoken生成失败")
		return ""
	}

	if cmt.AccessTokenUpdateTime != nil && nowTime.Unix()-cmt.AccessTokenUpdateTime.Unix() < 3600 {
		return *cmt.AccessToken
	}

	wxOpen := wxopenapi.GWxOpen
	wxOpen.Init(*cmt.Token, *cmt.AesKey, *cmt.Appid, *cmt.AppSecret)
	wxOpen.SetInfo(wxopenapi.COMPONENT_VERIFY_TICKET, *cmt.VerifyTicket, time.Now().Unix(), 600)

	wxOpen.UpdateAccessToken()
	accessToken := wxOpen.GetInfo(wxopenapi.COMPONENT_ACCESS_TOKEN).Info

	if accessToken == "" {
		fmt.Printf("component accesstoken生成失败")
		return ""
	}

	wxmodel.UpdateWeixinComponent(cmtID, wxmodel.WeixinComponent{
		AccessToken:           &accessToken,
		AccessTokenUpdateTime: &nowTime,
	})

	return accessToken
}

func WxOpenInit(cmtID int) *wxopenapi.WxOpen {
	cmt := wxmodel.GetWeixinComponent(cmtID)
	if cmt == nil {
		fmt.Printf("未找到ComponentApp\n")
		return nil
	}
	wxOpen := wxopenapi.GWxOpen
	wxOpen.Init(*cmt.Token, *cmt.AesKey, *cmt.Appid, *cmt.AppSecret)
	wxOpen.SetInfo(wxopenapi.COMPONENT_VERIFY_TICKET, *cmt.VerifyTicket, time.Now().Unix(), 600)

	accessToken := RefreshComponentAccessToken(cmtID)
	wxOpen.SetInfo(wxopenapi.COMPONENT_ACCESS_TOKEN, accessToken, time.Now().Unix(), 3600)

	return wxOpen
}
