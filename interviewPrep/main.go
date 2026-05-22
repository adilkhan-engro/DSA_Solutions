package main

import "log"

// func reverseString(s []byte) {
// 	i := 0
// 	for i < len(s)/2 {
// 		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
// 		i++
// 	}

// }

func main() {

	// s := "adil"
	// bytes := []byte(s)

	// reverseString(bytes)
	// print(string(bytes))
	arr := []int{4, 3, 2, 7, 8, 2, 3, 1}
	result := findDuplicates(arr)
	log.Println(result)
}
func reverseString(s []byte) {
	i := 0
	for i < len(s)/2 {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
		i++
	}

}
func findDuplicates(nums []int) []int {
	duplicates := map[int]int{}
	result := []int{}
	for _, value := range nums {
		log.Println(value)
		if duplicates[value] == 1 {
			result = append(result, value)

		} else {
			duplicates[value] += 1
		}
	}
	return result

}

func removeDuplicates(nums []int) int {
	i := 1
	j := 0
	for i = range len(nums) {
		if nums[j] == nums[i] {
			continue
		} else {
			nums[j+1] = nums[i]
			j++
		}

	}
	return j + 1
}

func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}
	for i,num:= range nums{
        remain:=target-num
        if value ,found:=hashMap[remain];found{
            return []int{value,i}
        }
        hashMap[num]=i

    }

	return []int{}
}

