package timeutil

import (
	"fmt"
	"testing"
	"time"
)

func TestExampleFormat(t *testing.T) {
	tt := time.Unix(1340244776, 0)
	utc, _ := time.LoadLocation("UTC")
	tt = tt.In(utc)
	fmt.Println(Format("%Y-%m-%d %H:%M:%S", tt))
	// Output:
	// 2012-06-21 02:12:56
}

func TestParseExcelDate(t *testing.T) {
	date := time.Date(2019, 12, 1, 0, 0, 0, 0, time.Local)

	dateNumStr := "43800"
	parsedNumDate := ParseExcelDate(dateNumStr)
	if date != parsedNumDate {
		t.Errorf("Parse %v to %v != %v", dateNumStr, parsedNumDate, date)
	}

	dateDashStr := "2019-12-01"
	parsedDashDate := ParseExcelDate(dateDashStr)
	if date != parsedDashDate {
		t.Errorf("Parse %v to %v != %v", dateDashStr, parsedDashDate, date)
	}

	dateSlashStr := "2019/12/1"
	parsedSlashDate := ParseExcelDate(dateSlashStr)
	if date != parsedSlashDate {
		t.Errorf("Parse %v to %v != %v", dateSlashStr, parsedSlashDate, date)
	}

	dateSpaceStr := "2019-12-01 "
	parsedSpaceDate := ParseExcelDate(dateSpaceStr)
	if !parsedSpaceDate.IsZero() {
		t.Errorf("Parse %v to %v != Zero Date", dateSpaceStr, parsedSpaceDate)
	}
}
