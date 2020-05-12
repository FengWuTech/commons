package app

import (
	"testing"
)

func TestGetMsgOK(t *testing.T) {
	var code = SUCCESS
	msg := GetMsg(code)

	if msg != "ok" {
		t.Errorf("Msg of code %v !=  %v, should be %v", code, msg, "ok")
	}
}

func TestGetMsgForbidden(t *testing.T) {
	var code = FORBID
	msg := GetMsg(code)

	if msg != "操作被禁止" {
		t.Errorf("Msg of code %v !=  %v, should be %v", code, msg, "操作被禁止")
	}
}

func TestGetMsgNotDefined(t *testing.T) {
	var code = 123141
	msg := GetMsg(code)

	if msg != "fail" {
		t.Errorf("Msg of code %v !=  %v, should be %v", code, msg, "fail")
	}
}