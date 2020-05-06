package constant

import "time"

const (
	SYSTEM_NOTICE_WSS_CHANNEL = "PMS_CHANNEL_NOTICE"

	SYSTEM_NOTICE_READ_STATUS_SENDED = 1 //已发送
	SYSTEM_NOTICE_READ_STATUS_READED = 2 //已阅读

	SYSTEM_NOTICE_TYPE_TEXT                 = 1 //文本通知
	SYSTEM_NOTICE_TYPE_NEW_THING_APPEAL     = 2 //新报事客诉通知
	SYSTEM_NOTICE_TYPE_NEW_REPAIR_APPEAL    = 3 //新报修客诉通知
	SYSTEM_NOTICE_TYPE_NEW_COMPLAINT_APPEAL = 4 //新投诉客诉通知
	SYSTEM_NOTICE_TYPE_NEW_PRAISE_APPEAL    = 5 //新表扬客诉通知
	SYSTEM_NOTICE_TYPE_NEW_SERVICE_BOOKING  = 6 //新服务预订通知
	SYSTEM_NOTICE_TYPE_NEW_USER_AUDIT       = 7 //新用户审核通知
)

type SystemNotice struct {
	CompanyID int         `json:"company_id"`
	Type      int         `json:"type"`
	Content   interface{} `json:"content"`
}

type SystemNoticeNewAppeal struct {
	AppealID int    `json:"appeal_id"`
	Detail   string `json:"detail"`
}

type SystemNoticeNewThingAppeal SystemNoticeNewAppeal
type SystemNoticeNewRepairAppeal SystemNoticeNewAppeal
type SystemNoticeNewComplaintAppeal SystemNoticeNewAppeal
type SystemNoticeNewPraiseAppeal SystemNoticeNewAppeal

type SystemNoticeNewServiceBook struct {
	OrderID  int        `json:"order_id"`
	ItemID   int        `json:"item_id"`
	ItemName string     `json:"item_name"`
	BookTime *time.Time `json:"book_time"`
}

type SystemNoticeNewUserAuditApply struct {
	UserID       int    `json:"user_id"`
	UserName     string `json:"user_name"`
	RoomLocation string `json:"room_location"`
}
