package task33103

import (
	"math/rand"
	"sort"
	"testing"
)

func sampleData() []int {
	var data []int
	for i := 0; i < 1_000; i++ {
		data = append(data, rand.Intn(1000))
	}

	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
	return data
}

func BenchmarkInts(b *testing.B) {
	data := sampleData()
	for i := 0; i < b.N; i++ {
		sort.Ints(data)
	}
}
