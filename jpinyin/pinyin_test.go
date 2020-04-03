package jpinyin

import (
	"reflect"
	"testing"
)

func TestGetInitiallist(t *testing.T) {
	list := GetInitiallist("小区")
	if !reflect.DeepEqual(list, []string{"XQ"}) {
		t.Errorf("GetInitiallist %v != %v", list, []string{"XQ"})
	}
}