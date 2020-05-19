package constant

import "sort"

const (
	ESTATE_ROOM_AREA_SECTION = 1 //房屋的面积区间

	ESTATE_ROOM_AREA_SECTION_1 = 1 //50平及以下
	ESTATE_ROOM_AREA_SECTION_2 = 2 //50-90平
	ESTATE_ROOM_AREA_SECTION_3 = 3 //90-130平
	ESTATE_ROOM_AREA_SECTION_4 = 4 //130-200平
	ESTATE_ROOM_AREA_SECTION_5 = 5 //200-500平
	ESTATE_ROOM_AREA_SECTION_6 = 6 //>500平

	ESTATE_PROJECT_PHONE_CATEGORY_COMPANY   = 1 //项目电话分类-物业
	ESTATE_PROJECT_PHONE_CATEGORY_COMMUNITY = 2 //项目电话分类-社会
	ESTATE_PROJECT_PHONE_CATEGORY_AROUND    = 3 //项目电话分类-周边
)

var PHONE_CATEGORY = []map[string]interface{}{
	{
		"label": "物业紧急联络电话",
		"value": ESTATE_PROJECT_PHONE_CATEGORY_COMPANY,
	},
	{
		"label": "社会公共服务电话",
		"value": ESTATE_PROJECT_PHONE_CATEGORY_COMMUNITY,
	},
	{
		"label": "周边便民服务电话",
		"value": ESTATE_PROJECT_PHONE_CATEGORY_AROUND,
	},
}

var ESTATE_INFO_MAP = map[int]map[int]map[string]interface{}{

	ESTATE_ROOM_AREA_SECTION: {
		ESTATE_ROOM_AREA_SECTION_1: {
			"value":   "50平米及以下",
			"section": []float64{0, 50},
		},
		ESTATE_ROOM_AREA_SECTION_2: {
			"value":   "50-90平米",
			"section": []float64{50, 90},
		},
		ESTATE_ROOM_AREA_SECTION_3: {
			"value":   "90-130平米",
			"section": []float64{90, 130},
		},
		ESTATE_ROOM_AREA_SECTION_4: {
			"value":   "130-200平米",
			"section": []float64{130, 200},
		},
		ESTATE_ROOM_AREA_SECTION_5: {
			"value":   "200-500平米",
			"section": []float64{200, 500},
		},
		ESTATE_ROOM_AREA_SECTION_6: {
			"value":   ">500平米",
			"section": []float64{500},
		},
	},
}

//获取面积所在区间的ID
func GetAreaSection(area float64) int {

	// To store the keys in slice in sorted order
	var keys []int
	for k := range ESTATE_INFO_MAP[ESTATE_ROOM_AREA_SECTION] {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		sList := ESTATE_INFO_MAP[ESTATE_ROOM_AREA_SECTION][k]["section"]
		sectionList := sList.([]float64)

		if len(sectionList) == 2 {
			if area > sectionList[0] && area <= sectionList[1] {
				return k
			}
		} else if len(sectionList) == 1 {
			if area > sectionList[0] {
				return k
			}
		}
	}
	return 0
}

//获取前端的展示配置
func GetEstateFEConf(which int) []map[string]interface{} {

	slice := make([]map[string]interface{}, 0)
	// To store the keys in slice in sorted order
	var keys []int
	for k := range ESTATE_INFO_MAP[which] {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// To perform the opertion you want
	for _, k := range keys {
		v := ESTATE_INFO_MAP[which][k]

		item := map[string]interface{}{
			"value": k,
			"label": v["value"],
		}
		slice = append(slice, item)
	}
	return slice
}

func GetEstateSectionFEConf() []map[string]interface{} {

	return GetEstateFEConf(ESTATE_ROOM_AREA_SECTION)
}
