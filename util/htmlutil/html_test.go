package htmlutil

import (
	"fmt"
	"testing"
)

func TestTrimHtml(t *testing.T) {
	text := TrimHtml("<s>sss</s>")
	fmt.Printf("%v", text)
}
