package constant

import (
	"bytes"
	"html/template"
	"sort"
)

const (
	BILL_BIND_OBJECT_TYPE = 1 //账单绑定的对象类型（二级）
	BILL_PAY_STATUS       = 2 //账单支付状态

	BILL_BIND_OBJECT_TYPE_ROOM       = 1 //房屋
	BILL_BIND_OBJECT_TYPE_PARKING    = 2 //车位
	BILL_BIND_OBJECT_TYPE_USER       = 3 //住户
	BILL_BIND_OBJECT_TYPE_ROOM_METER = 4 //户表
	//BILL_BIND_OBJECT_TYPE_PUBLIC_METER = 5 //公表（暂时不用）
	BILL_BIND_OBJECT_TYPE_OTHER = 6 //其他

	//账单支付状态
	BILL_PAY_STATUS_UNPAID   = 0 //未支付
	BILL_PAY_STATUS_PAID     = 1 //已支付
	BILL_PAY_STATUS_REFUNDED = 2 //已退款
)

var BILL_INFO_MAP = map[int]map[int]map[string]interface{}{

	BILL_BIND_OBJECT_TYPE: {
		BILL_BIND_OBJECT_TYPE_ROOM: {
			"value":             "房屋",
			"chargeBindNameTPL": "{{.GroupName}}/{{.BuildingName}}/{{.UnitName}}/{{.RoomName}}",
		},
		BILL_BIND_OBJECT_TYPE_PARKING: {
			"value":             "车位",
			"chargeBindNameTPL": "{{.GroupName}}/{{.ParkingName}}",
		},
		BILL_BIND_OBJECT_TYPE_USER: {
			"value":             "住户",
			"chargeBindNameTPL": "{{.UserName}}({{.UserMobile}})",
		},
		BILL_BIND_OBJECT_TYPE_ROOM_METER: {
			"value":             "户表",
			"chargeBindNameTPL": "{{.GroupName}}/{{.BuildingName}}/{{.UnitName}}/{{.RoomName}}-{{.MeterName}}",
		},
		BILL_BIND_OBJECT_TYPE_OTHER: {
			"value": "其他",
		},
	},
}

type ChargeBindRoom struct {
	GroupName    string
	BuildingName string
	UnitName     string
	RoomName     string
}

type ChargeBindParking struct {
	GroupName   string
	ParkingName string
}

type ChargeBindUser struct {
	UserName   string
	UserMobile string
}

type ChargeBindMeterRoom struct {
	GroupName    string
	BuildingName string
	UnitName     string
	RoomName     string
	MeterName    string
}

//获取前端的展示配置
func GetBillFEConf(which int) []map[string]interface{} {

	slice := make([]map[string]interface{}, 0)
	// To store the keys in slice in sorted order
	var keys []int
	for k := range BILL_INFO_MAP[which] {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// To perform the opertion you want
	for _, k := range keys {
		v := BILL_INFO_MAP[which][k]

		item := map[string]interface{}{
			"value": k,
			"label": v["value"],
		}
		slice = append(slice, item)
	}
	return slice
}

func GetBillBindObjectTypeFEConf() []map[string]interface{} {

	return GetBillFEConf(BILL_BIND_OBJECT_TYPE)
}

//获取支付绑定对象的Name
func GenChargeBindName(tplType int, data interface{}) (string, error) {

	tplString := BILL_INFO_MAP[BILL_BIND_OBJECT_TYPE][tplType]["chargeBindNameTPL"].(string)

	tpl := template.New("name")
	tpl, _ = tpl.Parse(tplString)

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
