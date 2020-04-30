package wss

import "time"

const (
	CHANNEL_NOTICE = "PMS_CHANNEL_NOTICE"

	MESSAGE_TYPE_WELCOME        = 1
	MESSAGE_TYPE_NEW_APPEAL     = 2
	MESSAGE_TYPE_NEW_BOOK_ORDER = 3
	MESSAGE_TYPE_NEW_USER_AUDIT = 4
)

type Message struct {
	CompanyID int         `json:"company_id"`
	Type      int         `json:"type"`
	Content   interface{} `json:"content"`
}

type MessageWelcome struct {
	Welcome string `json:"welcome"`
}

type MessageNewAppealOrder struct {
	AppealID int    `json:"appeal_id"`
	Detail   string `json:"detail"`
}

type MessageNewBook struct {
	OrderID  int        `json:"order_id"`
	ItemID   int        `json:"item_id"`
	ItemName string     `json:"item_name"`
	BookTime *time.Time `json:"book_time"`
}

type MessageNewUserAudit struct {
	UserID       int    `json:"user_id"`
	UserName     string `json:"user_name"`
	RoomLocation string `json:"room_location"`
}
