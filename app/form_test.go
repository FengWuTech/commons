package app

import (
	"bytes"
	"github.com/FengWuTech/commons/logger"
	"github.com/FengWuTech/commons/setting"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// SetUp
	setting.AppSetting.LogSaveName = "app"
	setting.AppSetting.LogFileExt = "log"
	setting.AppSetting.LogSaveName = "logs/"
	setting.AppSetting.LogLevel = "info"
	setting.AppSetting.TimeFormat = "20060102"
	logger.Setup()
	// TestCase
	exitVal := m.Run()
	// TearDown
	os.Exit(exitVal)
}

type Req struct {
	Title     *string    `form:"title"`
	Status    *int       `form:"status"`
}

func TestBindAndValid(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("POST", "/?title=你好&status=1", bytes.NewBufferString(""))

	var form Req

	_, _, err := BindAndValid(c, &form)

	if err != nil {
		t.Errorf("Bind json fail err = %v", err)
	}

	if form.Title == nil || *form.Title != "你好" {
		t.Errorf("title != %v", "你好")
	}
	if form.Status == nil || *form.Status != 1 {
		t.Errorf("status != %v", 1)
	}
}
