package knowbasicmath

func Reverse_Digits_of_A_Number(n int) int {
	no := 0
	for n > 0 {
		num := n % 10
		no = no*10 + num
		n = n / 10

	}
	return no
}
