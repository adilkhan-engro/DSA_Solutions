package main

import (
	"fmt"
	sorting1 "sorting/sorting-1"
)

func main() {
	nums := []int{13, 46, 24, 52, 20, 9}
	sorting1.BubbleSort(nums)
	fmt.Print(nums)
}
