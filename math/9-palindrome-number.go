package math

func isPalindrome(x int) bool {
	reverse := 0
	dummy := x

	for dummy > 0 {
		reverse = reverse*10 + (dummy % 10)
		dummy /= 10
	}

	return reverse == x
}
