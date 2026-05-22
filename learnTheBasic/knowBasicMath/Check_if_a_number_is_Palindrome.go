package knowbasicmath

// func Check_number_is_Palindrome(n int) bool {
// 	copy := n
// 	reverseNo := 0
// 	for n > 0 {
// 		num := n % 10
// 		reverseNo = reverseNo*10 + num
// 		n = n / 10
// 	}
// 	if reverseNo == copy {
// 		return true
// 	}
// 	return false
// }
func Check_number_is_Palindrome(n int)bool{
	reverseHalf:=0
	for n>reverseHalf{
		reverseHalf=reverseHalf*10+n%10
		n/=10
	}
	if (reverseHalf == n )||(n==reverseHalf/10) {
		return true
	}
	return  false
	
}
