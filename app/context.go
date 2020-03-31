package app

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
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

func GetCompanyID(c *gin.Context) (int, bool) {
	companyID := GetContextData(c, "company_id")
	if companyID != nil {
		return companyID.(int), true
	} else {
		return -1, false
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

func GetStaffID(c *gin.Context) (int, bool) {
	staffID := GetContextData(c, "staff_id")
	if staffID != nil {
		return int(staffID.(int)), true
	} else {
		return -1, false
	}
}

func SetUserID(c *gin.Context, userID int) {
	if c.Keys == nil {
		c.Keys = make(map[string]interface{})
	}
	c.Keys["userID"] = userID
}

func GetUserID(c *gin.Context) int {
	if c.Keys == nil {
		return 0
	}

	v, ok := c.Keys["userID"]
	if ok {
		return int(v.(int))
	} else {
		return 0
	}
}

func GetAppID(c *gin.Context) (appID int) {
	strAppID := c.Param("appID")
	return com.StrTo(strAppID).MustInt()
}

func GetGroupID(c *gin.Context) int {
	groupID := c.GetHeader("groupid")
	return com.StrTo(groupID).MustInt()
}
