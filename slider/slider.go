package slider

import (
	"github.com/FengWuTech/commons/logger"
	"github.com/FengWuTech/commons/setting"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/afs"
)

var (
	REGION_ID     = setting.AliyunSetting.RegionID
	ACCESS_ID     = setting.AliyunSetting.AccessID
	ACCESS_SECRET = setting.AliyunSetting.AccessSecret
	APPKEY        = setting.AliyunSetting.SliderAppKey
)

type SliderReq struct {
	Sig       string `json:"sig" valid:"Required"`
	SessionId string `json:"session_id" valid:"Required"`
	Token     string `json:"token" valid:"Required"`
	Scene     string `json:"scene" valid:"Required"`
}

func CheckAliyunUserSlider(form SliderReq, clientIp string) bool {
	client, err := afs.NewClientWithAccessKey(REGION_ID, ACCESS_ID, ACCESS_SECRET)
	if err != nil {
		return false
	}
	request := afs.CreateAuthenticateSigRequest()
	request.Domain = "afs.aliyuncs.com"
	request.Sig = form.Sig
	request.SessionId = form.SessionId
	request.Token = form.Token
	request.Scene = form.Scene
	request.AppKey = APPKEY
	// TODO use real ip
	request.RemoteIp = clientIp
	response, err := client.AuthenticateSig(request)
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	logger.Info("阿里云验证返回: %v", response)
	return response.Code == 100
}