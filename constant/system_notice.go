package constant

const (
	SYSTEM_NOTICE_TYPE_TEXT                = 1 //文本通知
	SYSTEM_NOTICE_TYPE_NEW_APPEAL          = 2 //新客诉通知
	SYSTEM_NOTICE_TYPE_NEW_SERVICE_BOOKING = 3 //新服务预订通知
	SYSTEM_NOTICE_TYPE_NEW_USER_AUDIT      = 4 //新用户审核通知
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
