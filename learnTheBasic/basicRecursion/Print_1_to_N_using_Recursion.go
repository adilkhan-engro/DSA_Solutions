package basicrecursion

func Print_1_to_N_using_Recursion(n int) {
	if n < 1 {

		return
	}
	Print_1_to_N_using_Recursion(n - 1)
	println(n)

}

// printsomthngNtime
func PrintsomthngNtime(n int) {
	if n < 1 {
		return
	}
	println("adil")
	PrintsomthngNtime(n - 1)

}
