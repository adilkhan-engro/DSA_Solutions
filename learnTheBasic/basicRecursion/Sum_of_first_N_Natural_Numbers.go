package basicrecursion

func Sum_of_first_N_Natural_Numbers(n int) int {
	if n <= 0 {
		return 0
	}

	return n + Sum_of_first_N_Natural_Numbers(n-1)
}
