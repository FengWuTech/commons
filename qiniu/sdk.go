package qiniu

import (
	"bytes"
	"context"
	"fmt"
	"github.com/FengWuTech/commons/setting"
	"github.com/parnurzeal/gorequest"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	uuid "github.com/satori/go.uuid"
	"time"
)

func GetUploadToken() string {
	bucket := setting.AppSetting.QBoxBucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	accessKey := setting.AppSetting.QBoxAccess
	secretKey := setting.AppSetting.QBoxSecret
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	return upToken
}

func UploadHttpFile(url string, dstFileName string) string {
	var timenow = time.Now()
	dstFileName = fmt.Sprintf("%04d%02d%02d/%s/%s", timenow.Year(), timenow.Month(), timenow.Day(), uuid.NewV4().String(), dstFileName)
	bucket := setting.AppSetting.QBoxBucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	accessKey := setting.AppSetting.QBoxAccess
	secretKey := setting.AppSetting.QBoxSecret
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone: &storage.ZoneHuabei,
	}
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	_, content, _ := gorequest.New().Get(url).End()
	data := []byte(content)
	dataLen := int64(len(data))
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	err := formUploader.Put(context.Background(), &ret, upToken, dstFileName, bytes.NewReader(data), dataLen, &putExtra)

	if err != nil {
		fmt.Printf("%v\n", err)
		return ""
	}
	return setting.AppSetting.CdnUrl + "/" + dstFileName
}
