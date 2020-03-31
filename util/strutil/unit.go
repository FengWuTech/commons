package strutil

import "fmt"

func YuanToFen(val float64) int {
	return int(val * 100.0)
}

func FenToYuan(val int) float64 {
	return float64(val) / 100.0
}

func FormatFen(val int) string {
	fen := FenToYuan(val)
	return fmt.Sprintf("%.2f", fen)
}
