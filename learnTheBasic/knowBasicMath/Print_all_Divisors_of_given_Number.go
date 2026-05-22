package knowbasicmath

import "log"

func Print_all_Divisors_of_given_Number(n int) []int {
	result := []int{}
	i := 1
	for n/2 >= i {
		if n%i == 0 {
			result = append(result, i)
			log.Println(result)

		}
		i++
	}
	result = append(result, n)
	return result
}
