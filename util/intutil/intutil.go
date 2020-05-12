package intutil

import "fmt"

func InArray(num int, slice []int) bool {
	for _, v := range slice {
		if num == v {
			return true
		}
	}
	return false
}

func Intersect(sliceA []int, sliceB []int) []int {
	var result []int

	for _, a := range sliceA {
		for _, b := range sliceB {
			if a == b {
				result = append(result, a)
			}
		}
	}

	return result
}

func ToInt(i interface{}) int {
	switch i.(type) {
	case int:
		return i.(int)
	case int8:
		return int(i.(int8))
	case int16:
		return int(i.(int16))
	case int32:
		return int(i.(int32))
	case int64:
		return int(i.(int64))
	case float32:
		return int(i.(float32))
	case float64:
		return int(i.(float64))
	}
	panic(fmt.Sprintf("存在无法转换的数据类型: %v", i))
}
