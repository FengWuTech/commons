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

	SMS_TPL_DISPATCH_ORDER_NEW_NOTICE      = 1  //您有新的任务，请访问公众号及时处理。
	SMS_TPL_LOGIN_CODE                     = 2  //验证码${code}，您正在登录，若非本人操作，请勿泄露。
	SMS_TPL_DISPATCH_ORDER_REMIND          = 3  //催单提醒：您的工单${no}处理时间过长，请登录公众号及时处理。
	SMS_TPL_PAY_BILL_NEW                   = 4  //您有新的物业费需要及时缴纳，请登录公众号${name}处理，您也可以前往物业服务中心现场缴费。
	SMS_TPL_BIRTHDAY_WISHES                = 16 //亲爱的业主，生日快乐！愿您一生被爱，永远快乐！愿你心想事成，平安顺遂。
	SMS_TPL_WEIXIN_RECON                   = 17 //对账失败，对账内容：微信账单VS本地账单
	SMS_TPL_USER_CHECKIN_AUDIT_PASS        = 18 //您的迁入的申请已经审核通过，您可以登录公众号查询详情。
	SMS_TPL_USER_CHECKIN_AUDIT_FAIL        = 19 //您的迁入的申请被驳回，您可以登录公众号重新申请。
	SMS_TPL_INNER_NOTICE                   = 22 //您有新的内部通知，请登录员工公众号查看
	SMS_TPL_RESET_PASSWORD_CODE            = 23 //验证码${code}，您正在尝试修改登录密码，请妥善保管账户信息。
	SMS_TPL_DISPATCH_ORDER_REMIND_LEADER   = 46 //您的员工{1}存在超时工单{2}，请关注处理。
	SMS_TPL_DISPATCH_ORDER_VERIFY_FAIL     = 54 //您的工单{1}核验不合格，请重新处理。
	SMS_TPL_NEW_BILL_NOTICE                = 55 //您有新的物业账单需要缴费，您可以前往公众号{1}缴纳。
	SMS_TPL_DISPATCH_ORDER_CANCEL          = 56 //工单{1}已经取消，请关注。
	SMS_TPL_EXPORT_CONFIRM                 = 57 //您正在使用高级导出功能，验证码{1}
	SMS_TPL_USER_CHECKIN_NOTICE            = 59 //主人，您已迁入{1}，可以登录公众号{2}在线获取物业服务。登录用户名:{3} 密码:{4}
	SMS_TPL_PASSWORD_RESET_NOTICE          = 60 //您的密码已被重置，新密码为:{1}
	SMS_TPL_COMPANY_CREATE_PASSWORD_NOTICE = 61 //您被设置为锋物物管平台管理员，账号:{1} 密码:{2}
	SMS_TPL_VISITOR_ARRIVED                = 63 //您邀请的访客{1}已经到访，请您关注。
	SMS_TPL_STAFF_ADD_NOTICE               = 64 //您的员工账号为用户名:{1} 密码:{2}，请及时修改密码，避免密码泄露。
)
