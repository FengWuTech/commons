package app

import (
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/FengWuTech/commons/logger"
	"github.com/gin-gonic/gin"
)

type ExcelRow []string

type ExcelData struct {
	Header []string
	Content []ExcelRow
	headerMap map[string]int
}

func (data *ExcelData) InitHeaderMap()  {
	// 使用表头字段的index获取对应数据字段
	data.headerMap = map[string]int{}
	for index, header := range data.Header {
		data.headerMap[header] = index
	}
}

func (data *ExcelData) GetField(row ExcelRow, column string) string {
	index, ok := data.headerMap[column]
	if ok && index < len(row) {
		return row[index]
	}
	return ""
}

func BindExcel(c *gin.Context, data *ExcelData) error {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		logger.Warnf("Form表单获取Excel出错: %s", err.Error())
		return errors.New("上传文件出错，请选择excel文件")
	}
	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		logger.Warnf("打开上传单Excel文件出错: %s", err.Error())
		return errors.New("文件出错，请提供符合要求的excel文件")
	}
	firstSheetIndex := 1
	sheetName := xlsx.GetSheetName(firstSheetIndex)
	rows, _ := xlsx.GetRows(sheetName)
	if len(rows) <= 0 {
		return errors.New("excel文件无内容，请补充后上传")
	}
	for index, row := range rows {
		var newRow ExcelRow
		for _, colCell := range row {
			// header row
			if index == 0 {
				data.Header = append(data.Header, colCell)
			} else {
				newRow = append(newRow, colCell)
			}
		}
		// data row
		if index != 0 {
			if len(newRow) < len(data.Header) {
				for i :=len(newRow); i < len(data.Header); i++ {
					newRow = append(newRow, "")
				}
			}
			data.Content = append(data.Content, newRow)
		}
	}
	data.InitHeaderMap()
	return nil
}
