package basicrecursion

func Palindrome(s string, i int) bool {
	if i >= len(s)/2 {
		return true
	}

	if s[i] != s[len(s)-1-i] {
		return false
	}

	return Palindrome(s, i+1)
}
