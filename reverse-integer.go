func reverse(x int) int {
	rangeL := int64(math.Pow(float64(2), 31)) * -1
	rangeR := int64(math.Pow(float64(2), 31) - 1)

	mul := 1
	if x < 0 {
		x = x * -1
		mul = -1
	}

	reversed := 0
	for x > 0 {
		n := x % 10
		reversed = reversed*10 + n
		x = x / 10
		fmt.Println(reversed, x)
	}

	if rangeL <= int64(reversed) && int64(reversed) <= rangeR {
		return reversed * mul
	} else {
		return 0
	}
}