package sorting1

func SelectionSort(nums []int) {
	// pick first one and compare with all element if find min value swap the index
	for i := 0; i < len(nums); i++ {
		minValueIndex := i
		for j := i; j < len(nums); j++ {
			if nums[minValueIndex] > nums[j] {
				minValueIndex = j
			}
		}
		nums[i], nums[minValueIndex] = nums[minValueIndex], nums[i]
	}

}
