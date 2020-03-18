package constant

import "sort"

const (
	KEY_VISIT_REASON = 1

	//住户状态
	VISIT_REASON_RELATIVE = 1 //亲朋拜访
	VISIT_REASON_PARTY    = 2 //聚会
	VISIT_REASON_DECORATE = 3 //装修方形
	VISIT_REASON_WAIMAI   = 4 //外卖
	VISIT_REASON_CLEAN    = 5 //保洁
	VISIT_REASON_LEARN    = 6 //家教
	VISIT_REASON_OTHER    = 7 //其他
)

var VISIT_INFO_MAP = map[int]map[int]map[string]interface{}{

	KEY_VISIT_REASON: {
		VISIT_REASON_RELATIVE: {
			"value": "亲朋拜访",
		},
		VISIT_REASON_PARTY: {
			"value": "聚会",
		},
		VISIT_REASON_DECORATE: {
			"value": "装修放行",
		},
		VISIT_REASON_WAIMAI: {
			"value": "外卖送餐",
		},
		VISIT_REASON_CLEAN: {
			"value": "保洁",
		},
		VISIT_REASON_LEARN: {
			"value": "家教",
		},
		VISIT_REASON_OTHER: {
			"value": "其他",
		},
	},
}

//获取前端的展示配置
func GetVisitFEConf(which int) []map[string]interface{} {

	slice := make([]map[string]interface{}, 0)

	// To store the keys in slice in sorted order
	var keys []int
	for k := range VISIT_INFO_MAP[which] {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// To perform the opertion you want
	for _, k := range keys {
		v := VISIT_INFO_MAP[which][k]

		item := map[string]interface{}{
			"value": k,
			"label": v["value"],
		}
		slice = append(slice, item)
	}
	return slice
}

func GetVisitReasonFEConf() []map[string]interface{} {
	return GetVisitFEConf(KEY_VISIT_REASON)
}
