package wxmodel

import (
	"github.com/FengWuTech/commons/wechat/wxdb"
	"time"
)

type WeixinComponent struct {
	Id                    *int       `gorm:"id;default:null" json:"id,omitempty"`                                             // ID
	Type                  *int       `gorm:"type;default:null" json:"type,omitempty"`                                         // 类型 1社区端 2员工端
	Appid                 *string    `gorm:"appid;default:''" json:"appid,omitempty"`                                         // 三方平台appid
	VerifyTicket          *string    `gorm:"verify_ticket;default:''" json:"verify_ticket,omitempty"`                         // 票据
	TicketUpdateTime      *time.Time `gorm:"ticket_update_time;default:null" json:"ticket_update_time,omitempty"`             // 票据更新时间
	AppSecret             *string    `gorm:"app_secret;default:''" json:"app_secret,omitempty"`                               // 秘钥
	AccessToken           *string    `gorm:"access_token;default:''" json:"access_token,omitempty"`                           // 授权token
	AccessTokenUpdateTime *time.Time `gorm:"access_token_update_time;default:null" json:"access_token_update_time,omitempty"` // 更新时间
	Token                 *string    `gorm:"token;default:''" json:"token,omitempty"`                                         // 校验token
	AesKey                *string    `gorm:"aes_key;default:''" json:"aes_key,omitempty"`                                     // 加密串
	CreateTime            *time.Time `gorm:"create_time;default:CURRENT_TIMESTAMP" json:"create_time,omitempty"`              // 创建时间
	UpdateTime            *time.Time `gorm:"update_time;default:CURRENT_TIMESTAMP" json:"update_time,omitempty"`              // 更新时间
}

func GetWeixinComponent(id int) *WeixinComponent {
	var cpt WeixinComponent
	wxdb.DB().Where("id = ?", id).First(&cpt)
	if cpt.Id == nil {
		return nil
	} else {
		return &cpt
	}
}

func UpdateWeixinComponent(id int, upd WeixinComponent) {
	wxdb.DB().Model(&WeixinComponent{}).Where("id = ?", id).Update(upd)
}

func GetWeixinComponentByType(typ int) *WeixinComponent {
	var cpt WeixinComponent
	wxdb.DB().Where("type = ?", typ).First(&cpt)
	if cpt.Id == nil {
		return nil
	} else {
		return &cpt
	}
}

func GetWeixinComponentByAppID(appID string) *WeixinComponent {
	var cpt WeixinComponent
	wxdb.DB().Where("appid = ?", appID).First(&cpt)
	if cpt.Id == nil {
		return nil
	} else {
		return &cpt
	}
}

func UpdateWeixinComponentByAppID(appID string, upd WeixinComponent) {
	wxdb.DB().Model(&WeixinComponent{}).Where("appid = ?", appID).Update(&upd)
}

func UpdateWeixinComponentByType(typ int, upd WeixinComponent) {
	wxdb.DB().Model(&WeixinComponent{}).Where("type = ?", typ).Update(&upd)
}
