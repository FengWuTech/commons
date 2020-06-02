package vision

import (
	"github.com/FengWuTech/commons/setting"
	"testing"
)

func New() *Client {
	return NewClient(setting.VisionSetting.BaseURL)
}

//测试警报结果通知
func TestSyncWarningResult(t *testing.T) {

	param := WarningProcessResultOptions{
		CommunityID:   5, //云趣园,
		DeviceID:      "5309ee0e3fb14296b17d962055c07dcf",
		AlarmID:       "h3ai34ye",
		ProcessName:   "张晓华",
		ProcessStatus: 3,
	}

	err, result := New().SyncWarningProcessResult(param)
	if err != nil || result.Code != 200 {
		t.Errorf("request failed, code[%v]msg[%v]", result.Code, result.Msg)
	}

}
