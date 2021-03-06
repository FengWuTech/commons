package util

import (
	"encoding/json"
	"reflect"
	"strconv"

	jsoniter "github.com/json-iterator/go"

	"github.com/siddontang/go/bson"
)

func Interface2String(inter interface{}, param ...int) string {
	switch inter.(type) {
	case string:
		return inter.(string)
	case int, int32:
		return strconv.FormatInt(int64(inter.(int)), 10)
	case int64:
		return strconv.FormatInt(inter.(int64), 10)
	case float32:
		return strconv.FormatFloat(float64(inter.(float32)), 'f', param[0], 32)
	case float64:
		perc := 2
		if param != nil {
			perc = param[0]
		}
		return strconv.FormatFloat(inter.(float64), 'f', perc, 64)
	}
	return ""
}

//判断interface是否为nil
func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

func StructToMapUseJson(st interface{}) map[string]interface{} {
	var ret map[string]interface{}
	jsonStr, _ := json.Marshal(st)
	json.Unmarshal(jsonStr, &ret)
	return ret
}

//struct转成map
func StructToMap(st interface{}) map[string]interface{} {
	types := reflect.TypeOf(st)
	values := reflect.ValueOf(st)

	var data = make(map[string]interface{})
	for i := 0; i < types.NumField(); i++ {
		k := types.Field(i).Name
		var v = values.Field(i).Interface()
		if !IsNil(v) {
			switch v.(type) {
			case int:
				data[k] = v
			case *string:
				vv := v.(*string)
				data[k] = *vv
			case string:
				data[k] = v
			}
		}
	}

	return data
}

func DeepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = DeepCopy(v)
		}

		return newSlice
	} else if valueMap, ok := value.(bson.M); ok {
		newMap := make(bson.M)
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}
	}
	return value
}

//严格的struct拷贝
func StrictStructCopy(src interface{}, dst interface{}) {
	sval := reflect.ValueOf(src).Elem()
	dval := reflect.ValueOf(dst).Elem()

	for i := 0; i < sval.NumField(); i++ {
		value := sval.Field(i)
		name := sval.Type().Field(i).Name

		dvalue := dval.FieldByName(name)
		if dvalue.IsValid() == false {
			continue
		}
		dvalue.Set(value)
	}
}

//利用json做中间值进行拷贝
func StructCopyUseJson(src interface{}, dst interface{}) {
	//srcJson, _ := json.Marshal(src)
	//_ = json.Unmarshal(srcJson, dst)

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonBytes, _ := json.Marshal(src)
	_ = json.Unmarshal(jsonBytes, dst)
}
