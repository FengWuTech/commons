package wxopen

import (
	"fmt"
	"github.com/FengWuTech/commons/wechat/wxmodel"
)

func GetComponentAccessToken(cmtID int) string {
	cmt := wxmodel.GetWeixinComponent(cmtID)
	if cmt == nil {
		fmt.Printf("无法查询ComponentApp\n")
		return ""
	}

	return *cmt.AccessToken
}

func SetVerifyTicket(cmtAppID string, verifyTicket string) {
	wxmodel.UpdateWeixinComponentByAppID(cmtAppID, wxmodel.WeixinComponent{
		VerifyTicket: &verifyTicket,
	})
}
