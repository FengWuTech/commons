package slider

import (
	"fmt"
	"github.com/FengWuTech/commons/logger"
	"github.com/FengWuTech/commons/setting"
	captcha "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/captcha/v20190722"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"

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

type CaptchaReq struct {
	Ticket  string `json:"ticket" valid:"Required"`
	Randstr string `json:"randstr" valid:"Required"`
}

func CheckTencentCaptcha(req CaptchaReq, userIp string) bool {
	credential := common.NewCredential(
		setting.TencentCloudSetting.SecretId,
		setting.TencentCloudSetting.SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "captcha.tencentcloudapi.com"
	client, _ := captcha.NewClient(credential, regions.Beijing, cpf)

	request := captcha.NewDescribeCaptchaResultRequest()
	request.CaptchaType = common.Uint64Ptr(9)
	request.Ticket = common.StringPtr(req.Ticket)
	request.UserIp = common.StringPtr(userIp)
	request.Randstr = common.StringPtr(req.Randstr)
	request.CaptchaAppId = common.Uint64Ptr(setting.TencentCloudSetting.CaptchaAppId)
	request.AppSecretKey = common.StringPtr(setting.TencentCloudSetting.CaptchaAppSecretKey)
	response, err := client.DescribeCaptchaResult(request)
	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return false
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		return false
	}
	fmt.Printf("%s", response.ToJsonString())
	return *response.Response.CaptchaCode == 1
}