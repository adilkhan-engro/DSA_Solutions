package sorting1

func BubbleSort(nums []int) {
	for j := 0; j < len(nums); j++ {
		for i := 0; i < len(nums)-1-j; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}

		}
		print(j)
	}

}
