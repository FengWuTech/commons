package constant

import (
	"sort"
)

const (
	CHARGE_TYPE_SEQUENCE_PAY = 1 //顺序支付
	CHARGE_TYPE_RANDOM_PAY   = 2 //随机支付

	CHARGE_PROJECT_TYPE          = 1 //费项类别:0周期性、1临时性、2押金性
	CHARGE_PROJECT_FEE_PRECISE   = 2 //费用精确位，枚举：1元、2角、3分
	CHARGE_PROJECT_DEFAULT_POINT = 3 //默认收费周期 0当期收当期 1当期收上期 2当期收下期
	CHARGE_PROJECT_LESS_METHOD   = 4 //不足周期计算方式  0按一个周期算  1按实际天数算
	CHARGE_PROJECT_LATEFEE_POINT = 5 //违约金开始计算时间 0账单开始时间  1账单结束时间

	CHARGE_STANDARD_TOTAL_CALCULATION_METHOD = 6  //总费用计算方式
	CHARGE_STANDARD_COUNT_TYPE               = 17 //数量的计量方式【为了兼容，要一个全集】
	CHARGE_STANDARD_COUNT_TYPE_CYCLE         = 7  //数量的计量方式（周期性收费项目）
	CHARGE_STANDARD_COUNT_TYPE_TEMP_PLEDGE   = 17 //数量的计量方式（临时性/押金性收费项目）

	CHARGE_RPINT_TPL_TYPE                           = 8  //打印模板类型 1通知单模板、2收据模板
	CHARGE_RPINT_TPL_CONTENT_RECEIVE_HEADER_ELEMENT = 9  //收据-页眉显示内容
	CHARGE_RPINT_TPL_CONTENT_RECEIVE_FOOTER_ELEMENT = 10 //收据-页脚显示内容
	CHARGE_RPINT_TPL_CONTENT_RECEIVE_BODY_ELEMENT   = 11 //收据-内容区显示内容
	CHARGE_RPINT_TPL_CONTENT_NOTICE_HEADER_ELEMENT  = 12 //通知-页眉显示内容
	CHARGE_RPINT_TPL_CONTENT_NOTICE_FOOTER_ELEMENT  = 13 //通知-页脚显示内容
	CHARGE_RPINT_TPL_CONTENT_NOTICE_BODY_ELEMENT    = 14 //通知-内容区显示内容

	CHARGE_PAY_METHOD = 15 //付费方式
	CHARGE_PAY_CYCLE  = 16 //收费周期
	//17用掉了
	//---------------------------------------------------
	CHARGE_PAY_METHOD_ALL = 18 //所有收费方式，统计用

	//费项基本分类
	CHARGE_PROJECT_TYPE_BASIC_CYCLE   = 0 //周期类
	CHARGE_PROJECT_TYPE_BASIC_TMP     = 1 //临时类
	CHARGE_PROJECT_TYPE_BASIC_DEPOSIT = 2 //押金类

	//费项精确位
	CHARGE_PROJECT_FEE_PRECISE_YUAN = 1 //元
	CHARGE_PROJECT_FEE_PRECISE_JIAO = 2 //角
	CHARGE_PROJECT_FEE_PRECISE_FEN  = 3 //分

	//费项的默认收费点
	CHARGE_PROJECT_DEFAULT_POINT_CUR    = 0 //当期收当期
	CHARGE_PROJECT_DEFAULT_POINT_AFTER  = 1 //当期收上期
	CHARGE_PROJECT_DEFAULT_POINT_BEFORE = 2 //当期收下期

	//不足周期计算方式
	CHARGE_PROJECT_LESS_METHOD_CYCLE = 1 //按一个周期算
	CHARGE_PROJECT_LESS_METHOD_DAY   = 2 //按实际天数算

	//违约金开始计算时间
	CHARGE_PROJECT_LATEFEE_POINT_START = 1 //账单开始时间
	CHARGE_PROJECT_LATEFEE_POINT_END   = 2 //账单结束时间

	//总费用计算方式
	CHARGE_STANDARD_TOTAL_CALCULATION_METHOD_UC     = 1 //单价*数量
	CHARGE_STANDARD_TOTAL_CALCULATION_METHOD_SINGLE = 2 //每户单独输入
	CHARGE_STANDARD_TOTAL_CALCULATION_METHOD_FIXED  = 3 //固定金额
	CHARGE_STANDARD_TOTAL_CALCULATION_METHOD_CUSTOM = 4 //自定义公式

	//数量的计量方式
	CHARGE_STANDARD_COUNT_TYPE_BUILD_AREA   = 1 //建筑面积(平方米)
	CHARGE_STANDARD_COUNT_TYPE_USE_AREA     = 2 //套内面积(平方米)
	CHARGE_STANDARD_COUNT_TYPE_POOL_AREA    = 3 //公摊面积(平方米)
	CHARGE_STANDARD_COUNT_TYPE_AMOUNT       = 4 //用量(止度-起度)
	CHARGE_STANDARD_COUNT_TYPE_PARKING_AREA = 5 //车位面积(平方米)
	CHARGE_STANDARD_COUNT_TYPE_CUSTOM_INPUT = 6 //自定义数量

	//打印模板-类型
	CHARGE_RPINT_TPL_TYPE_NOTICE  = 1 //1通知单模板
	CHARGE_RPINT_TPL_TYPE_RECEIVE = 2 //2收据模板

	//打印模板-页眉页脚 显示内容（通用，大家来挑）
	CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_TITLE      = "title"       //1标题
	CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_CODE       = "code"        //2编号
	CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_LOGO       = "logo"        //3LOGO
	CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_RPINTDATE  = "print_date"  //4打印日期
	CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_GROUP      = "group"       //5小区
	CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_ROOM       = "room"        //6房号
	CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_NAME       = "name"        //7住户姓名
	CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_MOBILE     = "mobile"      //8住户手机号
	CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_STATISTICS = "statistics"  //9合计
	CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_RDATE      = "rdate"       //10收款日期
	CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_PAYUSER    = "pay_user"    //11付款人
	CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_RECVUSER   = "recv_user"   //12收款人
	CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_RECVREMARK = "recv_remark" //13收款备注
	CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_RECVMETHOD = "recv_method" //14收款方式
	CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_REMARK     = "remark"      //15说明

	//内容区显示内容:
	CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_CODE         = "code"          //内容区显示内容: 1房号/车位号/住户
	CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_BUILDING     = "building"      //2楼宇
	CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_UNIT         = "unit"          //3单元
	CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_RECVPROJECT  = "recv_project"  //4收费项目
	CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_RECVSTANDARD = "recv_standard" //5收费标准
	CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_STIME        = "stime"         //6开始时间
	CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_ETIME        = "etime"         //7结束时间
	CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_SMETER       = "smeter"        //8起度
	CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_EMETER       = "emeter"        //9止度
	CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_AMOUNT       = "amount"        //10数量
	CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_PRICE        = "price"         //11单价
	CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_TOTAL        = "total"         //12金额
	CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_DISCOUNT     = "discount"      //13优惠
	CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_LATEFEE      = "latefee"       //14违约金
	CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_SHOULDFEE    = "should_fee"    //15应收合计
	CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_BATCH        = "batch"         //16账单期数
	CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_REMARK       = "remark"        //17备注

	//费用类型
	CHARGE_BIND_TYPE_ROOM         = 1 //物业费
	CHARGE_BIND_TYPE_PARKING      = 2 //车位费
	CHARGE_BIND_TYPE_USER         = 3 //用户定向收费
	CHARGE_BIND_TYPE_METER_ROOM   = 4 //户表费
	CHARGE_BIND_TYPE_METER_PUBLIC = 5 //公表费
	CHARGE_BIND_TYPE_OTHER        = 6 //其他费用

	//支付渠道
	CHARGE_PAY_METHOD_OFFLINE_CASH           = 1  //线下支付-现金
	CHARGE_PAY_METHOD_OFFLINE_POS            = 2  //线下支付-POS机刷卡
	CHARGE_PAY_METHOD_OFFLINE_WEIXIN         = 3  //线下支付-微信扫码
	CHARGE_PAY_METHOD_OFFLINE_ZHIFUBAO       = 4  //线下支付-支付宝扫码
	CHARGE_PAY_METHOD_OFFLINE_BANK           = 5  //线下支付-银行卡转账
	CHARGE_PAY_METHOD_ONLINE_WEIXIN          = 6  //线上支付-微信
	CHARGE_PAY_METHOD_ONLINE_ZHIFUBAO        = 7  //线上支付-支付宝
	CHARGE_PAY_METHOD_ONLINE_BANK_COLLECTION = 8  //线上支付-银行托收
	CHARGE_PAY_METHOD_OTHER                  = 9  //其他收款方式
	CHARGE_PAY_METHOD_DEDUCTE                = 10 //预充值抵扣（不传给fe）
	CHARGE_PAY_METHOD_BACKEND                = 11 //后台支付（不传给fe）

	//支付状态
	CHARGE_PAY_STATUS_CREATED  = 1 //账单已创建
	CHARGE_PAY_STATUS_PAID     = 2 //账单已支付
	CHARGE_PAY_STATUS_REFUNDED = 3 //账单已退款

	//收费周期
	CHARGE_PAY_CYCLE_MONTH    = 1 //每月
	CHARGE_PAY_CYCLE_SEASON   = 2 //每季度
	CHARGE_PAY_CYCLE_HALFYEAR = 3 //每半年
	CHARGE_PAY_CYCLE_YEAR     = 4 //每年

	ACCOUNT_FLOW_TYPE_RECHARGE          = 1 //充值
	ACCOUNT_FLOW_TYPE_REFUND_TO_ACCOUNT = 2 //退款到用户账户
	ACCOUNT_FLOW_TYPE_DEDUCT            = 3 //划扣
	ACCOUNT_FLOW_TYPE_REFUND_TO_USER    = 4 //退款给用户
)

