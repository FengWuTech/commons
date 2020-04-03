package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func TestGetContextData(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	key := "key"
	value := "value"

	c.Set(key, value)
	v := GetContextData(c, key)

	if v != value {
		t.Errorf("Get Context of %v !=  %v, should be %v", key, v, value)
	}
}

func TestSetContextData(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	key := "key"
	value := "value"

	SetContextData(c, key, value)
	v := GetContextData(c, key)

	if v != value {
		t.Errorf("Set Context of %v !=  %v, should be %v", key, v, value)
	}
}

func TestSetContextIds(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	testCases := []struct {
		Key       string
		Value     int
		SetMethod func(*gin.Context, int)
		GetMethod func(*gin.Context) int
		Remark    string
	}{
		{"companyID", 1, SetCompanyID, GetCompanyID,"公司ID"},
		{"staffID", 1, SetStaffID, GetStaffID,"员工ID"},
		{"userID", 1, SetUserID, GetUserID,"用户ID"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Test %s: set & get %s", tc.Remark, tc.Key), func(t *testing.T) {
			tc.SetMethod(c, tc.Value)
			v := tc.GetMethod(c)
			if v != tc.Value {
				t.Errorf("%v got %v; want %v", tc.Remark, v, tc.Value)
			}
		})
	}
}