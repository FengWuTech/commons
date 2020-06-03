package vision

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/astaxie/beego/validation"
	"github.com/parnurzeal/gorequest"
)

type WarningProcessResultOptions struct {
	CommunityID interface{} `json:"community_id"  valid:"Required"` //小区ID(是fpms的group id)
	DeviceID    string      `json:"device_id"  valid:"Required"`    //设备编号
	AlarmID     string      `json:"alarm_id"  valid:"Required"`     //报警id(按照该id去sync处理结果)

	ProcessName   string `json:"process_name"  valid:"Required"`      //处理人姓名
	ProcessStatus int    `json:"process_status"  valid:"Range(1, 4)"` //处理状态 1 未处理 2 系统自动处理 3 人工处理 4 误报
}

//报警处理结果sync给社区
func (cli *Client) SyncWarningProcessResult(param WarningProcessResultOptions) (error, Response) {

	path := "/api/alarm/sync"
	query := cli.genQuery(nil)
	rawURL := cli.BaseURL + path + "?" + query

	// param check
	valid := validation.Validation{}
	b, err := valid.Valid(&param)
	if err != nil {
		return err, Response{}
	}
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		return errors.New("参数错误"), Response{}
	}

	//做一下两个系统的id映射
	err, gMap := GetGroupMap(param.CommunityID.(int))
	if err != nil {
		return err, Response{}
	}
	param.CommunityID = gMap.COMMUNITY

	//发请求
	paramJSON, _ := json.Marshal(param)
	paramStr := string(paramJSON)

	var urlResp Response
	request := gorequest.New()
	_, _, errs := request.Post(rawURL).Send(paramStr).EndStruct(&urlResp)
	if len(errs) > 0 {
		return errs[0], urlResp
	}
	return nil, urlResp
}
