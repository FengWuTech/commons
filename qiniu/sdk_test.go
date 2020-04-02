package qiniu

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// SetUp
	//env := "dev"
	//conf.Setup(&env, true)
	// TestCase
	exitVal := m.Run()
	// TearDown
	// ...
	os.Exit(exitVal)
}

func TestUploadHttpFile(t *testing.T) {
	tk := GetUploadToken()
	fmt.Println(tk)
	var http = "http://mmbiz.qpic.cn/mmbiz_jpg/EmHVrGftic5WEcpqib6ZuS11yicKtuCj9QpGYNh6L8jdgtl2MaHdBb5TCsREZibcGsvtQmqCbEl4RSp5yc3tehCB2A/0"
	var dst = "qrcode.jpg"
	dst = UploadHttpFile(http, dst)
	fmt.Println(dst)
}
