package sort

import (
	"math/rand"
	"testing"
)

//region bubbleSort

func BenchmarkBubbleSort(b *testing.B) {

	b.Run("Max is ten times less than the previous one", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(1, 100)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("small arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 100)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("Max is five times greater than the previous one", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(50, 100)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("Max is fifty times greater than the previous one", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(5000, 1000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

}

//endregion

//region selectionSort

func BenchmarkSelectionSort(b *testing.B) {

	b.Run("Max is ten times less than the previous one", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(1, 100)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 100)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("Max is five times greater than the previous one", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(50, 100)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("Max is fifty times greater than the previous one", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(5000, 1000)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})
}

//endregion

//region insertionSort

func BenchmarkInsertionSort(b *testing.B) {

	b.Run("Max is ten times less than the previous one", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(1, 100)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 100)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("Max is five times greater than the previous one", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(50, 100)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("Max is fifty times greater than the previous one", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(5000, 1000)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})
}

//endregion

//region mergeSort

func BenchmarkMergeSort(b *testing.B) {

	b.Run("Max is ten times less than the previous one", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(1, 100)
			b.StartTimer()
			mergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 100)
			b.StartTimer()
			mergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("Max is five times greater than the previous one", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(50, 100)
			b.StartTimer()
			mergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			mergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("Max is fifty times greater than the previous one", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(5000, 1000)
			b.StartTimer()
			mergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			mergeSort(ar)
			b.StopTimer()
		}
	})
}

//endregion

//region quickSort

func BenchmarkQuickSort(b *testing.B) {

	b.Run("Max is ten times less than the previous one", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(1, 100)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 100)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("Max is five times greater than the previous one", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(50, 100)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("Max is fifty times greater than the previous one", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(5000, 1000)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})
}

//endregion

//region worst case

func BenchmarkWorstCase(b *testing.B) {

	b.Run("Bubble - worst", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := worstSlice()
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("Quick worst", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := worstSlice()
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("Insert worst", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := worstSlice()
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})
}

//endregion

//region best case

func BenchmarkBestCase(b *testing.B) {

	b.Run("Bubble - best", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := bestSlice()
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("Quick best", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := bestSlice()
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("Insert best", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := bestSlice()
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})
}

//endregion

//region best slice

func bestSlice() []int {
	size := 1000
	ar := make([]int, size)
	for i := 0; i < size; i++ {
		ar[i] = i
	}
	return ar
}

//endregion

//region worst slice

func worstSlice() []int {
	size := 1000
	ar := make([]int, size)
	for i := size; i < 0; i-- {
		ar[i] = i
	}
	return ar
}

//endregion

//region random slice

func generateSlice(max, size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = rand.Intn(max*2) - max
	}

	return ar
}

//endregion
