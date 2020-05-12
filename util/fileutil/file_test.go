package fileutil

import (
	"testing"
)

func TestExists(t *testing.T) {
	cases := map[string]bool{
		"file.go":   true,
		"file.go.1": false,

		"file_test.go":   true,
		"file_test.go.1": false,

		"../fileutil":   true,
		"../fileutil_1": false,

		"../../": true,
	}

	for fpath, expect := range cases {
		if Exists(fpath) == expect {
			t.Logf("pass, fpath=%s, expect= %v", fpath, expect)
		} else {
			t.Errorf("fail, fpath=%s, expect= %v", fpath, expect)
		}

	}

}

