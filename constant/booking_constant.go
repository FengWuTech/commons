package constant

import "sort"

const (
	BookingNeedDateYes = 1
	BookingNeedDateNo  = 0
)

const (
	BookingOrderStatusNew = iota
	BookingOrderStatusConfirmed
	BookingOrderStatusDone
	BookingOrderStatusCanceled
)

const (
	BOOKING_TYPE = 1 //服务分类

	BOOKING_TYPE_LIGHT       = 1 //灯具电路
	BOOKING_TYPE_CLEANING    = 2 //居家保洁
	BOOKING_TYPE_KITCHEN     = 3 //厨卫洁具
	BOOKING_TYPE_ELECTRONICS = 4 //家电维修
	BOOKING_TYPE_COMPUTER    = 5 //电脑网络
	BOOKING_TYPE_PIPELINE    = 6 //管道疏通
	BOOKING_TYPE_FURNITURE   = 7 //家具门窗
	BOOKING_TYPE_HEATING     = 8 //地热暖气
	BOOKING_TYPE_FLOOR       = 9 //地面墙面
)

var BOOKING_INFO_MAP = map[int]map[int]map[string]interface{}{

	BOOKING_TYPE: {
		BOOKING_TYPE_LIGHT: {
			"value": "灯具电路",
		},
		BOOKING_TYPE_CLEANING: {
			"value": "居家保洁",
		},
		BOOKING_TYPE_KITCHEN: {
			"value": "厨卫洁具",
		},
		BOOKING_TYPE_ELECTRONICS: {
			"value": "家电维修",
		},
		BOOKING_TYPE_COMPUTER: {
			"value": "电脑网络",
		},
		BOOKING_TYPE_PIPELINE: {
			"value": "管道疏通",
		},
		BOOKING_TYPE_FURNITURE: {
			"value": "家具门窗",
		},
		BOOKING_TYPE_HEATING: {
			"value": "地热暖气",
		},
		BOOKING_TYPE_FLOOR: {
			"value": "地面墙面",
		},
	},
}

//获取前端的展示配置
func GetBookingFEConf(which int) []map[string]interface{} {

	slice := make([]map[string]interface{}, 0)

	// To store the keys in slice in sorted order
	var keys []int
	for k := range BOOKING_INFO_MAP[which] {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// To perform the opertion you want
	for _, k := range keys {
		v := BOOKING_INFO_MAP[which][k]

		item := map[string]interface{}{
			"value": k,
			"label": v["value"],
		}
		slice = append(slice, item)
	}
	return slice
}

func GetBookingTypeFEConf() []map[string]interface{} {
	return GetBookingFEConf(BOOKING_TYPE)
}

func BookingTypeToName(type_ int) string {
       types := BOOKING_INFO_MAP[BOOKING_TYPE]
       for k, v := range types {
               value, _ := v["value"]
               if k == type_ {
                       return value.(string)
               }
       }
       return "其他"
}