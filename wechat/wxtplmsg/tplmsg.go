package wxtplmsg

import (
	"encoding/json"
	"fmt"
	"github.com/FengWuTech/commons/wechat"
	"github.com/FengWuTech/commons/wechat/wxmodel"
	"github.com/chanxuehong/wechat/mp/message/template"
)

func GetTplMsgIDOrCreate(appID int, shortMsgID string) string {
	clt := wechat.InitComponentAppClient(appID)
	tplMsg := wxmodel.IsTplMsgExist(appID, shortMsgID)
	if tplMsg != nil {
		return *tplMsg.TemplateId
	}

	templateID, _ := template.AddPrivateTemplate(clt, shortMsgID)
	if templateID == "" {
		return ""
	}
	id := wxmodel.AddWeixinAppTplMsg(wxmodel.WeixinAppTplMsg{
		AppId:      &appID,
		ShortId:    &shortMsgID,
		TemplateId: &templateID,
	})
	if id > 0 {
		return templateID
	}

	return ""
}

func SendTplMsgToUser(companyID int, msg interface{}) {
	app := wxmodel.GetCompanyUserWeixinApp(companyID)
	if app == nil {
		fmt.Printf("没有业主端，请授权后发送消息\n")
		return
	}
	clt := wechat.InitComponentAppClient(*app.Id)
	if clt == nil {
		return
	}
	_, err := template.Send(clt, msg)

	str, _ := json.Marshal(msg)
	if err != nil {
		fmt.Printf("用户消息发送失败 %v", string(str))
	} else {
		str, _ := json.Marshal(msg)
		fmt.Printf("用户消息发送成功 %v", string(str))
	}
}

func SendTplMsgToStaff(companyID int, msg interface{}) {
	app := wxmodel.GetCompanyStaffWeixinApp(companyID)
	if app == nil {
		fmt.Printf("没有员工端，请授权后发送消息\n")
		return
	}
	clt := wechat.InitComponentAppClient(*app.Id)
	if clt == nil {
		return
	}

	_, err := template.Send(clt, msg)

	str, _ := json.Marshal(msg)
	if err != nil {
		fmt.Printf("员工消息发送失败 %v", string(str))
	} else {
		str, _ := json.Marshal(msg)
		fmt.Printf("员工消息发送成功 %v", string(str))
	}
}
