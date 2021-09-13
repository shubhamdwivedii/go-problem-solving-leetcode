func isHappy(n int) bool {
	if n == 1 {
		return true
	}

	str := strconv.Itoa(n)
	for len(str) > 1 || str == "7" { // 7 is only exception
		sum := 0
		for _, ch := range str {
			num, _ := strconv.Atoi(string(ch))
			sum += num * num
		}

		if sum == 1 {
			return true
		}
		str = strconv.Itoa(sum)
	}
	return false
}