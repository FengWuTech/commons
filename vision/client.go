package vision

import "github.com/FengWuTech/commons/setting"

type Client struct {
	AppID     string
	AppSecret string
	BaseURL   string
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Time string      `json:"time"`
}

func NewClient() *Client {
	return &Client{
		AppID:     setting.VisionSetting.AppID,     //直接赋值好了
		AppSecret: setting.VisionSetting.AppSecret, //直接赋值好了
		BaseURL:   setting.VisionSetting.BaseURL,
	}
}
