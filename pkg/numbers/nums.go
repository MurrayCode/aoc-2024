package numbers

func Sum(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}
