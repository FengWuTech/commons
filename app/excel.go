package app

import (
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/FengWuTech/commons/logger"
	"github.com/gin-gonic/gin"
	"regexp"
	"strconv"
	"strings"
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

func (g *Gin) ExcelResponse(filename string, data *ExcelData) {
	g.C.Header("Content-Type", "application/octet-stream")
	g.C.Header("Content-Disposition", "attachment; filename="+filename)
	g.C.Header("Content-Transfer-Encoding", "binary")

	xlsx := excelize.NewFile()
	sheetName := "Sheet1"
	sheetId := xlsx.NewSheet(sheetName)
	xlsx.SetActiveSheet(sheetId)

	headerAxis := "A1"
	xlsx.SetSheetRow(sheetName, headerAxis, &data.Header)
	lockAndRedStyle, _ := xlsx.NewStyle(`{"font":{"color":"#FF0000"}, "protection":{"locked":true}}`)
	unlockStyle, _ := xlsx.NewStyle(`{"protection":{"locked": false}}`)

	for index, header := range data.Header {
		rowName := "1"
		colName, _ := excelize.ColumnNumberToName(index + 1)
		// 默认设置不保护
		xlsx.SetColStyle(sheetName, colName, unlockStyle)
		// 保护表头带星号单元格
		if strings.Contains(header, "*") {
			xlsx.SetCellStyle(sheetName, colName+rowName, colName+rowName, lockAndRedStyle)
		}
	}

	re, _ := regexp.Compile(`\[(.*?)\]`)
	// 因为选项不支持跨sheet，所以放在同一个sheet中，在有内容区域往右放置选项
	optionOffset := len(data.Header) + 10

	for rowNum, _ := range data.Content {
		rowName := strconv.Itoa(rowNum + 2)
		obj := data.Content[rowNum]
		for index, field := range obj {
			colName, _ := excelize.ColumnNumberToName(index + 1)
			// 提供下拉样式
			if re.MatchString(field) {
				matched := re.FindStringSubmatch(field)
				if matched != nil {
					optionColName, _ := excelize.ColumnNumberToName(index + optionOffset)
					// 第二行开始放置选项，只用处理一次，把同一列都使用选项约束
					if v, _ := xlsx.GetCellValue(sheetName, optionColName + "2"); v == "" {
						elements := strings.Split(matched[1], ",")
						for optionIndex, element := range elements {
							optionRowName := strconv.Itoa(optionIndex + 2)
							xlsx.SetSheetRow(sheetName, optionColName+optionRowName, &[]interface{}{element,})
						}
						// https://xuri.me/excelize/zh-hans/data.html#AddDataValidation
						dvRange := excelize.NewDataValidation(true)
						// 需要增加选项的列
						dvRange.Sqref = fmt.Sprintf("%s%s:%s%s", colName, "2", colName, strconv.Itoa(len(data.Content) + 2))
						// 选项来源的列
						dvRange.SetSqrefDropList(fmt.Sprintf("$%s$%s:$%s$%s", optionColName, "2", optionColName, strconv.Itoa(len(elements) + 2)), true)
						xlsx.AddDataValidation(sheetName, dvRange)
					}
					obj[index] = ""
				}
			}
		}
		xlsx.SetSheetRow(sheetName, "A"+rowName, &obj)
	}

	// 保护工作簿
	// xlsx.ProtectSheet(sheetName, nil)

	_ = xlsx.Write(g.C.Writer)
}


type HeaderItem struct {
	Name  string
	Width int
}
type PayFlowRow []string
type BillRow []string
type BillRowList []BillRow
type PayFlowExcelRow struct {
	PayFlowRow  PayFlowRow
	BillRowList BillRowList
}

type PayFlowExcelData struct {
	Header  []HeaderItem
	Content []PayFlowExcelRow
}

func (g *Gin) PayFlowExcelResponse(filename string, data *PayFlowExcelData) {
	g.C.Header("Content-Type", "application/octet-stream")
	g.C.Header("Content-Disposition", "attachment; filename="+filename)
	g.C.Header("Content-Transfer-Encoding", "binary")

	xlsx := excelize.NewFile()
	sheetName := "Sheet1"
	sheetId := xlsx.NewSheet(sheetName)
	xlsx.SetActiveSheet(sheetId)

	//headerStyle, _ := xlsx.NewStyle(`{"horizontal":"center", "vertical":"center"}`)
	contentStyle, _ := xlsx.NewStyle(`{"alignment":{"horizontal":"left", "vertical":"center"}}`)

	headerAxis := "A1"
	_ = xlsx.SetSheetRow(sheetName, headerAxis, &data.Header)

	xlsx.SetRowHeight(sheetName, 1, 24)
	for index, item := range data.Header {
		rowName := "1"
		colName, _ := excelize.ColumnNumberToName(index + 1)
		xlsx.SetCellValue(sheetName, colName+rowName, item.Name)
		xlsx.SetColWidth(sheetName, colName, colName, float64(item.Width))
		xlsx.SetColStyle(sheetName, colName, contentStyle)
	}

	mergeStartRow := 2
	mergeEndRow := mergeStartRow
	for _, rowData := range data.Content {
		payFlowRow := rowData.PayFlowRow
		billRowList := rowData.BillRowList

		billRowHeight := len(billRowList)
		if billRowHeight <= 0 {
			billRowHeight = 1
		}
		mergeEndRow += billRowHeight - 1

		for i := 1; i <= len(rowData.PayFlowRow); i++ {
			column, _ := excelize.ColumnNumberToName(i)
			mergeHCell := fmt.Sprintf("%s%d", column, mergeStartRow)
			mergeVCell := fmt.Sprintf("%s%d", column, mergeEndRow)
			_ = xlsx.MergeCell(sheetName, mergeHCell, mergeVCell)
		}

		for idx, value := range payFlowRow {
			column, _ := excelize.ColumnNumberToName(idx + 1)
			axis := fmt.Sprintf("%s%d", column, mergeStartRow)
			_ = xlsx.SetCellValue(sheetName, axis, value)
		}

		payFlowRowWidth := len(payFlowRow)
		for rowIdx, billRow := range billRowList {
			for widthIdx, value := range billRow {
				column, _ := excelize.ColumnNumberToName(payFlowRowWidth + widthIdx + 1)
				axis := fmt.Sprintf("%s%d", column, mergeStartRow+rowIdx)
				_ = xlsx.SetCellValue(sheetName, axis, value)
			}
		}

		for i := 0; i < billRowHeight; i++ {
			xlsx.SetRowHeight(sheetName, mergeStartRow+i, 24)
		}

		mergeStartRow += billRowHeight
		mergeEndRow = mergeStartRow
	}

	_ = xlsx.Write(g.C.Writer)
}