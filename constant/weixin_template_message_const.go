package constant

const (
	//物业电子投票通知
	//{{first.DATA}}
	//投票议题：{{keyword1.DATA}}
	//投票内容：{{keyword2.DATA}}
	//投票范围：{{keyword3.DATA}}
	//开始时间：{{keyword4.DATA}}
	//结束时间：{{keyword5.DATA}}
	//{{remark.DATA}}
	TEMPLATE_MESSAGE_SURVEY_NOTICE = "OPENTM400166107"

	//投票结果通知
	//{{first.DATA}}
	//投票姓名：{{keyword1.DATA}}
	//活动票数：{{keyword2.DATA}}
	//投票排名：{{keyword3.DATA}}
	//{{remark.DATA}}
	TEMPLATE_MESSAGE_SURVEY_RESULT_NOTICE = "OPENTM406039008"

	//物业管理通知
	//{{first.DATA}}
	//标题：{{keyword1.DATA}}
	//发布时间：{{keyword2.DATA}}
	//内容：{{keyword3.DATA}}
	//{{remark.DATA}}
	TEMPLATE_MESSAGE_PMS_MANAGE_NOTICE = "OPENTM204594069"

	//物业费缴纳通知
	//{{first.DATA}}
	//社区名称：{{keyword1.DATA}}
	//楼宇户号：{{keyword2.DATA}}
	//应缴费用：{{keyword3.DATA}}
	//截止日期：{{keyword4.DATA}}
	//{{remark.DATA}}
	TEMPLATE_MESSAGE_PAYMENT_NOTICE = "OPENTM411067136"

	//审核通知
	//{{first.DATA}}
	//审核人：{{keyword1.DATA}}
	//审核内容：{{keyword2.DATA}}
	//审核日期：{{keyword3.DATA}}
	//{{remark.DATA}}
	TEMPLATE_MESSAGE_AUDIT_NOTICE = "OPENTM417949050"

	//受理通知
	//{{first.DATA}}
	//申请时间：{{keyword1.DATA}}
	//服务项目：{{keyword2.DATA}}
	//地址：{{keyword3.DATA}}
	//{{remark.DATA}}
	TEMPLATE_MESSAGE_ACCEPTANCE = "OPENTM407308160"

	//物业管理通知
	//{{first.DATA}}
	//标题：{{keyword1.DATA}}
	//发布时间：{{keyword2.DATA}}
	//内容：{{keyword3.DATA}}
	//{{remark.DATA}}
	TEMPLATE_MESSAGE_MANAGE_NOTICE = "OPENTM204594069"

	//工单超时通知
	//{{first.DATA}}
	//工单号：{{keyword1.DATA}}
	//工单类别：{{keyword2.DATA}}
	//工单状态：{{keyword3.DATA}}
	//处理人员：{{keyword4.DATA}}
	//具体内容：{{keyword5.DATA}}
	//{{remark.DATA}}
	TEMPLATE_MESSAGE_ORDER_TIMEOUT = "OPENTM401096729"

	//工单处理提醒
	//{{first.DATA}}
	//工单关联位置：{{keyword1.DATA}}
	//工单编号：{{keyword2.DATA}}
	//工单内容：{{keyword3.DATA}}
	//操作人：{{keyword4.DATA}}
	//时间：{{keyword5.DATA}}
	//{{remark.DATA}}
	TEMPLATE_MESSAGE_ORDER_REMIND = "OPENTM415863052"

	//报修处理结果
	//{{first.DATA}}
	//工单编号：{{keyword1.DATA}}
	//房屋编号：{{keyword2.DATA}}
	//维修主题：{{keyword3.DATA}}
	//维修工程师：{{keyword4.DATA}}
	//维修完成时间：{{keyword5.DATA}}
	//{{remark.DATA}}
	TEMPLATE_MESSAGE_REPAIR_RESULT = "OPENTM202578956"

	//新工单提醒
	//{{first.DATA}}
	//工单内容：{{keyword1.DATA}}
	//工单地址：{{keyword2.DATA}}
	//派单时间：{{keyword3.DATA}}
	//派单人员：{{keyword4.DATA}}
	//{{remark.DATA}}
	TEMPLATE_MESSAGE_ORDER_NEW = "OPENTM417876520"
)
