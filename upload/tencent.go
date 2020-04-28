package upload

import (
	"context"
	"fmt"
	"github.com/FengWuTech/commons/setting"
	uuid "github.com/satori/go.uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Cos struct {
	AppID     string
	CosBucket string
	CosRegion string
	SecretID  string
	SecretKey string
}

func NewClient(secretID string, secretKey string, appID string, bucket string, region string) *Cos {
	return &Cos{
		AppID:     appID,
		CosBucket: bucket,
		CosRegion: region,
		SecretID:  secretID,
		SecretKey: secretKey,
	}
}

func (cs *Cos) FilePath(suffix string) string {
	timeNow := time.Now()
	if suffix == "" {
		return fmt.Sprintf("pubsrv/%d%d%d/%s", timeNow.Year(), timeNow.Month(), timeNow.Day(), uuid.NewV1().String())
	} else {
		return fmt.Sprintf("pubsrv/%d%d%d/%s.%s", timeNow.Year(), timeNow.Month(), timeNow.Day(), uuid.NewV1().String(), suffix)
	}
}

type UploadToken struct {
	Token       sts.Credentials `json:"token"`
	FileKey     string          `json:"file_key"`
	StartTime   int64           `json:"start_time"`
	ExpiredTime int64           `json:"expired_time"`
	Bucket      string          `json:"bucket"`
	Region      string          `json:"region"`
}

func (cs *Cos) GetUploadToken(suffix string) *UploadToken {
	appid := cs.AppID
	bucket := cs.CosBucket
	region := cs.CosRegion
	cli := sts.NewClient(cs.SecretID, cs.SecretKey, nil)
	opt := &sts.CredentialOptions{
		DurationSeconds: int64(time.Hour.Seconds()),
		Region:          region,
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					Action: []string{
						"name/cos:PostObject",
						"name/cos:PutObject",
					},
					Effect: "allow",
					Resource: []string{
						"qcs::cos:" + bucket + ":uid/" + appid + ":" + bucket + "/pubsrv/*",
					},
				},
			},
		},
	}
	res, err := cli.GetCredential(opt)

	if err != nil {
		fmt.Printf("get upload token err: %v", err)
		return nil
	}

	path := cs.FilePath(suffix)

	return &UploadToken{
		Token:       *res.Credentials,
		FileKey:     path,
		StartTime:   time.Now().Unix(),
		ExpiredTime: time.Now().Unix() + 3600,
		Bucket:      bucket,
		Region:      region,
	}
}

func (cs *Cos) UploadBytes(content string) string {
	//u, _ := url.Parse("https://" + setting.TencentCloudSetting.CosBucket)
	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", setting.TencentCloudSetting.CosBucket, setting.TencentCloudSetting.CosRegion))
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  cs.SecretID,
			SecretKey: cs.SecretKey,
		},
	})
	name := cs.FilePath("")
	f := strings.NewReader(content)

	_, err := c.Object.Put(context.Background(), name, f, nil)
	if err != nil {
		fmt.Printf("upload err: %v", err)
		return ""
	}

	return name
}
