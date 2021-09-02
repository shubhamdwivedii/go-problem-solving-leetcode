func longestCommonPrefix(strs []string) string {
	prefix := ""
	curr := 0

	if len(strs) == 0 {
		return prefix
	}

	for {
		if len(strs[0]) == 0 {
			return prefix
		}
		if curr >= len(strs[0]) {
			return prefix
		}
		firstch := strs[0][curr]
		matched := true
		for _, str := range strs {
			if curr >= len(str) {
				return prefix
			}

			if str[curr] == firstch {
				continue
			} else {
				matched = false
				return prefix
			}
		}
		if matched {
			prefix += string(firstch)
		}
		curr++
	}
}