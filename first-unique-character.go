func firstUniqChar(s string) int {
	charMap := make(map[string]int)

	for i, ch := range s {
		if _, ok := charMap[string(ch)]; !ok {
			charMap[string(ch)] = i
		} else {
			charMap[string(ch)] = -1
		}
	}
	min := len(s)
	for _, i := range charMap {
		if i > -1 {
			if i < min {
				min = i
			}
		}
	}
	if min == len(s) {
		return -1
	}
	return min
}