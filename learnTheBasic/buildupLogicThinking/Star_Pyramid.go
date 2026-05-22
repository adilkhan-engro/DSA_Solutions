package buildupLogicThinking

import "fmt"

func Star_Pyramid(n int) {

	for i := 1; i <= n; i++ {

		// spaces
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}

		// stars
		for k := 1; k <= (2*i)-1; k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
