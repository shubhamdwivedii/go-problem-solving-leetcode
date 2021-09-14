func longestPalindrome(s string) string {
	memo := make(map[string]int)

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if _, ok := memo[s[i:j+1]]; ok {
				continue
			}
			isPalim := isPalindrome(s[i : j+1])
			if isPalim {
				memo[s[i:j+1]] = len(s[i : j+1])
			}
		}
	}

	max := -1
	maxStr := string(s[0])
	for st, ln := range memo {
		if ln > max {
			max = ln
			maxStr = st
		}
	}
	return maxStr
}

func isPalindrome(s string) bool {
	ptrL := 0
	ptrR := len(s) - 1
	for ptrL < ptrR {
		if s[ptrL] != s[ptrR] {
			return false
		}
		ptrL++
		ptrR--
	}
	return true
}