package knowbasicmath

func Check_number_is_prime(n int) bool {
	i := 2
	for n/2 >= i {
		if n%i == 0 {
			return false
		}
		i++
	}
	return true

}
func FindAllPrimeNo(n int) []int {
	i := 2
	primeList := []int{}
	for n >= i {
		if Check_number_is_prime(i) {
			primeList = append(primeList, i)
		}

		i++
	}
	return primeList
}
