package knowbasicmath

func Count_digits_in_a_number(n int) int {
	count := 0
	if n < 0 {
		n = n * -1
	}
	for n > 0 {
		count += 1
		n = n / 10

	}
	return count
}