var ACCOUNT_FLOW_TYPE = []map[string]interface{}{
	{
		"type":  ACCOUNT_FLOW_TYPE_RECHARGE,
		"label": "充值",
	},
	{
		"type":  ACCOUNT_FLOW_TYPE_REFUND_TO_ACCOUNT,
		"label": "退款",
	},
	{
		"type":  ACCOUNT_FLOW_TYPE_DEDUCT,
		"label": "划扣",
	},
	{
		"type":  ACCOUNT_FLOW_TYPE_REFUND_TO_USER,
		"label": "退款",
	},
}

//打印模板配置
var CHARGE_RPINT_TPL_INFO = map[int][]map[string]interface{}{
	CHARGE_RPINT_TPL_CONTENT_RECEIVE_HEADER_ELEMENT: []map[string]interface{}{
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_TITLE,
			"label": "标题", //这个是允许自定义的，只是很多type不展示label,此时无意义而已.
			//value是设置死的，不需要动态变化，一旦设置了就永远是这个value
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_LOGO,
			"label": "LOGO",
			//value是设置死的，不需要动态变化，一旦设置了就永远是这个value
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_CODE,
			"label": "编号",
			"value": "【title】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_RPINTDATE,
			"label": "打印日期",
			"value": "【print_date】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_GROUP,
			"label": "小区",
			"value": "【group】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_ROOM,
			"label": "房号",
			"value": "【room】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_NAME,
			"label": "住户姓名",
			"value": "【name】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_MOBILE,
			"label": "住户手机号",
			"value": "【mobile】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_STATISTICS,
			"label": "合计",
			"value": "【staticstics】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_RDATE,
			"label": "收款日期",
			"value": "【rdate】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_PAYUSER,
			"label": "付款人",
			"value": "【pay_user】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_RECVUSER,
			"label": "收款人",
			"value": "【recv_user】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
	},
	CHARGE_RPINT_TPL_CONTENT_RECEIVE_BODY_ELEMENT: []map[string]interface{}{
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_CODE,
			"label": "房号/车位号/住户",
			"value": "【code】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_BUILDING,
			"label": "楼宇",
			"value": "【building】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_UNIT,
			"label": "单元",
			"value": "【unit】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_RECVPROJECT,
			"label": "收费项目",
			"value": "【charge_project】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_RECVSTANDARD,
			"label": "收费标准",
			"value": "【charge_standard】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_STIME,
			"label": "开始时间",
			"value": "【stime】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_ETIME,
			"label": "结束时间",
			"value": "【etime】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_SMETER,
			"label": "起度",
			"value": "【smeter】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_EMETER,
			"label": "止度",
			"value": "【emeter】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_AMOUNT,
			"label": "数量",
			"value": "【amount】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_PRICE,
			"label": "单价",
			"value": "【price】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_TOTAL,
			"label": "金额",
			"value": "【total】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_DISCOUNT,
			"label": "优惠",
			"value": "【discount】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_LATEFEE,
			"label": "违约金",
			"value": "【latefee】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_SHOULDFEE,
			"label": "应收合计",
			"value": "【should_fee】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_BATCH,
			"label": "账单期数",
			"value": "【batch】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_BODY_ELEMENT_REMARK,
			"label": "备注",
			"value": "【remark】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
	},
	CHARGE_RPINT_TPL_CONTENT_RECEIVE_FOOTER_ELEMENT: []map[string]interface{}{
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_RPINTDATE,
			"label": "打印日期",
			"value": "【print_date】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_GROUP,
			"label": "小区",
			"value": "【group】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_ROOM,
			"label": "房号",
			"value": "【room】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_NAME,
			"label": "住户姓名",
			"value": "【name】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_MOBILE,
			"label": "住户手机号",
			"value": "【mobile】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_STATISTICS,
			"label": "合计",
			"value": "【staticstics】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_RDATE,
			"label": "收款日期",
			"value": "【rdate】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_PAYUSER,
			"label": "付款人",
			"value": "【pay_user】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_RECVUSER,
			"label": "收款人",
			"value": "【recv_user】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_RECVREMARK,
			"label": "收款备注",
			"value": "【recv_remark】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_RECVMETHOD,
			"label": "收款方式",
			"value": "【recv_method】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
		map[string]interface{}{
			"type":  CHARGE_RPINT_TPL_CONTENT_HF_ELEMENT_REMARK,
			"label": "说明",
			"value": "【remark】", //说明value是一个要被替换的变量, 这里用【】包裹就是说这个是要被替换到的变量的意思.
		},
	},
}
var CHARGE_INFO_MAP = map[int]map[int]map[string]interface{}{

	/*
		CHARGE_RPINT_TPL_BODY_RECEIVE_HEADER_ELEMENT: {
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_TITLE: {
				"value": "标题",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_CODE: {
				"value": "编号",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_LOGO: {
				"value": "LOGO",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_RPINTDATE: {
				"value": "打印日期",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_GROUP: {
				"value": "小区",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_ROOM: {
				"value": "房号",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_NAME: {
				"value": "住户姓名",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_MOBILE: {
				"value": "住户手机号",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_STATISTICS: {
				"value": "合计",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_RDATE: {
				"value": "收款日期",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_PAYUSER: {
				"value": "付款人",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_RECVUSER: {
				"value": "收款人",
			},
		},
		CHARGE_RPINT_TPL_BODY_RECEIVE_FOOTER_ELEMENT: {
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_RPINTDATE: {
				"value": "打印日期",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_GROUP: {
				"value": "小区",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_ROOM: {
				"value": "房号",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_NAME: {
				"value": "住户姓名",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_MOBILE: {
				"value": "住户手机号",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_STATISTICS: {
				"value": "合计",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_RDATE: {
				"value": "收款日期",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_PAYUSER: {
				"value": "付款人",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_RECVUSER: {
				"value": "收款人",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_REVCOMMENT: {
				"value": "收款备注",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_PAYMETHOD: {
				"value": "收款方式",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_DESC: {
				"value": "说明",
			},
		},
		CHARGE_RPINT_TPL_BODY_RECEIVE_CONTENT_ELEMENT: {
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_OBJECT: {
				"value": "内容区显示内容: 1房号/车位号/住户",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_BUILDING: {
				"value": "楼宇",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_UNIT: {
				"value": "单元",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_REVPROJECT: {
				"value": "收费项目",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_REVSTANDARD: {
				"value": "收费标准",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_STIME: {
				"value": "开始时间",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_ETIME: {
				"value": "结束时间",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_SMETER: {
				"value": "起度",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_EMETER: {
				"value": "止度",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_AMOUNT: {
				"value": "数量",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_PRICE: {
				"value": "单价",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_TOTAL: {
				"value": "金额",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_DISCOUNT: {
				"value": "优惠",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_LATEFEE: {
				"value": "违约金",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_SHOULDFEE: {
				"value": "应收合计",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_BATCH: {
				"value": "账单期数",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_REMARK: {
				"value": "备注",
			},
		},

		CHARGE_RPINT_TPL_BODY_NOTICE_HEADER_ELEMENT: {
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_TITLE: {
				"value": "标题",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_CODE: {
				"value": "编号",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_LOGO: {
				"value": "LOGO",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_RPINTDATE: {
				"value": "打印日期",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_GROUP: {
				"value": "小区",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_ROOM: {
				"value": "房号",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_NAME: {
				"value": "住户姓名",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_MOBILE: {
				"value": "住户手机号",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_STATISTICS: {
				"value": "合计",
			},
		},
		CHARGE_RPINT_TPL_BODY_NOTICE_FOOTER_ELEMENT: {
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_RPINTDATE: {
				"value": "打印日期",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_GROUP: {
				"value": "小区",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_ROOM: {
				"value": "6房号",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_NAME: {
				"value": "住户姓名",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_MOBILE: {
				"value": "住户手机号",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_STATISTICS: {
				"value": "合计",
			},
			CHARGE_RPINT_TPL_BODY_HF_ELEMENT_DESC: {
				"value": "说明",
			},
		},
		CHARGE_RPINT_TPL_BODY_NOTICE_CONTENT_ELEMENT: {
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_OBJECT: {
				"value": "房号/车位号/住户",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_BUILDING: {
				"value": "楼宇",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_UNIT: {
				"value": "单元",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_REVPROJECT: {
				"value": "收费项目",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_REVSTANDARD: {
				"value": "收费标准",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_STIME: {
				"value": "开始时间",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_ETIME: {
				"value": "结束时间",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_SMETER: {
				"value": "起度",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_EMETER: {
				"value": "止度",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_AMOUNT: {
				"value": "数量",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_PRICE: {
				"value": "单价",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_TOTAL: {
				"value": "金额",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_DISCOUNT: {
				"value": "优惠",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_LATEFEE: {
				"value": "违约金",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_SHOULDFEE: {
				"value": "应收合计",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_BATCH: {
				"value": "账单期数",
			},
			CHARGE_RPINT_TPL_BODY_CONTENT_ELEMENT_REMARK: {
				"value": "备注",
			},
		},
	*/
	CHARGE_RPINT_TPL_TYPE: {
		CHARGE_RPINT_TPL_TYPE_NOTICE: {
			"value": "通知单模板",
		},
		CHARGE_RPINT_TPL_TYPE_RECEIVE: {
			"value": "收据模板",
		},
	},
	CHARGE_PROJECT_LESS_METHOD: {
		CHARGE_PROJECT_LESS_METHOD_CYCLE: {
			"value": "按一个周期算",
		},
		CHARGE_PROJECT_LESS_METHOD_DAY: {
			"value": "按实际天数算",
		},
	},
	CHARGE_PROJECT_LATEFEE_POINT: {
		CHARGE_PROJECT_LATEFEE_POINT_START: {
			"value": "账单开始时间",
		},
		CHARGE_PROJECT_LATEFEE_POINT_END: {
			"value": "账单结束时间",
		},
	},
	CHARGE_PROJECT_TYPE: {

		CHARGE_PROJECT_TYPE_BASIC_CYCLE: {
			"value": "周期类",
		},
		CHARGE_PROJECT_TYPE_BASIC_TMP: {
			"value": "临时类",
		},
		CHARGE_PROJECT_TYPE_BASIC_DEPOSIT: {
			"value": "押金类",
		},
	},
	CHARGE_PROJECT_FEE_PRECISE: {
		CHARGE_PROJECT_FEE_PRECISE_YUAN: {
			"value": "元",
		},
		CHARGE_PROJECT_FEE_PRECISE_JIAO: {
			"value": "角",
		},
		CHARGE_PROJECT_FEE_PRECISE_FEN: {
			"value": "分",
		},
	},
	CHARGE_PROJECT_DEFAULT_POINT: {
		CHARGE_PROJECT_DEFAULT_POINT_CUR: {
			"value": "当期收当期",
		},
		CHARGE_PROJECT_DEFAULT_POINT_AFTER: {
			"value": "当期收上期",
		},
		CHARGE_PROJECT_DEFAULT_POINT_BEFORE: {
			"value": "当期收下期",
		},
	},
	CHARGE_STANDARD_TOTAL_CALCULATION_METHOD: {
		CHARGE_STANDARD_TOTAL_CALCULATION_METHOD_UC: {
			"value": "单价*数量",
		},
		CHARGE_STANDARD_TOTAL_CALCULATION_METHOD_SINGLE: {
			"value": "每户单独输入",
		},
		CHARGE_STANDARD_TOTAL_CALCULATION_METHOD_FIXED: {
			"value": "固定金额",
		},
		CHARGE_STANDARD_TOTAL_CALCULATION_METHOD_CUSTOM: {
			"value": "自定义公式",
		},
	},
	CHARGE_STANDARD_COUNT_TYPE_TEMP_PLEDGE: {
		CHARGE_STANDARD_COUNT_TYPE_BUILD_AREA: {
			"value": "建筑面积(平方米)",
		},
		CHARGE_STANDARD_COUNT_TYPE_USE_AREA: {
			"value": "套内面积(平方米)",
		},
		CHARGE_STANDARD_COUNT_TYPE_POOL_AREA: {
			"value": "公摊面积(平方米)",
		},
		CHARGE_STANDARD_COUNT_TYPE_AMOUNT: {
			"value": "用量(止度-起度)",
		},
		CHARGE_STANDARD_COUNT_TYPE_PARKING_AREA: {
			"value": "车位面积(平方米)",
		},
		CHARGE_STANDARD_COUNT_TYPE_CUSTOM_INPUT: {
			"value": "自定义数量",
		},
	},
	CHARGE_STANDARD_COUNT_TYPE_CYCLE: {
		CHARGE_STANDARD_COUNT_TYPE_BUILD_AREA: {
			"value": "建筑面积(平方米)",
		},
		CHARGE_STANDARD_COUNT_TYPE_USE_AREA: {
			"value": "套内面积(平方米)",
		},
		CHARGE_STANDARD_COUNT_TYPE_POOL_AREA: {
			"value": "公摊面积(平方米)",
		},
		CHARGE_STANDARD_COUNT_TYPE_AMOUNT: {
			"value": "用量(止度-起度)",
		},
		CHARGE_STANDARD_COUNT_TYPE_PARKING_AREA: {
			"value": "车位面积(平方米)",
		},
	},
	CHARGE_PAY_METHOD_ALL: {
		CHARGE_PAY_METHOD_OFFLINE_CASH: {
			"value": "线下支付-现金",
		},
		CHARGE_PAY_METHOD_OFFLINE_POS: {
			"value": "线下支付-POS机刷卡",
		},
		CHARGE_PAY_METHOD_OFFLINE_WEIXIN: {
			"value": "线下支付-微信扫码",
		},
		CHARGE_PAY_METHOD_OFFLINE_ZHIFUBAO: {
			"value": "线下支付-支付宝扫码",
		},
		CHARGE_PAY_METHOD_OFFLINE_BANK: {
			"value": "线下支付-银行卡转账",
		},
		CHARGE_PAY_METHOD_ONLINE_WEIXIN: {
			"value": "线上支付-微信",
		},
		CHARGE_PAY_METHOD_ONLINE_ZHIFUBAO: {
			"value": "线上支付-支付宝",
		},
		CHARGE_PAY_METHOD_ONLINE_BANK_COLLECTION: {
			"value": "线上支付-银行托收",
		},
		CHARGE_PAY_METHOD_OTHER: {
			"value": "其他收款方式",
		},
	},
	CHARGE_PAY_METHOD: {
		CHARGE_PAY_METHOD_OFFLINE_CASH: {
			"value": "线下支付-现金",
		},
		CHARGE_PAY_METHOD_OFFLINE_POS: {
			"value": "线下支付-POS机刷卡",
		},
		CHARGE_PAY_METHOD_OFFLINE_WEIXIN: {
			"value": "线下支付-微信转账",
		},
		CHARGE_PAY_METHOD_OFFLINE_ZHIFUBAO: {
			"value": "线下支付-支付宝转账",
		},
		CHARGE_PAY_METHOD_OFFLINE_BANK: {
			"value": "线下支付-银行卡转账",
		},
		CHARGE_PAY_METHOD_OTHER: {
			"value": "其他收款方式",
		},
	},
	CHARGE_PAY_CYCLE: {
		CHARGE_PAY_CYCLE_MONTH: {
			"value": "每月",
		},
		CHARGE_PAY_CYCLE_SEASON: {
			"value": "每季度",
		},
		CHARGE_PAY_CYCLE_HALFYEAR: {
			"value": "每半年",
		},
		CHARGE_PAY_CYCLE_YEAR: {
			"value": "每年",
		},
	},
}

