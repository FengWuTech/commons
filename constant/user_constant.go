package constant

import "sort"

const (
	KEY_USER_STATUS           = 1
	KEY_USER_ROLE             = 2
	KEY_USER_WORK_TYPE        = 3
	KEY_USER_CERTIFICATE_TYPE = 4
	KEY_USER_EDU_LEVEL        = 5
	KEY_USER_AGE_GROUP        = 6
	KEY_USER_SEX              = 7
	KEY_USER_CHECK_ACTION     = 8

	//住户状态
	USER_STATUS_CHECKIN  = 1 //已迁入
	USER_STATUS_AUDIT    = 2 //待审核
	USER_STATUS_FAIL     = 3 //未通过
	USER_STATUS_CHECKOUT = 4 //已迁出
	USER_STATUS_OTHER    = 5 //其他

	//住户身份
	USER_ROLE_OWNER      = 1 //业主本人
	USER_ROLE_KIN        = 2 //亲戚
	USER_ROLE_TENANT     = 3 //租客
	USER_ROLE_FRINED     = 4 //朋友
	USER_ROLE_COLLEAGUE  = 5 //同事
	USER_ROLE_NANNY      = 6 //保姆
	USER_ROLE_DRIVER     = 7 //司机
	USER_ROLE_RENOVATION = 8 //装修人员
	USER_ROLE_OTHER      = 9 //其他

	//工作职业
	USER_WORK_UNKNOW  = 0 //未知
	USER_WORK_CIVIL   = 1 //公务员
	USER_WORK_TEACHER = 2 //教师
	USER_WORK_DOCTOR  = 3 //医生
	USER_WORK_NORMAL  = 4 //普通职员
	USER_WORK_SENIOR  = 5 //高级白领
	USER_WORK_PUBLIC  = 6 //事业单位

	//证件类型
	KEY_USER_CERTIFICATE_ID       = 0 //身份证
	KEY_USER_CERTIFICATE_PASSPORT = 1 //护照
	KEY_USER_CERTIFICATE_MILITARY = 2 //军官证
	KEY_USER_CERTIFICATE_HK       = 3 //港澳通行证
	KEY_USER_CERTIFICATE_TW       = 4 //台胞证

	//教育程度
	KEY_USER_EDU_UNKNOW        = 0 //未知
	KEY_USER_EDU_JUNIOR_MIDDLE = 1 //初中及以下
	KEY_USER_EDU_SENIOR_MIDDLE = 2 //高中/职校
	KEY_USER_EDU_JUNIOR_HIGH   = 3 //大专
	KEY_USER_EDU_SENIOR_HIGH   = 4 //本科
	KEY_USER_EDU_MASTER        = 5 //硕士
	KEY_USER_EDU_DOCTOR        = 6 //博士

	//年龄段
	KEY_USER_AGE_UNKNOW = 0 //未知
	KEY_USER_AGE_0      = 1 //25岁以下
	KEY_USER_AGE_25     = 2 //25-30岁
	KEY_USER_AGE_30     = 3 //31-35岁
	KEY_USER_AGE_35     = 4 //36-40岁
	KEY_USER_AGE_40     = 5 //41-45岁
	KEY_USER_AGE_45     = 6 //46-50岁
	KEY_USER_AGE_50     = 7 //51-55岁
	KEY_USER_AGE_55     = 8 //55岁以上

	//性别
	KEY_USER_SEX_MALE   = 0 //男
	KEY_USER_SEX_FEMALE = 1 //女
	KEY_USER_SEX_OTHER  = 2 //其他

	//操作动作
	KEY_USER_CHECK_ACTION_IN  = 0 //迁入
	KEY_USER_CHECK_ACTION_OUT = 1 //迁出

	USER_CHECKIN_AUDIT_TYPE_DEFAULT = 0 // 迁入审核来源 0-后台审核 1-业主邀请住户 2-业主主动审核
	USER_CHECKIN_AUDIT_TYPE_INVITE  = 1
	USER_CHECKIN_AUDIT_TYPE_APPROVE = 2
)

