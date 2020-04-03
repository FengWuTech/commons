package strutil

import "testing"

func TestFenToYuan(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FenToYuan(tt.args.val); got != tt.want {
				t.Errorf("FenToYuan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatFen(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatFen(tt.args.val); got != tt.want {
				t.Errorf("FormatFen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYuanToFen(t *testing.T) {
	type args struct {
		val float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YuanToFen(tt.args.val); got != tt.want {
				t.Errorf("YuanToFen() = %v, want %v", got, tt.want)
			}
		})
	}
}