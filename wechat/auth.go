package wechat

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/FengWuTech/commons/logger"
	"github.com/FengWuTech/commons/wechat/wxmodel"
	"github.com/FengWuTech/commons/wechat/wxopen"
	"github.com/FengWuTech/commons/wechat/wxutil"
	"github.com/chanxuehong/wechat/mp/menu"
	"github.com/jmniu/weixin-third-dev/wxopenapi"
	"time"
)

func PreAuthCode(cmtID int) string {
	cmt := wxmodel.GetWeixinComponent(cmtID)
	if cmt == nil {
		fmt.Printf("无法查找服务商 %d \n", cmtID)
		return ""
	}
	wxOpen := wxopen.WxOpenInit(*cmt.Id)
	wxOpen.UpdatePreAuthCode()
	return wxOpen.GetInfo(wxopenapi.PRE_AUTH_CODE).Info
}

func AuthFinish(companyID int, cmtType int, authCode string) int {
	cmt := wxmodel.GetWeixinComponentByType(cmtType)
	wxOpen := wxopen.WxOpenInit(*cmt.Id)
	authAccessToken := wxOpen.GetAuthAccessToken(authCode)

	//保存授权信息
	var appType = 1
	var appID = authAccessToken.AuthorizerAppid
	var authInfo = wxOpen.GetAuthInfo(appID)

	authInfoStr, _ := json.Marshal(authInfo)
	logger.Infof("authInfo: %v", authInfoStr)

	qrcodeURL := wxutil.UploadFromURL(authInfo.AuthorizerInfo.QrcodeUrl)
	existApp := wxmodel.IsWeixinAppExist(companyID, appID, *cmt.Id)
	if existApp != nil {
		wxmodel.UpdateWeixinApp(*existApp.Id, wxmodel.WeixinApp{
			AuthorizerAppName:      &authInfo.AuthorizerInfo.NickName,
			QrcodeUrl:              &qrcodeURL,
			AuthorizerRefreshToken: &authInfo.AuthorizationInfo.AuthorizerRefreshToken,
		})
		return *existApp.Id
	} else {
		var nowTime = time.Now()
		var accessToken = authAccessToken.AuthorizerAccessToken
		var expiresIn = authAccessToken.ExpiresIn
		var refreshToken = authAccessToken.AuthorizerRefreshToken
		var funcInfo, _ = json.Marshal(authAccessToken.FuncInfos)
		var strFuncInfo = string(funcInfo)
		newAppID := wxmodel.AddWeixinApp(wxmodel.WeixinApp{
			CompanyId:                       &companyID,
			ComponentId:                     cmt.Id,
			AppType:                         &appType,
			AuthorizerAppid:                 &appID,
			AuthorizerAppName:               &authInfo.AuthorizerInfo.NickName,
			AuthorizerAccessToken:           &accessToken,
			AuthorizerAccessTokenExpiresIn:  &expiresIn,
			AuthorizerRefreshToken:          &refreshToken,
			AuthorizerFuncInfo:              &strFuncInfo,
			AuthorizerAccessTokenUpdateTime: &nowTime,
			QrcodeUrl:                       &qrcodeURL,
		})
		return newAppID
	}
}

func AuthEventHandle(cmtType int, rawXmlData string, msgSignature string, timestamp string, nonce string) {
	cmt := wxmodel.GetWeixinComponentByType(cmtType)
	if cmt == nil {
		logger.Infof("Component为空")
		return
	}
	wxopen.WxOpenInit(*cmt.Id)
	//解密xml
	_, decodedXml := wxopenapi.GWxOpen.Decrypt(msgSignature, timestamp, nonce, rawXmlData)
	logger.Infof("收到微信调用：%s", decodedXml)
	if decodedXml == "" {
		return
	}

	//解析解密后的数据
	var notifyAuth wxopenapi.NotifyAuth
	_ = xml.Unmarshal([]byte(decodedXml), &notifyAuth)

	//处理不同事件请求
	switch notifyAuth.InfoType {
	case "component_verify_ticket":
		app := wxmodel.GetWeixinComponentByAppID(notifyAuth.AppId)
		if app == nil {
			fmt.Printf("未找到Component\n")
			return
		}
		nowTime := time.Now()
		wxmodel.UpdateWeixinComponent(*app.Id, wxmodel.WeixinComponent{
			VerifyTicket:     &notifyAuth.ComponentVerifyTicket,
			TicketUpdateTime: &nowTime,
		})
	case "authorized":
		logger.Warnf("authorized %v", notifyAuth)
	case "unauthorized":
		//删除公众号菜单
		app := wxmodel.GetWeixinAppByAuthorizerAppID(notifyAuth.AppId, notifyAuth.AuthorizerAppid)
		clt := InitComponentAppClient(*app.Id)
		err := menu.Delete(clt)
		if err != nil {
			logger.Warnf("删除公众号菜单失败 %v", err)
		}

		//删除本机记录
		wxmodel.DeleteWeixinApp(notifyAuth.AppId, notifyAuth.AuthorizerAppid)
	case "updateauthorized":
	default:
		logger.Warnf("未找到消息处理函数 infoType[%s] notifyAuth[%v]", notifyAuth.InfoType, notifyAuth)
	}
}

func GetOpenIDByAuthCode(appID int, code string) string {
	app := wxmodel.GetWeixinApp(appID)
	if app == nil {
		return ""
	}
	wxOpen := wxopen.WxOpenInit(*app.ComponentId)

	cmt := wxmodel.GetWeixinComponent(*app.ComponentId)

	authInfo := wxOpen.GetUserAuthInfo(*app.AuthorizerAppid, code, *cmt.Appid)
	if authInfo == nil {
		return ""
	}

	return authInfo.OpenID
}
