func isPowerOfThree(n int) bool {
	if n == 3 || n == 1 {
		return true
	}
	if n%3 != 0 {
		return false
	}
	num := 9

	for num < n {
		num *= 3
	}

	if num == n {
		return true
	}
	return false
}