package basicrecursion

func Print_N_to_1_using_Recursion(n int) {
	if n < 1 {
		return
	}
	println(n)
	Print_N_to_1_using_Recursion(n - 1)

}
