package arraysslices

// Sum takes numbers and returns their sum
func Sum(numbers []int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int
	for _, nums := range numbers {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(nums[1:]))
		}
	}
	return sums
}
