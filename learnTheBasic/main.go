package main

import (
	"fmt"
	basichashing "mydsa/basicHashing"
)

func main() {
	arr := []int{10,5,10,15,10,5}
	// basicrecursion.Reverse_Array(arr, 0, (len(arr) - 1))
	// log.Println(arr)
	hightest, loweset := basichashing.Highest_lowest_frequency_element(arr)
	fmt.Println("result:", hightest, loweset)
}
