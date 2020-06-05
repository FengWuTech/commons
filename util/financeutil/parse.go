package financeutil

import (
	"strconv"
	"strings"
)

func float64ToString(num float64, prec int) string {
	return strconv.FormatFloat(num, 'f', prec, 64)
}

func float32ToString(num float32, prec int) string {
	return strconv.FormatFloat(float64(num), 'f', prec, 32)
}

func RMBFenToYuan(num int64) float64 {
	return float64(num) / 100
}

func RMBFenToYuanString(num int64) string {
	return float64ToString(RMBFenToYuan(num), 2)
}

func ParseFloat64(num string) float64 {
	num = strings.ReplaceAll(num, ",", "")
	num = strings.ReplaceAll(num, "", "")
	flt64, _ := strconv.ParseFloat(num, 64)
	return flt64
}

func ParseFloat32(num string) float32 {
	return float32(ParseFloat64(num))
}
