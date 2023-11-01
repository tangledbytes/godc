package util

import (
	"fmt"
	"math/rand"
	"strings"
)

// GenerateIntSeries generates a slice of integers from start to end
// where end is inclusive.
func GenerateIntSeries(start, end int) []int {
	var series []int
	for i := start; i <= end; i++ {
		series = append(series, i)
	}
	return series
}

// GenerateRandomIntSeries generates a slice of integers from start to end
// where end is inclusive and the slice is shuffled.
func GenerateRandomIntSeries(start, end int) []int {
	series := GenerateIntSeries(start, end)

	rand.Shuffle(end-start+1, func(i, j int) {
		series[i], series[j] = series[j], series[i]
	})

	return series
}

// CompareSliceUnordered compares two slices of comparable types
// and returns true if they contain the same elements.
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

// Assert panics if the condition is false.
func Assert(cond bool, msg ...string) {
	if !cond {
		fmsg := "assertion failed"
		if len(msg) != 0 {
			fmsg = fmt.Sprintf("%s: %s", fmsg, strings.Join(msg, ", "))
		}

		panic(fmsg)
	}
}
