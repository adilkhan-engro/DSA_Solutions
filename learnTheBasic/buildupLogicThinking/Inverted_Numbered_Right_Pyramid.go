package buildupLogicThinking

func Inverted_Numbered_Right_Pyramid(n int) {
	i := n
	for i > 0 {
		j := 1
		for j < i+1 {
			print(j)
			j += 1
		}
		i -= 1
		println()

	}
}
