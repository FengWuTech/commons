package wechat

import (
	"fmt"
	"github.com/FengWuTech/commons/wechat/wxmodel"
	"github.com/FengWuTech/commons/wechat/wxopen"
	"github.com/chanxuehong/wechat/mp/core"
)

type ThirdDevAccessTokenServer struct {
	ComponentID int
	AppID       int
}

func NewThirdDevAccessTokenServer(cmtID int, appID int) *ThirdDevAccessTokenServer {
	return &ThirdDevAccessTokenServer{
		ComponentID: cmtID,
		AppID:       appID,
	}
}

func (srv *ThirdDevAccessTokenServer) Token() (token string, err error) {
	return srv.GetAccessToken(), nil
}

func (srv *ThirdDevAccessTokenServer) RefreshToken(currentToken string) (token string, err error) {
	return srv.GetAccessToken(), nil
}

func (srv *ThirdDevAccessTokenServer) IID01332E16DF5011E5A9D5A4DB30FED8E1() {

}

func (srv *ThirdDevAccessTokenServer) GetAccessToken() string {
	if srv.AppID == 0 {
		return wxopen.GetComponentAccessToken(srv.ComponentID)
	}

	return wxopen.GetWeixinAppAccessToken(srv.AppID)
}

func InitComponentClient(cmtType int) *core.Client {
	wxopen.WxOpenInit(cmtType)
	var srv = NewThirdDevAccessTokenServer(cmtType, 0)
	client := core.NewClient(srv, nil)
	return client
}

func InitComponentAppClient(appID int) *core.Client {
	app := wxmodel.GetWeixinApp(appID)
	if app == nil {
		fmt.Printf("App为空\n")
		return nil
	}
	cmt := wxmodel.GetWeixinComponent(*app.ComponentId)
	if cmt == nil {
		fmt.Printf("Component为空\n")
		return nil
	}
	wxopen.WxOpenInit(*cmt.Type)
	var srv = NewThirdDevAccessTokenServer(*cmt.Id, appID)
	client := core.NewClient(srv, nil)
	return client
}
