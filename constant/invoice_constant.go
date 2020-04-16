package constant

import "sort"

const (
	INVOICE_STATUS               = 0
	INVOICE_STATUS_START         = 1 //申请中
	INVOICE_STATUS_DONE          = 2 //开票成功
	INVOICE_STATUS_FAIL          = 3 //开票失败
	INVOICE_STATUS_CANCEL_DONE   = 4 //销票成功
	INVOICE_STATUS_CANCEL_FAIL   = 5 //销票失败

	SPECIAL_INVOICE_STATUS       = 2
	SPECIAL_INVOICE_STATUS_START = 6 //专票待处理
	SPECIAL_INVOICE_STATUS_DONE  = 7 //专票已开票

	INVOICE_TYPE                 = 1
	INVOICE_TYPE_PERSONAL        = 1 //个人
	INVOICE_TYPE_COMPANY         = 2 //公司普票
	INVOICE_TYPE_COMPANY_SPECIAL = 3 //公司专票
)

var INVOICE_INFO_MAP = map[int]map[int]map[string]interface{}{
	INVOICE_STATUS: {
		INVOICE_STATUS_START: {
			"value": "申请中",
		},
		INVOICE_STATUS_DONE: {
			"value": "已开票",
		},
		INVOICE_STATUS_FAIL: {
			"value": "开票失败",
		},
		INVOICE_STATUS_CANCEL_DONE: {
			"value": "已销票",
		},
		INVOICE_STATUS_CANCEL_FAIL: {
			"value": "销票失败",
		},
	},
	SPECIAL_INVOICE_STATUS: {
		SPECIAL_INVOICE_STATUS_START: {
			"value": "专票待处理",
		},
		SPECIAL_INVOICE_STATUS_DONE: {
			"value": "专票已开票",
		},
	},
	INVOICE_TYPE: {
		INVOICE_TYPE_PERSONAL: {
			"value": "个人",
		},
		INVOICE_TYPE_COMPANY: {
			"value": "公司普票",
		},
		INVOICE_TYPE_COMPANY_SPECIAL: {
			"value": "公司专票",
		},
	},
}

//获取前端的展示配置
func GetInvoiceFEConf(which int) []map[string]interface{} {

	slice := make([]map[string]interface{}, 0)

	// To store the keys in slice in sorted order
	var keys []int
	for k := range INVOICE_INFO_MAP[which] {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// To perform the opertion you want
	for _, k := range keys {
		v := INVOICE_INFO_MAP[which][k]

		item := map[string]interface{}{
			"value": k,
			"label": v["value"],
		}
		slice = append(slice, item)
	}
	return slice
}

func GetInvoiceStatusFEConf() []map[string]interface{} {

	return GetInvoiceFEConf(INVOICE_STATUS)
}

func GetInvoiceTypeFEConf() []map[string]interface{} {

	return GetInvoiceFEConf(INVOICE_TYPE)
}
