
func reverseOnlyLetters(s string) string {
	ptrL := 0
	ptrR := len(s) - 1

	for ptrL < ptrR {
		if isLetter(rune(s[ptrL])) {
			if isLetter(rune(s[ptrR])) {
				// s[ptrL], s[ptrR] = s[ptrR], s[ptrL]
				fmt.Println("Swap a and b", s[ptrL], s[ptrR])
				s = s[:ptrL] + string(s[ptrR]) + s[ptrL+1:ptrR] + string(s[ptrL]) + s[ptrR+1:]
				ptrL++
				ptrR--
			} else {
				ptrR--
			}
		} else {
			ptrL++
		}
	}
	return s
}

func isLetter(c rune) bool {
	if c >= 'A' && c <= 'Z' {
		return true
	}

	if c >= 'a' && c <= 'z' {
		return true
	}

	fmt.Println("Not a letter", string(c))

	return false
}