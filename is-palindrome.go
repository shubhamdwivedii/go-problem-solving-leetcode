func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reverse := 0

	y := x

	for y > 0 {
		n := y % 10
		y = y / 10
		reverse *= 10
		reverse += n
	}

	if reverse == x {
		return true
	}
	return false
}