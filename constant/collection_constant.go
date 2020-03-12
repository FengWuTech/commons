package constant

import "sort"

//银行托收模板

const (
	COLLECTION_FILE_TYPE           = 1 //文件的类型
	COLLECTION_COLUMN_SEND_TYPE    = 2 //送盘的字段类型
	COLLECTION_COLUMN_RECEIVE_TYPE = 3 //回盘的字段类型

	COLLECTION_FILE_TYPE_EXCEL = 1 //文件的类型(excel)
	COLLECTION_FILE_TYPE_TXT   = 2 //文件的类型(txt)

	COLLECTION_COLUMN_SEND_TYPE_SERIAL_NUMBER = 1 //序号
	COLLECTION_COLUMN_SEND_TYPE_NAME          = 2 //住户姓名
	COLLECTION_COLUMN_SEND_TYPE_BANK_ACCOUNT  = 3 //银行账户
	COLLECTION_COLUMN_SEND_TYPE_FEE           = 4 //金额
	COLLECTION_COLUMN_SEND_TYPE_FEE_NAME      = 5 //费用名称
	COLLECTION_COLUMN_SEND_TYPE_DATE          = 6 //送盘日期

	//todo
	COLLECTION_COLUMN_RECEIVE_TYPE_SERIAL_NUMBER = 1 //序号
	COLLECTION_COLUMN_RECEIVE_TYPE_NAME          = 2 //住户姓名
	COLLECTION_COLUMN_RECEIVE_TYPE_BANK_ACCOUNT  = 3 //银行账户
	COLLECTION_COLUMN_RECEIVE_TYPE_FEE           = 4 //应收金额
	COLLECTION_COLUMN_RECEIVE_TYPE_REAL_FEE      = 5 //实际划扣金额
	COLLECTION_COLUMN_RECEIVE_TYPE_FEE_NAME      = 6 //费用名称
	COLLECTION_COLUMN_RECEIVE_TYPE_DATE          = 7 //回盘日期
	COLLECTION_COLUMN_RECEIVE_TYPE_TAG           = 8 //扣款标志 Y或N

)

var COLLECTION_INFO_MAP = map[int]map[int]map[string]interface{}{

	COLLECTION_FILE_TYPE: {
		COLLECTION_FILE_TYPE_EXCEL: {
			"value": "excel",
		},
		COLLECTION_FILE_TYPE_TXT: {
			"value": "txt",
		},
	},
	COLLECTION_COLUMN_SEND_TYPE: {
		COLLECTION_COLUMN_SEND_TYPE_SERIAL_NUMBER: {
			"value": "序号",
		},
		COLLECTION_COLUMN_SEND_TYPE_NAME: {
			"value": "住户姓名",
		},
		COLLECTION_COLUMN_SEND_TYPE_BANK_ACCOUNT: {
			"value": "银行账户",
		},
		COLLECTION_COLUMN_SEND_TYPE_FEE: {
			"value": "金额",
		},
		COLLECTION_COLUMN_SEND_TYPE_FEE_NAME: {
			"value": "费用名称",
		},
		COLLECTION_COLUMN_SEND_TYPE_DATE: {
			"value": "送盘日期",
		},
	},
	COLLECTION_COLUMN_RECEIVE_TYPE: {
		COLLECTION_COLUMN_RECEIVE_TYPE_SERIAL_NUMBER: {
			"value": "序号",
		},
		COLLECTION_COLUMN_RECEIVE_TYPE_NAME: {
			"value": "住户姓名",
		},
		COLLECTION_COLUMN_RECEIVE_TYPE_BANK_ACCOUNT: {
			"value": "银行账户",
		},
		COLLECTION_COLUMN_RECEIVE_TYPE_FEE: {
			"value": "金额",
		},
		COLLECTION_COLUMN_RECEIVE_TYPE_REAL_FEE: {
			"value": "金额",
		},
		COLLECTION_COLUMN_RECEIVE_TYPE_FEE_NAME: {
			"value": "费用名称",
		},
		COLLECTION_COLUMN_RECEIVE_TYPE_DATE: {
			"value": "回盘日期",
		},
		COLLECTION_COLUMN_RECEIVE_TYPE_TAG: {
			"value": "扣款标志",
		},
	},
}

//获取前端的展示配置
func GetCollectionFEConf(which int) []map[string]interface{} {

	slice := make([]map[string]interface{}, 0)

	// To store the keys in slice in sorted order
	var keys []int
	for k := range COLLECTION_INFO_MAP[which] {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// To perform the opertion you want
	for _, k := range keys {
		v := COLLECTION_INFO_MAP[which][k]

		item := map[string]interface{}{
			"value": k,
			"label": v["value"],
		}
		slice = append(slice, item)
	}
	return slice
}

//获取文件类型
func GetCollectionFileTypeFEConf() []map[string]interface{} {

	return GetCollectionFEConf(COLLECTION_FILE_TYPE)
}

//获取送盘字段列表
func GetCollectionColumnSendTypeFEConf() []map[string]interface{} {

	return GetCollectionFEConf(COLLECTION_COLUMN_SEND_TYPE)
}

//获取回盘字段列表
func GetCollectionColumnReceiveTypeFEConf() []map[string]interface{} {

	return GetCollectionFEConf(COLLECTION_COLUMN_RECEIVE_TYPE)
}
