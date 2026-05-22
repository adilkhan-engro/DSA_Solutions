package basicrecursion

func Reverse_Array(arr []int, left, right int) {
	if left >= right {
		return
	}

	arr[left], arr[right] = arr[right], arr[left]

	Reverse_Array(arr, left+1, right-1)
}
