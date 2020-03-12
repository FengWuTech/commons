package constant

import "sort"

const (
	//住户身份
	PARKING_SPACE_TYPE   = 1 //车位类型
	PARKING_SPACE_STATUS = 2 //车位状态

	PARKING_SPACE_TYPE_CIVIL    = 1 //人防车位
	PARKING_SPACE_TYPE_PUBLIC   = 2 //公共车位
	PARKING_SPACE_TYPE_PROPERTY = 3 //产权车位

	PARKING_SPACE_STATUS_SALED    = 1 //已售
	PARKING_SPACE_STATUS_RENT     = 2 //已租
	PARKING_SPACE_STATUS_EMPTY    = 3 //空置
	PARKING_SPACE_STATUS_SELFUSE  = 4 //自用
	PARKING_SPACE_STATUS_WAITSALE = 5 //待售
	PARKING_SPACE_STATUS_WAITRENT = 6 //待租
)

var SPACE_INFO_MAP = map[int]map[int]map[string]interface{}{

	PARKING_SPACE_TYPE: {

		PARKING_SPACE_TYPE_CIVIL: {
			"value": "人防车位",
		},
		PARKING_SPACE_TYPE_PUBLIC: {
			"value": "公共车位",
		},
		PARKING_SPACE_TYPE_PROPERTY: {
			"value": "产权车位",
		},
	},
	PARKING_SPACE_STATUS: {
		PARKING_SPACE_STATUS_SALED: {
			"value": "已售",
		},
		PARKING_SPACE_STATUS_RENT: {
			"value": "已租",
		},
		PARKING_SPACE_STATUS_EMPTY: {
			"value": "空置",
		},
		PARKING_SPACE_STATUS_SELFUSE: {
			"value": "自用",
		},
		PARKING_SPACE_STATUS_WAITSALE: {
			"value": "待售",
		},
		PARKING_SPACE_STATUS_WAITRENT: {
			"value": "待租",
		},
	},
}

//获取前端的展示配置
func GetParkingFEConf(which int) []map[string]interface{} {

	slice := make([]map[string]interface{}, 0)

	// To store the keys in slice in sorted order
	var keys []int
	for k := range SPACE_INFO_MAP[which] {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// To perform the opertion you want
	for _, k := range keys {
		v := SPACE_INFO_MAP[which][k]

		item := map[string]interface{}{
			"value": k,
			"label": v["value"],
		}
		slice = append(slice, item)
	}
	return slice
}

func GetParkingTypeFEConf() []map[string]interface{} {

	return GetParkingFEConf(PARKING_SPACE_TYPE)
}

func GetParkingStatusFEConf() []map[string]interface{} {

	return GetParkingFEConf(PARKING_SPACE_STATUS)
}
