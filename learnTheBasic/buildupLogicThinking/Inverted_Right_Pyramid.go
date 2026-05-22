package buildupLogicThinking

func Inverted_Right_Pyramid(n int) {
	i := n
	for i>0 {
		j := 0
		for j < i {
			print("*")
			j+=1

		}
		i-=1
		println()
	}
}
