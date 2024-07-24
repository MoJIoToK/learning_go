package statistics

func Avg(nums []float64) float64 {
	var sum float64

	if len(nums) == 0 {
		return 0
	}

	for _, n := range nums {
		sum += n
	}
	return sum / float64(len(nums))
}

// Задание 33.4.1
func MaxN(nums []float64) float64 {
	max := nums[0]
	for _, v := range nums {
		if max < v {
			max = v
		}
	}
	return max
}
