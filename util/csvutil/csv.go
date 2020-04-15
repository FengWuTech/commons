package csvutil

import (
	"strings"

	"github.com/axgle/mahonia"
)

func GBKToUtf8(src string) string {
	decoder := mahonia.NewDecoder("GBK")
	return decoder.ConvertString(src)
}

func ParseCSVToMap(content string, ignoreNum int) []map[string]string {
	if content == "" {
		return nil
	}

	var titleItems []string
	var dataList = make([]map[string]string, 0)

	rows := strings.Split(content, "\n")
	titleItems = strings.Split(rows[ignoreNum], ",")

	for rowID := ignoreNum + 1; rowID < len(rows); rowID++ {
		rowItems := strings.Split(rows[rowID], ",")
		dataItem := make(map[string]string)
		for k, v := range rowItems {
			key := strings.TrimSpace(titleItems[k])
			key = GBKToUtf8(key)
			value := strings.TrimSpace(v)
			value = strings.TrimLeft(value, "'")
			value = GBKToUtf8(value)
			dataItem[key] = value
		}
		dataList = append(dataList, dataItem)
	}

	return dataList
}
