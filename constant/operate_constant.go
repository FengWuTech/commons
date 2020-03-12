package constant

import "sort"

const (
	OPERATE_TYPE = 1 //操作类型

	//操作类型
	OPERATE_TYPE_LOGIN             = 0 //登录
	OPERATE_TYPE_STAFF_ROLE_GRANT  = 1 //增加管理员
	OPERATE_TYPE_STAFF_ROLE_CANCEL = 2 //删除管理员
	OPERATE_TYPE_STAFF_SET_GROUP   = 3 //员工关联组团
	OPERATE_TYPE_STAFF_SET_ACL     = 4 //员工权限
	OPERATE_TYPE_BASIC_INFO_EDIT   = 5 //编辑基础信息
	OPERATE_TYPE_USER_BIND         = 6 //住户绑定
	OPERATE_TYPE_IMPORT_EXPORT     = 7 //导入导出
)

var OPERATE_HISTORY_MAP = map[int]map[int]map[string]interface{}{
	OPERATE_TYPE: {
		OPERATE_TYPE_LOGIN: {
			"value": "登录",
			"tpl":   "%v%v登录了系统",
		},
		OPERATE_TYPE_STAFF_ROLE_GRANT: {
			"value": "增加管理员",
		},
		OPERATE_TYPE_STAFF_ROLE_CANCEL: {
			"value": "删除管理员",
		},
		OPERATE_TYPE_STAFF_SET_GROUP: {
			"value": "员工关联组团",
		},
		OPERATE_TYPE_STAFF_SET_ACL: {
			"value": "员工权限",
		},
		OPERATE_TYPE_BASIC_INFO_EDIT: {
			"value": "编辑基础信息",
		},
		OPERATE_TYPE_USER_BIND: {
			"value": "住户绑定",
		},
		OPERATE_TYPE_IMPORT_EXPORT: {
			"value": "导入导出",
		},
	},
}

//获取前端的展示配置
func GetOperateFEConf(which int) []map[string]interface{} {

	slice := make([]map[string]interface{}, 0)

	// To store the keys in slice in sorted order
	var keys []int
	for k := range OPERATE_HISTORY_MAP[which] {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// To perform the operation you want
	for _, k := range keys {
		v := OPERATE_HISTORY_MAP[which][k]

		item := map[string]interface{}{
			"value": k,
			"label": v["value"],
		}
		slice = append(slice, item)
	}
	return slice
}

func GetOperateTypeFEConf() []map[string]interface{} {

	return GetOperateFEConf(OPERATE_TYPE)
}
