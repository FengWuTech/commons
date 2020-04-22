package constant

const (
	SMS_TPL_TYPE_VERIFY_CODE = 0 //验证码
	SMS_TPL_TYPE_NORMAL      = 1 //普通短信
	SMS_TPL_TYPE_AD          = 2 //广告

	SMS_TPL_REGION_DOMESTIC      = 0 //国内短信
	SMS_TPL_REGION_INTERNATIONAL = 1 //港澳台和国际短信

	SMS_TPL_AUDITING   = 0 //审核中
	SMS_TPL_AUDIT_PASS = 1 //审核通过
	SMS_TPL_AUDIT_FAIL = 2 //审核失败
)
