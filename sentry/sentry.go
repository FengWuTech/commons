package sentry

import (
	"github.com/FengWuTech/commons/setting"
	"github.com/getsentry/sentry-go"
)

func Setup() {
	dsn := setting.RemoteSetting.SentryDSN
	sentry.Init(sentry.ClientOptions{
		Dsn: dsn,
		AttachStacktrace: true,
	})
}
