package knowbasicmath

func Check_number_Armstrong_Number(n int) bool {
	AN := 0
	copyN := n
	no := n
	count := 0
	for n > 0 {
		count += 1
		n /= 10
	}
	for no > 0 {
		digit := no % 10
		AN = AN + pow(digit, count)
		no /= 10
	}
	if AN == copyN {
		return true
	}
	return false

}
func pow(n int, value int) int {
	powValue := n
	for value > 1 {
		powValue = powValue * n
		value -= 1
	}
	return powValue
}
