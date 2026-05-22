package buildupLogicThinking

func Inverted_Star_Pyramid(n int) {
	for i := n; i > 0; i-- {
		for k := i; k < n; k++ {
			print(" ")
		}
		// print star
		for j := 2*i - 1; j > 0; j-- {
			print("*")
		}

		println()

	}
}
