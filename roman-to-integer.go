func romanToInt(s string) int {
	sum := getVal(rune(s[len(s)-1]))
	curr := len(s) - 1
	for curr >= 1 {
		sum += getDiff(rune(s[curr]), rune(s[curr-1]))
		curr--
	}
	return sum
}

func getDiff(ch rune, prv rune) int {
	if getVal(prv) >= getVal(ch) {
		return getVal(prv)
	} else {
		return -getVal(prv)
	}
}

func getVal(ch rune) int {
	switch ch {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return -1
	}
}