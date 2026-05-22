package basichashing

func FrequencyAlphabet(str string) map[string]int {
	// byteString := []byte(s)

	mapAlpha := map[string]int{}
	for _, alpha := range str {

		mapAlpha[string(alpha)] += 1

	}
	return mapAlpha

}

func FrequencyNum(nums []int) map[int]int {
	result := make(map[int]int)
	for _, num := range nums {
		result[num] += 1

	}
	return result
}
