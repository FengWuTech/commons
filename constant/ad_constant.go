package constant

import "sort"

const (
	AD_POS = 1 //广告的位置

	AD_POS_TOP       = 1 //首页轮播图
	AD_POS_RECOMMEND = 2 //首页推荐位
	AD_POS_ARTICLE   = 3 //文章底部
	AD_POS_SERVICE   = 4 //服务业轮播图
)

var AD_INFO_MAP = map[int]map[int]map[string]interface{}{

	AD_POS: {
		AD_POS_TOP: {
			"value": "首页轮播图",
		},
		AD_POS_RECOMMEND: {
			"value": "首页推荐位",
		},
		AD_POS_ARTICLE: {
			"value": "文章底部",
		},
		AD_POS_SERVICE: {
			"value": "服务业轮播图",
		},
	},
}

//获取前端的展示配置
func GetAdFEConf(which int) []map[string]interface{} {

	slice := make([]map[string]interface{}, 0)
	// To store the keys in slice in sorted order
	var keys []int
	for k := range AD_INFO_MAP[which] {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// To perform the opertion you want
	for _, k := range keys {
		v := AD_INFO_MAP[which][k]

		item := map[string]interface{}{
			"value": k,
			"label": v["value"],
		}
		slice = append(slice, item)
	}
	return slice
}

func GetAdPosFEConf() []map[string]interface{} {

	return GetAdFEConf(AD_POS)
}
