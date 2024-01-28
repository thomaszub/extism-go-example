package calc

import "math"

func StdDev(nums []float64) float64 {
	mean := Mean(nums)
	std := 0.0
	for _, n := range nums {
		std += math.Pow(n-mean, 2)
	}
	std /= float64(len(nums) - 1)
	return math.Sqrt(std)
}

func Mean(nums []float64) float64 {
	mean := 0.0
	for _, n := range nums {
		mean += n
	}
	return mean / float64(len(nums))
}
