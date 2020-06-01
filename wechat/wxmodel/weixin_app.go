package wxmodel

import (
	"github.com/FengWuTech/commons/wechat/wxcons"
	"github.com/FengWuTech/commons/wechat/wxdb"
	"time"
)

type WeixinApp struct {
	Id                              *int       `gorm:"id;default:null" json:"id,omitempty"`
	ComponentId                     *int       `gorm:"component_id;default:0" json:"component_id,omitempty"`                                                  // 三方平台ID
	CompanyId                       *int       `gorm:"company_id;default:0" json:"company_id,omitempty"`                                                      // 公司ID
	AppType                         *int       `gorm:"app_type;default:0" json:"app_type,omitempty"`                                                          // 1公众号 2小程序
	AuthorizerAppid                 *string    `gorm:"authorizer_appid;default:''" json:"authorizer_appid,omitempty"`                                         // 授权方APPID
	AuthorizerAppName               *string    `gorm:"authorizer_app_name;default:''" json:"authorizer_app_name,omitempty"`                                   // 授权方APP名称
	AuthorizerAccessToken           *string    `gorm:"authorizer_access_token;default:''" json:"authorizer_access_token,omitempty"`                           // 授权方accesstoken
	AuthorizerAccessTokenExpiresIn  *int       `gorm:"authorizer_access_token_expires_in;default:0" json:"authorizer_access_token_expires_in,omitempty"`      // 授权方accessToken有效周期
	AuthorizerAccessTokenUpdateTime *time.Time `gorm:"authorizer_access_token_update_time;default:null" json:"authorizer_access_token_update_time,omitempty"` // 授权方accessToken更新时间
	AuthorizerRefreshToken          *string    `gorm:"authorizer_refresh_token;default:''" json:"authorizer_refresh_token,omitempty"`                         // 刷新token
	AuthorizerFuncInfo              *string    `gorm:"authorizer_func_info;default:''" json:"authorizer_func_info,omitempty"`                                 // func_info
	NickName                        *string    `gorm:"nick_name;default:''" json:"nick_name,omitempty"`                                                       // 昵称
	MchId                           *string    `gorm:"mch_id;default:''" json:"mch_id,omitempty"`                                                             // 对应的商户ID
	QrcodeUrl                       *string    `gorm:"qrcode_url;default:''" json:"qrcode_url,omitempty"`                                                     // 二维码URL
	CreateTime                      *time.Time `gorm:"create_time;default:CURRENT_TIMESTAMP" json:"create_time,omitempty"`                                    // 创建时间
	UpdateTime                      *time.Time `gorm:"update_time;default:CURRENT_TIMESTAMP" json:"update_time,omitempty"`                                    // 更新时间
}

func GetCompanyWeixinApp(companyID int) []WeixinApp {
	var list []WeixinApp
	wxdb.DB().Where("company_id = ?", companyID).Find(&list)
	return list
}

func GetCompanyUserWeixinApp(companyID int) *WeixinApp {
	list := GetCompanyWeixinApp(companyID)
	for _, app := range list {
		cmt := GetWeixinComponent(*app.ComponentId)
		if *cmt.Type == wxcons.COMPONENT_TYPE_USER {
			return &app
		}
	}
	return nil
}

func GetCompanyStaffWeixinApp(companyID int) *WeixinApp {
	list := GetCompanyWeixinApp(companyID)
	for _, app := range list {
		cmt := GetWeixinComponent(*app.ComponentId)
		if *cmt.Type == wxcons.COMPONENT_TYPE_STAFF {
			return &app
		}
	}
	return nil
}

func IsWeixinAppExist(companyID int, appID string, cmtID int) *WeixinApp {
	var app WeixinApp
	wxdb.DB().Where("company_id = ? and component_id = ? and authorizer_appid = ?", companyID, cmtID, appID).First(&app)
	if app.Id == nil {
		return nil
	} else {
		return &app
	}
}

func GetWeixinApp(id int) *WeixinApp {
	var app WeixinApp
	wxdb.DB().Where("id = ?", &id).First(&app)
	if app.Id == nil {
		return nil
	} else {
		return &app
	}
}

func UpdateWeixinApp(id int, upd WeixinApp) {
	wxdb.DB().Model(&WeixinApp{}).Where("id = ?", id).Update(upd)
}

func AddWeixinApp(app WeixinApp) int {
	wxdb.DB().Create(&app)
	return *app.Id
}

func DeleteWeixinApp(appID string, authorizerAppid string) {
	cmt := GetWeixinComponentByAppID(appID)
	wxdb.DB().Delete(&WeixinApp{}, "component_id = ? and authorizer_appid = ?", *cmt.Id, authorizerAppid)
}
