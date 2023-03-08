package util

func GenerateIntSeries(start, end int) []int {
	var series []int
	for i := start; i <= end; i++ {
		series = append(series, i)
	}
	return series
}

func CompareSliceUnordered[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	mpa := make(map[T]struct{}, len(a))
	for _, v := range a {
		mpa[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := mpa[v]; !ok {
			return false
		}
	}

	return true
}