//获取前端的展示配置
func GetChargeFEConf(which int) []map[string]interface{} {

	slice := make([]map[string]interface{}, 0)

	// To store the keys in slice in sorted order
	var keys []int
	for k := range CHARGE_INFO_MAP[which] {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// To perform the opertion you want
	for _, k := range keys {
		v := CHARGE_INFO_MAP[which][k]
		item := map[string]interface{}{
			"value": k,
			"label": v["value"],
		}
		slice = append(slice, item)
	}
	return slice
}

func GetChargePayCycleFEConf() []map[string]interface{} {

	return GetChargeFEConf(CHARGE_PAY_CYCLE)
}

func GetAllChargePayMethodFEConf() []map[string]interface{} {

	return GetChargeFEConf(CHARGE_PAY_METHOD_ALL)
}

func GetChargePayMethodFEConf() []map[string]interface{} {

	return GetChargeFEConf(CHARGE_PAY_METHOD)
}

func GetChargeTypeFEConf() []map[string]interface{} {

	return GetChargeFEConf(CHARGE_PROJECT_TYPE)
}

func GetChargeFeePreciseFEConf() []map[string]interface{} {

	return GetChargeFEConf(CHARGE_PROJECT_FEE_PRECISE)
}

func GetChargeDefaultPointFEConf() []map[string]interface{} {

	return GetChargeFEConf(CHARGE_PROJECT_DEFAULT_POINT)
}

func GetChargeLessMethodFEConf() []map[string]interface{} {

	return GetChargeFEConf(CHARGE_PROJECT_LESS_METHOD)
}

func GetChargeLatefeePointFEConf() []map[string]interface{} {

	return GetChargeFEConf(CHARGE_PROJECT_LATEFEE_POINT)
}

func GetChargeTotalCalculationMethodFEConf() []map[string]interface{} {

	return GetChargeFEConf(CHARGE_STANDARD_TOTAL_CALCULATION_METHOD)
}

//获取周期性收费标准的 计量方式
func GetChargeCountTypeCycleFEConf() []map[string]interface{} {

	return GetChargeFEConf(CHARGE_STANDARD_COUNT_TYPE_CYCLE)
}

//获取临时性/押金性性收费标准的 计量方式
func GetChargeCountTypeTempPledgeFEConf() []map[string]interface{} {

	return GetChargeFEConf(CHARGE_STANDARD_COUNT_TYPE_TEMP_PLEDGE)
}

//打印模板类型 1通知单模板、2收据模板
func GetChargePrintTplTypeFEConf() []map[string]interface{} {
	return GetChargeFEConf(CHARGE_RPINT_TPL_TYPE)
}

//获取前端的展示配置
/*
var CHARGE_RPINT_TPL_INFO = map[int]interface{}{
	CHARGE_RPINT_TPL_CONTENT_RECEIVE_HEADER_ELEMENT: []interface{}{
		map[string]interface{}{
			"type":  "title",
*/

//= 14 //收据-页眉显示内容
func GetChargePrintTplContentReceiveHeaderElementFEConf() []map[string]interface{} {
	//return GetChargeFEConf(CHARGE_RPINT_TPL_BODY_RECEIVE_HEADER_ELEMENT)
	return CHARGE_RPINT_TPL_INFO[CHARGE_RPINT_TPL_CONTENT_RECEIVE_HEADER_ELEMENT]
}

//= 15 //收据-页脚显示内容
func GetChargePrintTplContentReceiveFooterElementFEConf() []map[string]interface{} {
	//return GetChargeFEConf(CHARGE_RPINT_TPL_BODY_RECEIVE_FOOTER_ELEMENT)
	return CHARGE_RPINT_TPL_INFO[CHARGE_RPINT_TPL_CONTENT_RECEIVE_FOOTER_ELEMENT]
}

//= 16 //收据-内容区显示内容
func GetChargePrintTplContentReceiveBodyElementFEConf() []map[string]interface{} {
	//return GetChargeFEConf(CHARGE_RPINT_TPL_BODY_RECEIVE_CONTENT_ELEMENT)
	return CHARGE_RPINT_TPL_INFO[CHARGE_RPINT_TPL_CONTENT_RECEIVE_BODY_ELEMENT]
}

/*
//= 17 //通知-页眉显示内容
func GetChargePrintTplBodyNoticeHeaderElementFEConf() []map[string]interface{} {
	return GetChargeFEConf(CHARGE_RPINT_TPL_BODY_NOTICE_HEADER_ELEMENT)
}

//= 18 //通知-页脚显示内容
func GetChargePrintTplBodyNoticeFooterElementFEConf() []map[string]interface{} {
	return GetChargeFEConf(CHARGE_RPINT_TPL_BODY_NOTICE_FOOTER_ELEMENT)
}

//= 19 //通知-内容区显示内容
func GetChargePrintTplBodyNoticeContentElementFEConf() []map[string]interface{} {
	return GetChargeFEConf(CHARGE_RPINT_TPL_BODY_NOTICE_CONTENT_ELEMENT)
}
*/
