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

	SMS_TPL_DISPATCH_ORDER_NEW_NOTICE    = 1  //物业维修工单提醒
	SMS_TPL_LOGIN_CODE                   = 2  //登录确认验证码
	SMS_TPL_DISPATCH_ORDER_REMIND        = 3  //工单催单提醒
	SMS_TPL_PAY_BILL_NEW                 = 4  //物业缴费通知
	SMS_TPL_BIRTHDAY_WISHES              = 16 //生日祝福
	SMS_TPL_WEIXIN_RECON                 = 17 //微信对账
	SMS_TPL_USER_CHECKIN_AUDIT_PASS      = 18 //迁入申请审核通过
	SMS_TPL_USER_CHECKIN_AUDIT_FAIL      = 19 //迁入申请审核拒绝
	SMS_TPL_INNER_NOTICE                 = 22 //内部通知
	SMS_TPL_RESET_PASSWORD_CODE          = 23 //重置密码
	SMS_TPL_DISPATCH_ORDER_REMIND_LEADER = 46 //推送给领导的催单短信
	SMS_TPL_DISPATCH_ORDER_VERIFY_FAIL   = 54 //核验不成功
)
