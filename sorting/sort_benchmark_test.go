package mysorting

import (
	"sort"
	"testing"
)

func TestMergesort(t *testing.T) {
	array := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	expected := make([]int, len(array))
	copy(expected, array)
	sort.Ints(expected)

	sortedArray := mergesort(array)

	if !equals(expected, sortedArray) {
		t.Errorf("sortedArray not equal expected:\nsortedArray: %v\nExpected:%v", sortedArray, expected)

	}
}

func TestQuicksort(t *testing.T) {
	array := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	expected := make([]int, len(array))
	copy(expected, array)
	sort.Ints(expected)

	sortedArray := qsort(array)

	if !equals(expected, sortedArray) {
		t.Errorf("sortedArray not equal expected:\nsortedArray: %v\nExpected:%v", sortedArray, expected)
	}
}

func BenchmarkStdSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var dataset []int
		for x := 100; x > 0; x-- {
			dataset = append(dataset, x)
		}
		b.StartTimer()
		sort.Ints(dataset)
	}
}

func BenchmarkQuicksort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var dataset []int
		for x := 100; x > 0; x-- {
			dataset = append(dataset, x)
		}
		b.StartTimer()
		qsort(dataset)
	}
}

func BenchmarkMergesort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var dataset []int
		for x := 100; x > 0; x-- {
			dataset = append(dataset, x)
		}
		b.StartTimer()
		mergesort(dataset)
	}
}

func equals(x []int, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i, j := range x {
		if j != y[i] {
			return false
		}
	}
	return true
}
