package wxmodel

import (
	"github.com/FengWuTech/commons/wechat/wxdb"
	"time"
)

type WeixinAppTplMsg struct {
	Id         *int       `gorm:"id;default:null" json:"id,omitempty"`                                // ID
	AppId      *int       `gorm:"app_id;default:0" json:"app_id,omitempty"`                           // weixin_app表的ID
	ShortId    *string    `gorm:"short_id;default:''" json:"short_id,omitempty"`                      // 模板消息短ID
	TemplateId *string    `gorm:"template_id;default:''" json:"template_id,omitempty"`                // 模板消息实例后的ID
	CreateTime *time.Time `gorm:"create_time;default:CURRENT_TIMESTAMP" json:"create_time,omitempty"` // 创建时间
	UpdateTime *time.Time `gorm:"update_time;default:CURRENT_TIMESTAMP" json:"update_time,omitempty"` // 更新时间
}

func IsTplMsgExist(appID int, shortID string) *WeixinAppTplMsg {
	var tpl WeixinAppTplMsg
	wxdb.DB().Where("app_id = ? and short_id = ?", appID, shortID).First(&tpl)
	if tpl.Id == nil {
		return nil
	} else {
		return &tpl
	}
}

func AddWeixinAppTplMsg(tpl WeixinAppTplMsg) int {
	wxdb.DB().Create(&tpl)
	return *tpl.Id
}
