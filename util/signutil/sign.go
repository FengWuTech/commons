package signutil

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

type Sign struct {
	URL    string
	ApiKey string
	KV     map[string]string
	Body   string
}

func NewSign(apiKey string) *Sign {
	var sign Sign
	sign.KV = make(map[string]string)
	sign.ApiKey = apiKey
	return &sign
}

func (sign *Sign) AddQuery(key string, value interface{}) {
	sign.KV[key] = valueToString(value)
}

func (sign *Sign) SetBody(body string) {
	sign.Body = body
}

func (sign *Sign) GenSign() string {
	var keys []string
	for k := range sign.KV {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var query = make([]string, 0)
	for _, k := range keys {
		query = append(query, fmt.Sprintf("%s=%s", k, valueToString(sign.KV[k])))
	}

	sumStr := strings.Join(query, "&") + sign.Body + sign.ApiKey
	signStr := fmt.Sprintf("%x", md5.Sum([]byte(sumStr)))

	return signStr
}

func (sign *Sign) GenSignURL(host string) string {
	signStr := sign.GenSign()
	query := make([]string, 0)
	for k, v := range sign.KV {
		query = append(query, k+"="+v)
	}
	query = append(query, "sign="+signStr)
	return host + "?" + strings.Join(query, "&")
}

func (sign *Sign) VerifySign(requestURI string, body string) bool {
	signStr := ""
	u, _ := url.ParseRequestURI(requestURI)
	for k, v := range u.Query() {
		if k == "sign" {
			signStr = v[0]
		} else {
			sign.AddQuery(k, v[0])
		}
	}
	sign.SetBody(body)
	if signStr == sign.GenSign() {
		return true
	} else {
		return false
	}
}

func valueToString(inter interface{}) string {
	switch inter.(type) {
	case string:
		return fmt.Sprint(inter.(string))
		break
	case int:
		return fmt.Sprint(inter.(int))
		break
	case int64:
		return fmt.Sprint(inter.(int64))
		break
	case float64:
		return fmt.Sprint(inter.(float64))
		break
	}
	panic("valueToString 存在无法解析的类型")
}
