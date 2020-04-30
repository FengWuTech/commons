package wss

const (
	CHANNEL_NOTICE = "PMS_CHANNEL_NOTICE"

	MESSAGE_TYPE_WELCOME    = 1
	MESSAGE_TYPE_NEW_APPEAL = 2
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
