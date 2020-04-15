// Package timeutil 提供了一些列时间、日期格式化输出方法，是基于
// https://github.com/jehiah/go-strftime 类库的封装
package timeutil

import (
	"strings"
	"time"

	"github.com/noaway/dateparse"
	"github.com/unknwon/com"
)

// taken from time/format.go
var conversion = map[string]string{
	/*stdLongMonth      */ "B": "January",
	/*stdMonth          */ "b": "Jan",
	// stdNumMonth       */ "m": "1",
	/*stdZeroMonth      */ "m": "01",
	/*stdLongWeekDay    */ "A": "Monday",
	/*stdWeekDay        */ "a": "Mon",
	// stdDay            */ "d": "2",
	// stdUnderDay       */ "d": "_2",
	/*stdZeroDay        */ "d": "02",
	/*stdHour           */ "H": "15",
	// stdHour12         */ "I": "3",
	/*stdZeroHour12     */ "I": "03",
	// stdMinute         */ "M": "4",
	/*stdZeroMinute     */ "M": "04",
	// stdSecond         */ "S": "5",
	/*stdZeroSecond     */ "S": "05",
	/*stdLongYear       */ "Y": "2006",
	/*stdYear           */ "y": "06",
	/*stdPM             */ "p": "PM",
	// stdpm             */ "p": "pm",
	/*stdTZ             */ "Z": "MST",
	// stdISO8601TZ      */ "z": "Z0700",  // prints Z for UTC
	// stdISO8601ColonTZ */ "z": "Z07:00", // prints Z for UTC
	/*stdNumTZ          */ "z": "-0700", // always numeric
	// stdNumShortTZ     */ "b": "-07",    // always numeric
	// stdNumColonTZ     */ "b": "-07:00", // always numeric
}

// Format 格式化时间
// This is an alternative to time.Format because no one knows
// what date 040305 is supposed to create when used as a 'layout' string
// this takes standard strftime format options. For a complete list
// of format options see http://strftime.org/
func Format(format string, t time.Time) string {
	formatChunks := strings.Split(format, "%")
	var layout []string
	for _, chunk := range formatChunks {
		if len(chunk) == 0 {
			continue
		}
		if layoutCmd, ok := conversion[chunk[0:1]]; ok {
			layout = append(layout, layoutCmd)
			if len(chunk) > 1 {
				layout = append(layout, chunk[1:])
			}
		} else {
			layout = append(layout, "%", chunk)
		}
	}
	return t.Format(strings.Join(layout, ""))
}

//获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateZeroTimeOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTimeOfDay(d)
}

//获取传入的时间所在月份的最后一天，即某月最后一天的23:59点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateEndTimeOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1).AddDate(0, 1, -1)
	return GetEndTimeOfDay(d)
}

//-----------------------------------
//把string转为time

func Str2Time(st string) (time.Time, error) {
	return dateparse.ParseLocal(st)
}

//-----------------------------------

//获取某一天的0点时间
func GetZeroTimeOfDay(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

//获取某一天的最后的时间
func GetEndTimeOfDay(d time.Time) time.Time {
	nowStr := d.Format("2006-01-02")
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", nowStr+" 23:59:59", time.Local)
	return endTime
}

//获取今天同时同分同秒的时间
func GetSameTimeOfToday(d time.Time) time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), d.Hour(), d.Minute(), d.Second(), d.Nanosecond(), d.Location())
}

func GetLastDay() (year int, month int, day int) {
	nowTime := time.Now()
	lastTime := nowTime.AddDate(0, 0, -1)
	return lastTime.Year(), int(lastTime.Month()), lastTime.Day()
}

func GetLastMonth() (year int, month int) {
	nowTime := time.Now()
	lastTime := nowTime.AddDate(0, -1, 0)
	return lastTime.Year(), int(lastTime.Month())
}

func GetLastYear() (year int) {
	nowTime := time.Now()
	lastTime := nowTime.AddDate(-1, 0, 0)
	return lastTime.Year()
}

func GetYearStartTime(year int) time.Time {
	return time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)
}

func GetYearEndTime(year int) time.Time {
	return GetYearStartTime(year).AddDate(1, 0, 0)
}

func GetMonthStartTime(year int, month int) time.Time {
	return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
}

func GetMonthEndTime(year int, month int) time.Time {
	return GetMonthStartTime(year, month).AddDate(0, 1, 0)
}

func GetDayStartTime(year int, month int, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

func GetDayEndTime(year int, month int, day int) time.Time {
	return GetDayStartTime(year, month, day).AddDate(0, 0, 1)
}

//-----------------------------------

func ParseExcelDate(i string) time.Time {
	if i == "" {
		return time.Time{}
	}
	if strings.Contains(i, "-") || strings.Contains(i, "/") {
		t, _ := dateparse.ParseLocal(i)
		return t
	}
	date, _ := time.ParseInLocation("2006-01-02", "1900-01-01", time.Local)
	days := com.StrTo(i).MustInt()
	date = date.AddDate(0, 0, days-2)
	return date
}

func GetLastDayWithoutOver(year int, month int) int {
	now := time.Now()
	monthStartTime := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	nextMonthStartTime := monthStartTime.AddDate(0, 1, 0)

	statisticsTime := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)

	if statisticsTime.Unix() < monthStartTime.Unix() { //当前月之前的时间
		return statisticsTime.AddDate(0, 1, -1).Day()
	} else if statisticsTime.Unix() >= nextMonthStartTime.Unix() { //当前月之后的时间
		return 0
	} else {
		return now.Day()
	}
}

func GetLastMonthWithoutOver(year int) int {
	if year < time.Now().Year() {
		return 12
	} else if year > time.Now().Year() {
		return 0
	} else {
		return int(time.Now().Month())
	}
}
