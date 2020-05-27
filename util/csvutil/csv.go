package csvutil

import (
	"github.com/FengWuTech/commons/util/encodeutil"
	"strings"

	"github.com/axgle/mahonia"
)

func GBKToUtf8(src string) string {
	decoder := mahonia.NewDecoder("GBK")
	return decoder.ConvertString(src)
}

func ParseCsvToMap2(content string, ignoreNum int, fromCharset string) []map[string]string {
	if content == "" {
		return nil
	}

	var titleItems []string
	var dataList = make([]map[string]string, 0)

	decoder := mahonia.NewDecoder(fromCharset)
	content = decoder.ConvertString(content)

	rows := strings.Split(content, "\n")
	titleItems = strings.Split(rows[ignoreNum], ",")

	for rowID := ignoreNum + 1; rowID < len(rows); rowID++ {
		if rows[rowID] == "" {
			continue
		}
		rowItems := strings.Split(rows[rowID], ",")
		dataItem := make(map[string]string)
		for k, v := range rowItems {
			if k >= len(titleItems) {
				continue
			}
			key := strings.TrimSpace(titleItems[k])
			value := strings.TrimSpace(v)
			value = strings.TrimLeft(value, "'")
			dataItem[key] = value
		}
		dataList = append(dataList, dataItem)
	}

	return dataList
}

func ParseCSVToMap(content string, ignoreNum int) []map[string]string {
	if content == "" {
		return nil
	}

	var titleItems []string
	var dataList = make([]map[string]string, 0)

	if !encodeutil.IsUTF8([]byte(content)) {
		content = GBKToUtf8(content)
	}

	rows := strings.Split(content, "\n")
	titleItems = strings.Split(rows[ignoreNum], ",")

	for rowID := ignoreNum + 1; rowID < len(rows); rowID++ {
		if rows[rowID] == "" {
			continue
		}
		rowItems := strings.Split(rows[rowID], ",")
		dataItem := make(map[string]string)
		for k, v := range rowItems {
			if k >= len(titleItems) {
				continue
			}
			key := strings.TrimSpace(titleItems[k])
			value := strings.TrimSpace(v)
			value = strings.TrimLeft(value, "'")
			dataItem[key] = value
		}
		dataList = append(dataList, dataItem)
	}

	return dataList
}
