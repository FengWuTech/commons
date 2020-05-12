package sentry

import (
	"github.com/FengWuTech/commons/setting"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// SetUp
	setting.RemoteSetting.SentryDSN = ""
	// TestCase
	exitVal := m.Run()
	// TearDown
	os.Exit(exitVal)
}

func TestSetup(t *testing.T) {
	Setup()

	r := gin.New()
	r.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	if hub := sentry.CurrentHub().Clone(); hub == nil {
		t.Errorf("sentry setup error")
	}
}
