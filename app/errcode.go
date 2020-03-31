package app

const (
	SUCCESS        = 200
	INVALID_PARAMS = 400
	NOT_EXIST      = 401
	FORBID         = 403
	ERROR          = 500
	NOT_SUPPORT    = 501
)

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",
	NOT_EXIST:      "请求对象不存在",
	FORBID:         "操作被禁止",
	NOT_SUPPORT:    "不支持的操作",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}


