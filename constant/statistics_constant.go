package constant

import "sort"

const (
	STATISTICS_ROOM_X1 = 1 // 房屋统计：横轴x1类型
	STATISTICS_ROOM_X2 = 2 // 房屋统计：横轴x2类型
	STATISTICS_ROOM_Y  = 3 // 房屋统计：纵轴y类型

	STATISTICS_USER_X = 4 // 住户统计：横轴x类型
	STATISTICS_USER_Y = 5 // 住户统计：纵轴y类型

	STATISTICS_ROOM_X1_PROJECT = 1 //项目
	STATISTICS_ROOM_X1_GROUP   = 2 //组团

	STATISTICS_ROOM_X2_AREA_CONSTRUCTION = 1 //建筑面积
	STATISTICS_ROOM_X2_AREA_INNER        = 2 //套内面积

	STATISTICS_ROOM_Y_ROOM_CNT          = 1 // 房屋数量
	STATISTICS_ROOM_Y_AREA_CONSTRUCTION = 2 // 建筑面积
	STATISTICS_ROOM_Y_AREA_INNER        = 3 // 套内面积

	STATISTICS_USER_X_SEXY   = 1 //性别
	STATISTICS_USER_X_AGE    = 2 //年龄段
	STATISTICS_USER_X_JOB    = 3 //工作职位
	STATISTICS_USER_X_EDU    = 4 //教育程度
	STATISTICS_USER_X_NATION = 5 //民族

	STATISTICS_USER_Y_USER_CNT = 1 // 住户数量
)

var STATISTICS_MAP = map[int]map[int]map[string]interface{}{
	STATISTICS_ROOM_X1: {
		STATISTICS_ROOM_X1_PROJECT: {
			"value": "项目",
			"index": "belong_project_id",
		},
		STATISTICS_ROOM_X1_GROUP: {
			"value": "组团",
			"index": "belong_group_id",
		},
	},
	STATISTICS_ROOM_X2: {
		STATISTICS_ROOM_X2_AREA_CONSTRUCTION: {
			"value": "建筑面积",
			"index": "area_construction_section",
		},
		STATISTICS_ROOM_X2_AREA_INNER: {
			"value": "套内面积",
			"index": "area_inner_section",
		},
	},
	STATISTICS_ROOM_Y: {
		STATISTICS_ROOM_Y_ROOM_CNT: {
			"value": "总房屋数量",
			"index": "id", //count
		},
		STATISTICS_ROOM_Y_AREA_CONSTRUCTION: {
			"value": "总建筑面积",
			"index": "area_construction", //sum
		},
		STATISTICS_ROOM_Y_AREA_INNER: {
			"value": "总套内面积",
			"index": "area_inner", //sum
		},
	},
	STATISTICS_USER_X: {
		STATISTICS_USER_X_SEXY: {
			"value": "性别",
			"index": "basic_gender",
		},
		STATISTICS_USER_X_AGE: {
			"value": "年龄段",
			"index": "extra_age_group",
		},
		STATISTICS_USER_X_JOB: {
			"value": "工作职位",
			"index": "basic_work_type",
		},
		STATISTICS_USER_X_EDU: {
			"value": "教育程度",
			"index": "extra_edu_level",
		},
		STATISTICS_USER_X_NATION: {
			"value": "民族",
			"index": "extra_nation",
		},
	},
	STATISTICS_USER_Y: {
		STATISTICS_USER_Y_USER_CNT: {
			"value": "住户数量",
			"index": "id", //count
		},
	},
}

//获取前端的展示配置
func GetStatisticsFEConf(which int) []map[string]interface{} {

	slice := make([]map[string]interface{}, 0)

	// To store the keys in slice in sorted order
	var keys []int
	for k := range STATISTICS_MAP[which] {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// To perform the opertion you want
	for _, k := range keys {
		v := STATISTICS_MAP[which][k]

		item := map[string]interface{}{
			"value": k,
			"label": v["value"],
		}
		slice = append(slice, item)
	}
	return slice
}

func GetStatisticsRoomX1FEConf() []map[string]interface{} {

	return GetStatisticsFEConf(STATISTICS_ROOM_X1)
}

func GetStatisticsRoomX2FEConf() []map[string]interface{} {

	return GetStatisticsFEConf(STATISTICS_ROOM_X2)
}

func GetStatisticsRoomYFEConf() []map[string]interface{} {

	return GetStatisticsFEConf(STATISTICS_ROOM_Y)
}

func GetStatisticsUserXFEConf() []map[string]interface{} {

	return GetStatisticsFEConf(STATISTICS_USER_X)
}

func GetStatisticsUserYFEConf() []map[string]interface{} {

	return GetStatisticsFEConf(STATISTICS_USER_Y)
}
