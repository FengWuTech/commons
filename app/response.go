package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/FengWuTech/commons/gredis"
	"github.com/FengWuTech/commons/logger"
	"github.com/FengWuTech/commons/setting"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Time string      `json:"time"`
}

func SaveRawRequest(c *gin.Context) {
	logger.Info("SaveRawRequest begin")
	requestRawData, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestRawData))
	c.Set("http_request_raw_data", string(requestRawData))
	logger.Info("SaveRawRequest end")
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	msg := GetMsg(errCode)
	g.ResponseWithMsgAndCode(httpCode, errCode, msg, data)
}

func (g *Gin) ResponseSuccess(data interface{}) {
	g.Response(http.StatusOK, SUCCESS, data)
}

func (g *Gin) ResponseError(errCode int, data interface{}) {
	g.Response(http.StatusOK, errCode, data)
}

func (g *Gin) ResponseErrorWithMsg(errCode int, msg string) {
	g.ResponseWithMsgAndCode(http.StatusOK, errCode, msg, nil)
}

func (g *Gin) ResponseWithMsgAndCode(httpCode, errCode int, msg string, data interface{}) {
	var response = Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
		Time: time.Now().Format("2006-01-02 15:04:05"),
	}
	var resStr, _ = json.Marshal(response)
	var reqBody, _ = g.C.Get("http_request_raw_data")
	var reqHeader, _ = json.Marshal(g.C.Request.Header)
	logger.Infof("reqURL[%v] reqHeader[%v] reqRawData[%v] resBody[%v]", g.C.Request.RequestURI, string(reqHeader), reqBody, string(resStr))
	g.C.JSON(httpCode, response)
	return
}

func (g *Gin) buildCacheKey() string {
	companyID := GetCompanyID(g.C)
	host := g.C.Request.Host
	uri := g.C.Request.RequestURI
	key := fmt.Sprintf("%d-%s-%s", companyID, host, uri)
	return key
}

func (g *Gin) ResponseFromCache() bool {
	key := g.buildCacheKey()
	resData, err := gredis.Get(key)
	if err != nil {
		return false
	}
	g.C.Writer.Write(resData)
	return true
}

func (g *Gin) ResponseAndCache(code int, data interface{}, timeout int) bool {
	key := g.buildCacheKey()

	var s = Response{
		Code: code,
		Msg:  GetMsg(code),
		Data: data,
		Time: time.Now().Format("2006-01-02 15:04:05"),
	}
	rawData, _ := json.Marshal(s)
	//debug模式统一10秒过期
	if setting.ServerSetting.RunMode == "debug" {
		timeout = 10
	}
	err := gredis.Set(key, string(rawData), timeout)
	if err != nil {
		return false
	}

	g.C.Writer.Write(rawData)
	return true
}
