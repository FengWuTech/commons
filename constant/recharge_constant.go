package constant

import "sort"

const (
	RECHARGE_TYPE = 1 //操作类型

	//操作类型
	RECHARGE_TYPE_ADD               = 1 //充值
	RECHARGE_TYPE_REFUND_TO_ACCOUNT = 2 //退费
	RECHARGE_TYPE_DEDUCT            = 3 //抵扣
	RECHARGE_TYPE_REFUND_TO_USER    = 4 //退款

)

var RECHARGE_MAP = map[int]map[int]map[string]interface{}{
	RECHARGE_TYPE: {
		RECHARGE_TYPE_ADD: {
			"value": "充值",
		},
		RECHARGE_TYPE_REFUND_TO_ACCOUNT: {
			"value": "退费",
		},
		RECHARGE_TYPE_DEDUCT: {
			"value": "抵扣",
		},
		RECHARGE_TYPE_REFUND_TO_USER: {
			"value": "退款",
		},
	},
}

//获取前端的展示配置
func GetRechargeFEConf(which int) []map[string]interface{} {

	slice := make([]map[string]interface{}, 0)

	// To store the keys in slice in sorted order
	var keys []int
	for k := range RECHARGE_MAP[which] {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// To perform the opertion you want
	for _, k := range keys {
		v := RECHARGE_MAP[which][k]

		item := map[string]interface{}{
			"value": k,
			"label": v["value"],
		}
		slice = append(slice, item)
	}
	return slice
}

func GetRechargeTypeFEConf() []map[string]interface{} {

	return GetRechargeFEConf(OPERATE_TYPE)
}
