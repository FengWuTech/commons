package constant

import "sort"

const (
	METER_SHARE_TYPE = 1 //分摊方式

	METER_SHARE_TYPE_ROOM   = 1 //按户
	METER_SHARE_TYPE_AREA   = 2 //按建筑面积
	METER_SHARE_TYPE_AMOUNT = 3 //按用量
)

var METER_INFO_MAP = map[int]map[int]map[string]interface{}{

	METER_SHARE_TYPE: {
		METER_SHARE_TYPE_ROOM: {
			"value": "按户",
			"type":  0, //户摊
		},
		METER_SHARE_TYPE_AREA: {
			"value": "按建筑面积(平方米)",
			"type":  0,
		},
		METER_SHARE_TYPE_AMOUNT: {
			"value": "按用量",
			"type":  1, //表摊
		},
	},
}

//获取前端的展示配置
func GetMeterFEConf(which int) []map[string]interface{} {

	slice := make([]map[string]interface{}, 0)

	// To store the keys in slice in sorted order
	var keys []int
	for k := range METER_INFO_MAP[which] {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// To perform the opertion you want
	for _, k := range keys {
		v := METER_INFO_MAP[which][k]

		item := map[string]interface{}{
			"value": k,
			"label": v["value"],
			"type":  v["type"],
		}
		slice = append(slice, item)
	}
	return slice
}

func GetMeterShareTypeFEConf() []map[string]interface{} {

	return GetMeterFEConf(METER_SHARE_TYPE)
}