var USER_INFO_MAP = map[int]map[int]map[string]interface{}{

	KEY_USER_STATUS: {
		USER_STATUS_CHECKIN: {
			"value": "已迁入",
		},
		USER_STATUS_AUDIT: {
			"value": "待审核",
		},
		USER_STATUS_FAIL: {
			"value": "未通过",
		},
		USER_STATUS_CHECKOUT: {
			"value": "已迁出",
		},
		USER_STATUS_OTHER: {
			"value": "其他",
		},
	},
	KEY_USER_WORK_TYPE: {
		USER_WORK_UNKNOW: {
			"value": "未知",
		},
		USER_WORK_CIVIL: {
			"value": "公务员",
		},
		USER_WORK_TEACHER: {
			"value": "教师",
		},
		USER_WORK_DOCTOR: {
			"value": "医生",
		},
		USER_WORK_NORMAL: {
			"value": "普通职员",
		},
		USER_WORK_SENIOR: {
			"value": "高级白领",
		},
		USER_WORK_PUBLIC: {
			"value": "事业单位",
		},
	},
	KEY_USER_CERTIFICATE_TYPE: {
		KEY_USER_CERTIFICATE_ID: {
			"value": "身份证",
		},
		KEY_USER_CERTIFICATE_PASSPORT: {
			"value": "护照",
		},
		KEY_USER_CERTIFICATE_MILITARY: {
			"value": "军官证",
		},
		KEY_USER_CERTIFICATE_HK: {
			"value": "港澳通行证",
		},
		KEY_USER_CERTIFICATE_TW: {
			"value": "台胞证",
		},
	},
	KEY_USER_EDU_LEVEL: {
		KEY_USER_EDU_UNKNOW: {
			"value": "未知",
		},
		KEY_USER_EDU_JUNIOR_MIDDLE: {
			"value": "初中及以下",
		},
		KEY_USER_EDU_SENIOR_MIDDLE: {
			"value": "高中/职校",
		},
		KEY_USER_EDU_JUNIOR_HIGH: {
			"value": "大专",
		},
		KEY_USER_EDU_SENIOR_HIGH: {
			"value": "本科",
		},
		KEY_USER_EDU_MASTER: {
			"value": "硕士",
		},
		KEY_USER_EDU_DOCTOR: {
			"value": "博士",
		},
	},
	KEY_USER_AGE_GROUP: {
		KEY_USER_AGE_UNKNOW: {
			"value": "未知",
		},
		KEY_USER_AGE_0: {
			"value": "25岁以下",
		},
		KEY_USER_AGE_25: {
			"value": "25-30岁",
		},
		KEY_USER_AGE_30: {
			"value": "31-35岁",
		},
		KEY_USER_AGE_35: {
			"value": "36-40岁",
		},
		KEY_USER_AGE_40: {
			"value": "41-45岁",
		},
		KEY_USER_AGE_45: {
			"value": "46-50岁",
		},
		KEY_USER_AGE_50: {
			"value": "51-55岁",
		},
		KEY_USER_AGE_55: {
			"value": "55岁以上",
		},
	},
	KEY_USER_SEX: {
		KEY_USER_SEX_MALE: {
			"value": "男",
		},
		KEY_USER_SEX_FEMALE: {
			"value": "女",
		},
		KEY_USER_SEX_OTHER: {
			"value": "保密",
		},
	},
	KEY_USER_ROLE: {
		USER_ROLE_OWNER: {
			"value":      "业主",
			"validMonth": 0, //0代表无限
		},
		USER_ROLE_KIN: {
			"value":      "亲戚",
			"validMonth": 0, //0代表无限
		},
		USER_ROLE_TENANT: {
			"value":      "租客",
			"validMonth": 12, //12个月
		},
		USER_ROLE_FRINED: {
			"value":      "朋友",
			"validMonth": 12, //12个月
		},
		USER_ROLE_COLLEAGUE: {
			"value":      "同事",
			"validMonth": 3, //3个月
		},
		USER_ROLE_NANNY: {
			"value":      "保姆",
			"validMonth": 6, //6个月
		},
		USER_ROLE_DRIVER: {
			"value":      "司机",
			"validMonth": 6, //6个月
		},
		USER_ROLE_RENOVATION: {
			"value":      "装修人员",
			"validMonth": 3, //3个月
		},
		USER_ROLE_OTHER: {
			"value":      "其他",
			"validMonth": 3, //3个月
		},
	},
}

func GetUserRoleConstantByID(USER_ROLE_ID int) interface{} {

	if info, ok := USER_INFO_MAP[KEY_USER_ROLE][USER_ROLE_ID]; ok {
		return info
	}
	return nil
}

func GetAllUserConstant() interface{} {
	return USER_INFO_MAP
}

//获取前端的展示配置
func GetUserFEConf(which int) []map[string]interface{} {

	slice := make([]map[string]interface{}, 0)

	// To store the keys in slice in sorted order
	var keys []int
	for k := range USER_INFO_MAP[which] {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// To perform the opertion you want
	for _, k := range keys {
		v := USER_INFO_MAP[which][k]

		item := map[string]interface{}{
			"value": k,
			"label": v["value"],
		}
		slice = append(slice, item)
	}
	return slice
}

func GetUserRoleFEConf() []map[string]interface{} {
	var ret []map[string]interface{}
	for k, v := range USER_INFO_MAP[KEY_USER_ROLE] {
		ret = append(ret, map[string]interface{}{
			"label":      v["value"],
			"value":      k,
			"validMonth": v["validMonth"],
		})
	}
	return ret
}

func GetUserStatusFEConf() []map[string]interface{} {
	return GetUserFEConf(KEY_USER_STATUS)
}

func GetUserWorkTypeFEConf() []map[string]interface{} {
	return GetUserFEConf(KEY_USER_WORK_TYPE)
}

func GetUserEduLevelFEConf() []map[string]interface{} {
	return GetUserFEConf(KEY_USER_EDU_LEVEL)
}

func GetUserCertificateTypeFEConf() []map[string]interface{} {
	return GetUserFEConf(KEY_USER_CERTIFICATE_TYPE)
}

func GetUserUserAgeFEConf() []map[string]interface{} {
	return GetUserFEConf(KEY_USER_AGE_GROUP)
}

func GetUserUserSexFEConf() []map[string]interface{} {
	return GetUserFEConf(KEY_USER_SEX)
}

func confToMap(conf []map[string]interface{}) map[int]string {
	ret := map[int]string{}
	for _, item := range conf {
		k, v := -1, ""
		if value, ok := item["value"]; ok {
			k = value.(int)
		}
		if value, ok := item["label"]; ok {
			v = value.(string)
		}
		if k >= 0 {
			ret[k] = v
		}
	}
	return ret
}

var (
	UserSexyConfMap = confToMap(GetUserUserSexFEConf())
	UserAgeConfMap  = confToMap(GetUserUserAgeFEConf())
	UserEduConfMap  = confToMap(GetUserEduLevelFEConf())
	UserWorkConfMap = confToMap(GetUserWorkTypeFEConf())
)
