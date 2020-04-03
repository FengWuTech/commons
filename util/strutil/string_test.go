package strutil

import (
	"reflect"
	"testing"
)

func TestCreateRandomString(t *testing.T) {
	type args struct {
		len int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"random string length 10", args{10}, "",},
		{"random string length 20", args{20}, "",},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateRandomString(tt.args.len); len(got) != tt.args.len {
				t.Errorf("CreateRandomString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSliceToString(t *testing.T) {
	type args struct {
		ints  []int
		split string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"int slice", args{[]int{10, 20}, ","}, "10,20",},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntSliceToString(tt.args.ints, tt.args.split); got != tt.want {
				t.Errorf("IntSliceToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSliceToStringSlice(t *testing.T) {
	type args struct {
		ints []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"string slice", args{[]int{10, 20}}, []string{"10","20"},},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntSliceToStringSlice(tt.args.ints); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntSliceToStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToIntSlice(t *testing.T) {
	type args struct {
		str       string
		separator string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"string to int slice", args{"1,2", ","}, []int{1, 2},},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToIntSlice(tt.args.str, tt.args.separator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}