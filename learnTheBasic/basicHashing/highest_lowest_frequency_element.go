package basichashing

func Highest_lowest_frequency_element(nums []int) (int, int) {
	hightest := nums[0]
	loweset := nums[0]
	for _, num := range nums {
		if num > hightest {
			hightest = num

		}
		if loweset > num {
			loweset = num
		}

	}
	return hightest, loweset
}
