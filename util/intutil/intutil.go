package intutil

func InArray(num int, slice []int) bool {
	for _, v := range slice {
		if num == v {
			return true
		}
	}
	return false
}

func Intersect(sliceA []int, sliceB []int) []int {
	var result []int

	for _, a := range sliceA {
		for _, b := range sliceB {
			if a == b {
				result = append(result, a)
			}
		}
	}

	return result
}
