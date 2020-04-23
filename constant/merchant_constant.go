package constant

const (
	MERCHANT_TYPE_WEIXIN = 1

	MERCHANT_APPLY_STATUS_CREATED  = 1
	MERCHANT_APPLY_STATUS_APPLYING = 2
	MERCHANT_APPLY_STATUS_SUCCESS  = 3
	MERCHANT_APPLY_STATUS_FAIL     = 4
)

var MERCHANT_APPLY_STATUS = []map[string]interface{}{
	{
		"label": "已提交",
		"value": MERCHANT_APPLY_STATUS_CREATED,
	},
	{
		"label": "申请中",
		"value": MERCHANT_APPLY_STATUS_APPLYING,
	},
	{
		"label": "申请成功",
		"value": MERCHANT_APPLY_STATUS_SUCCESS,
	},
	{
		"label": "申请失败",
		"value": MERCHANT_APPLY_STATUS_FAIL,
	},
}
