package app

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

const (
	WEB_APP_CODE = "fpms"
	CompanyKey   = "companyID"
	StaffKey     = "staffID"
	UserKey      = "userID"
)

func SetContextData(c *gin.Context, key string, value interface{}) {
	if c.Keys == nil {
		c.Keys = make(map[string]interface{})
	}
	c.Keys[key] = value
}

func GetContextData(c *gin.Context, key string) interface{} {
	if c.Keys == nil {
		return nil
	}
	return c.Keys[key]
}

func SetCompanyID(c *gin.Context, companyID int) {
	SetContextData(c, CompanyKey, companyID)
}

func GetCompanyID(c *gin.Context) int {
	v := GetContextData(c, CompanyKey)
	if v != nil {
		return v.(int)
	} else {
		return 0
	}
}

func GetProjectID(c *gin.Context) (int, bool) {
	QueryProjectId := com.StrTo(c.DefaultQuery("project_id", "-1")).MustInt()
	HeaderProjectId := com.StrTo(c.Request.Header.Get("Project-Id")).MustInt()
	// TODO
	// check staff permission
	if QueryProjectId > 0 {
		return QueryProjectId, true
	}
	if HeaderProjectId > 0 {
		return HeaderProjectId, true
	}
	return -1, false
}

func SetStaffID(c *gin.Context, staffID int) {
	SetContextData(c, StaffKey, staffID)
}

func GetStaffID(c *gin.Context) int {
	v := GetContextData(c, StaffKey)
	if v != nil {
		return v.(int)
	} else {
		return 0
	}
}

func SetUserID(c *gin.Context, userID int) {
	SetContextData(c, UserKey, userID)
}

func GetUserID(c *gin.Context) int {
	v := GetContextData(c, UserKey)
	if v != nil {
		return v.(int)
	} else {
		return 0
	}
}

func GetAppID(c *gin.Context) int {
	strAppID := c.Param("appID")
	return com.StrTo(strAppID).MustInt()
}

func GetGroupID(c *gin.Context) int {
	groupID := c.GetHeader("groupid")
	return com.StrTo(groupID).MustInt()
}

func IsWebRequest(c *gin.Context) bool {
	appCode := GetContextData(c, "app_code").(string)
	return appCode == WEB_APP_CODE
}
