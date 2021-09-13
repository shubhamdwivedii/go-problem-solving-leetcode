func plusOne(digits []int) []int {
	buff := 1

	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + buff
		if sum > 9 {
			digits[i] = sum % 10
			buff = sum / 10
		} else {
			digits[i] = sum
			buff = 0
			break
		}
	}

	fmt.Println(buff, digits)

	if buff != 0 {
		digits = append([]int{buff}, digits...)
	}

	return digits
}