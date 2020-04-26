package constant

const (
	//支付渠道
	CHARGE_PAY_METHOD_OFFLINE_CASH           = 1  //线下支付-现金
	CHARGE_PAY_METHOD_OFFLINE_POS            = 2  //线下支付-POS机刷卡
	CHARGE_PAY_METHOD_OFFLINE_WEIXIN         = 3  //线下支付-微信扫码
	CHARGE_PAY_METHOD_OFFLINE_ZHIFUBAO       = 4  //线下支付-支付宝扫码
	CHARGE_PAY_METHOD_OFFLINE_BANK           = 5  //线下支付-转账
	CHARGE_PAY_METHOD_OFFLINE_QRCODE_SCAN    = 12 //线下支付-聚合二维码
	CHARGE_PAY_METHOD_ONLINE_WEIXIN          = 6  //线上支付-微信
	CHARGE_PAY_METHOD_ONLINE_ZHIFUBAO        = 7  //线上支付-支付宝
	CHARGE_PAY_METHOD_ONLINE_BANK_COLLECTION = 8  //线上支付-银行托收
	CHARGE_PAY_METHOD_OTHER                  = 9  //其他收款方式
	CHARGE_PAY_METHOD_DEDUCTE                = 10 //预充值抵扣（不传给fe）

	CHARGE_PAY_METHOD_NAME_OFFLINE_CASH           = "线下-现金支付"
	CHARGE_PAY_METHOD_NAME_OFFLINE_POS            = "线下-POS机刷卡"
	CHARGE_PAY_METHOD_NAME_OFFLINE_WEIXIN         = "线下-微信专用支付码"
	CHARGE_PAY_METHOD_NAME_OFFLINE_ZHIFUBAO       = "线下-支付宝专用支付码"
	CHARGE_PAY_METHOD_NAME_OFFLINE_QRCODE_SCAN    = "线下-聚合支付码"
	CHARGE_PAY_METHOD_NAME_OFFLINE_BANK           = "线下-转账"
	CHARGE_PAY_METHOD_NAME_ONLINE_WEIXIN          = "线上-微信公众号"
	CHARGE_PAY_METHOD_NAME_ONLINE_ZHIFUBAO        = "线上-支付宝小程序"
	CHARGE_PAY_METHOD_NAME_ONLINE_BANK_COLLECTION = "线上-银行托收"
	CHARGE_PAY_METHOD_NAME_OTHER                  = "其他收款方式"
	CHARGE_PAY_METHOD_NAME_DEDUCTE                = "预充值抵扣"
)

//所有的支付渠道
var CHARGE_PAY_METHOD_ALL = map[int]string{
	CHARGE_PAY_METHOD_OFFLINE_CASH:           CHARGE_PAY_METHOD_NAME_OFFLINE_CASH,
	CHARGE_PAY_METHOD_OFFLINE_POS:            CHARGE_PAY_METHOD_NAME_OFFLINE_POS,
	CHARGE_PAY_METHOD_OFFLINE_WEIXIN:         CHARGE_PAY_METHOD_NAME_OFFLINE_WEIXIN,
	CHARGE_PAY_METHOD_OFFLINE_ZHIFUBAO:       CHARGE_PAY_METHOD_NAME_OFFLINE_ZHIFUBAO,
	CHARGE_PAY_METHOD_OFFLINE_BANK:           CHARGE_PAY_METHOD_NAME_OFFLINE_BANK,
	CHARGE_PAY_METHOD_OFFLINE_QRCODE_SCAN:    CHARGE_PAY_METHOD_NAME_OFFLINE_QRCODE_SCAN,
	CHARGE_PAY_METHOD_ONLINE_WEIXIN:          CHARGE_PAY_METHOD_NAME_ONLINE_WEIXIN,
	CHARGE_PAY_METHOD_ONLINE_ZHIFUBAO:        CHARGE_PAY_METHOD_NAME_ONLINE_ZHIFUBAO,
	CHARGE_PAY_METHOD_ONLINE_BANK_COLLECTION: CHARGE_PAY_METHOD_NAME_ONLINE_BANK_COLLECTION,
	CHARGE_PAY_METHOD_OTHER:                  CHARGE_PAY_METHOD_NAME_OTHER,
}

//后台支付可以选择的支付类型
var CHARGE_PAY_METHOD = []int{
	CHARGE_PAY_METHOD_OFFLINE_CASH,
	CHARGE_PAY_METHOD_OFFLINE_POS,
	CHARGE_PAY_METHOD_OFFLINE_WEIXIN,
	CHARGE_PAY_METHOD_OFFLINE_ZHIFUBAO,
	CHARGE_PAY_METHOD_OFFLINE_BANK,
	CHARGE_PAY_METHOD_OFFLINE_QRCODE_SCAN,
}

//可以配置项目关联的渠道
var THIRD_PAY_METHOD = []int{
	CHARGE_PAY_METHOD_OFFLINE_WEIXIN,
	CHARGE_PAY_METHOD_OFFLINE_ZHIFUBAO,
	CHARGE_PAY_METHOD_OFFLINE_POS,
	CHARGE_PAY_METHOD_OFFLINE_QRCODE_SCAN,
	CHARGE_PAY_METHOD_OFFLINE_BANK,
}

//允许修改真实收款金额的渠道
var PAY_METHOD_ALLOW_EDIT_REAL_AMOUNT = []int{
	CHARGE_PAY_METHOD_OFFLINE_CASH,
	CHARGE_PAY_METHOD_OFFLINE_WEIXIN,
	CHARGE_PAY_METHOD_OFFLINE_ZHIFUBAO,
	CHARGE_PAY_METHOD_OFFLINE_QRCODE_SCAN,
	CHARGE_PAY_METHOD_OFFLINE_POS,
	CHARGE_PAY_METHOD_OFFLINE_BANK,
}

func GetPayMethodName(methodID int) string {
	name, ok := CHARGE_PAY_METHOD_ALL[methodID]
	if ok {
		return name
	} else {
		return "未知"
	}
}

func GetAllChargePayMethodFEConf() []map[string]interface{} {
	var ret []map[string]interface{}
	for value, label := range CHARGE_PAY_METHOD_ALL {
		ret = append(ret, map[string]interface{}{
			"label": label,
			"value": value,
		})
	}
	return ret
}

func GetChargePayMethodFEConf() []map[string]interface{} {
	var ret []map[string]interface{}
	for _, id := range CHARGE_PAY_METHOD {
		ret = append(ret, map[string]interface{}{
			"label": GetPayMethodName(id),
			"value": id,
		})
	}
	return ret
}

func GetThirdPayMethodFEConf() []map[string]interface{} {
	var ret []map[string]interface{}
	for _, id := range THIRD_PAY_METHOD {
		ret = append(ret, map[string]interface{}{
			"label": GetPayMethodName(id),
			"value": id,
		})
	}
	return ret
}
