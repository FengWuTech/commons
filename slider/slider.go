package slider

import (
	"github.com/FengWuTech/commons/logger"
	"github.com/FengWuTech/commons/setting"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/afs"
)

type SliderReq struct {
	Sig       string `json:"sig" valid:"Required"`
	SessionId string `json:"session_id" valid:"Required"`
	Token     string `json:"token" valid:"Required"`
	Scene     string `json:"scene" valid:"Required"`
}

func CheckAliyunUserSlider(form SliderReq, clientIp string) bool {
	var (
		regionId     = setting.AliyunSetting.RegionID
		accessId     = setting.AliyunSetting.AccessID
		accessSecret = setting.AliyunSetting.AccessSecret
		appkey       = setting.AliyunSetting.SliderAppKey
	)
	client, err := afs.NewClientWithAccessKey(regionId, accessId, accessSecret)
	if err != nil {
		return false
	}
	request := afs.CreateAuthenticateSigRequest()
	request.Domain = "afs.aliyuncs.com"
	request.Sig = form.Sig
	request.SessionId = form.SessionId
	request.Token = form.Token
	request.Scene = form.Scene
	request.AppKey = appkey
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