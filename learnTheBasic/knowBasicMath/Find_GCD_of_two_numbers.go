package knowbasicmath

func Find_GCD_of_two_numbers(n1, n2 int) int {
	smallNo := n2
	gcd := -1
	if n1 > n2 {
		if n1%n2 == 0 {
			return n2
		}

	} else if n2 > n1 {
		if n2%n1 == 0 {
			return n1
		}

		smallNo = n1
	}
	for i := 1; i < smallNo; i++ {
		if n1%i == 0 && n2%i == 0 {
			gcd = i
		}
	}
	return gcd

}
func FindGCD(n1, n2 int) int {
	for n2 != 0 {
		n1, n2 = n2, n1%n2
	}
	return n1
}

func FindGCD2(n1, n2 int) int {
	for n1%n2 != 0 {
		if n1 > n2 {
			subVal := n1 - n2
			if subVal > n2 {
				n1 = subVal
			} else {
				n1, n2 = n2, subVal
			}
		} else {
			n1, n2 = n2, n1
		}

	}
	return n2

}
