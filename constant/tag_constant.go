package constant

import "sort"

const (
	TAG_TYPE = 1 //标签分类

	TAG_TYPE_ROOM     = 1 //房屋
	TAG_TYPE_PARKING  = 2 //车位
	TAG_TYPE_USER     = 3 //住户
	TAG_TYPE_BUILDING = 4 //楼宇
)

var TAG_INFO_MAP = map[int]map[int]map[string]interface{}{

	TAG_TYPE: {

		TAG_TYPE_ROOM: {
			"value": "房屋",
		},
		TAG_TYPE_PARKING: {
			"value": "车位",
		},
		TAG_TYPE_USER: {
			"value": "住户",
		},
		TAG_TYPE_BUILDING: {
			"value": "楼宇",
		},
	},
}

//获取前端的展示配置
func GetTagFEConf(which int) []map[string]interface{} {

	slice := make([]map[string]interface{}, 0)

	// To store the keys in slice in sorted order
	var keys []int
	for k := range TAG_INFO_MAP[which] {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// To perform the opertion you want
	for _, k := range keys {
		v := TAG_INFO_MAP[which][k]

		item := map[string]interface{}{
			"value": k,
			"label": v["value"],
		}
		slice = append(slice, item)
	}
	return slice
}

func GetTagTypeFEConf() []map[string]interface{} {

	return GetTagFEConf(TAG_TYPE)
}
