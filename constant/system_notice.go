package constant

const (
	SYSTEM_NOTICE_TYPE_TEXT                 = 1 //文本通知
	SYSTEM_NOTICE_TYPE_NEW_THING_APPEAL     = 2 //新报事客诉通知
	SYSTEM_NOTICE_TYPE_NEW_REPAIR_APPEAL    = 3 //新报修客诉通知
	SYSTEM_NOTICE_TYPE_NEW_COMPLAINT_APPEAL = 4 //新投诉客诉通知
	SYSTEM_NOTICE_TYPE_NEW_PRAISE_APPEAL    = 5 //新表扬客诉通知
	SYSTEM_NOTICE_TYPE_NEW_SERVICE_BOOKING  = 6 //新服务预订通知
	SYSTEM_NOTICE_TYPE_NEW_USER_AUDIT       = 7 //新用户审核通知
)

type SystemNoticeNewAppeal struct {
	AppealID int `json:"appeal_id"`
}

type SystemNoticeNewServiceBooking struct {
	BookingOrderID int `json:"booking_order_id"`
}

type SystemNoticeNewUserAudit struct {
	AuditID int `json:"audit_id"`
}
