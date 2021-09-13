func isAnagram(s string, t string) bool {
	sCount := make(map[string]int)
	tCount := make(map[string]int)

	for _, ch := range s {
		if _, ok := sCount[string(ch)]; !ok {
			sCount[string(ch)] = 1
		} else {
			sCount[string(ch)]++
		}
	}

	for _, ch := range t {
		if _, ok := tCount[string(ch)]; !ok {
			tCount[string(ch)] = 1
		} else {
			tCount[string(ch)]++
		}
	}

	if len(sCount) != len(tCount) {
		return false
	}

	for k, v := range sCount {
		if count, ok := tCount[k]; !ok {
			return false
		} else {
			if count != v {
				return false
			}
		}
	}

	return true

}