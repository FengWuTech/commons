package util

import "testing"

func TestStructCopyUseJson(t *testing.T) {
	type args struct {
		src interface{}
		dst interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"copy use json", args{"json", ""},},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StructCopyUseJson(&tt.args.src, &tt.args.dst)
			if tt.args.src != tt.args.dst {
				t.Errorf("StructCopyUseJson() = %v, want %v", tt.args.src, tt.args.dst)
			}
		})
	}
}

func TestIsNil(t *testing.T) {
	type args struct {
		i interface{}
	}
	var temp *string
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"is nil pointer", args{temp}, true},
		{"is nil 1", args{1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNil(tt.args.i); got != tt.want {
				t.Errorf("IsNil() = %v, want %v", got, tt.want)
			}
		})
	}
}