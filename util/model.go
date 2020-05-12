package util

type KV map[string]interface{}
type List []interface{}
type Where map[string]interface{}
type PageData struct {
	Total int         `json:"total"`
	List  interface{} `json:"list"`
}

var SelectAll []string
