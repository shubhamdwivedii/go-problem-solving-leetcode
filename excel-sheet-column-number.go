func titleToNumber(columnTitle string) int {
	sum := float64(0)

	for i, ch := range columnTitle {
		twoSixTimes := len(columnTitle) - 1 - i

		power := math.Pow(26, float64(twoSixTimes))

		sum += float64(alphaToInt(ch)) * power
	}

	return int(sum)
}

func alphaToInt(a rune) int {
	sum := int(a) - int('A') + 1
	return sum
}