package slider

import (
	"testing"
)

func TestCheckTencentCaptcha(t *testing.T) {
	type args struct {
		req    CaptchaReq
		userIp string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test captcha", args{CaptchaReq{"t", "r"}, "1.1.1.1"}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckTencentCaptcha(tt.args.req, tt.args.userIp); got != tt.want {
				t.Errorf("CheckTencentCaptcha() = %v, want %v", got, tt.want)
			}
		})
	}
}