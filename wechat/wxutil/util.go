package wxutil

import (
	"github.com/FengWuTech/commons/setting"
	"github.com/FengWuTech/commons/upload"
	"github.com/parnurzeal/gorequest"
)

func UploadFromURL(url string) string {
	_, content, _ := gorequest.New().Get(url).End()
	cos := upload.NewClient(setting.TencentCloudSetting.SecretId, setting.TencentCloudSetting.SecretKey,
		setting.TencentCloudSetting.AppId, setting.TencentCloudSetting.CosBucket, setting.TencentCloudSetting.CosRegion)
	qrcodeURL := setting.AppSetting.CdnUrl + "/" + cos.UploadBytes(content)
	return qrcodeURL
}
