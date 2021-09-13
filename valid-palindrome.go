
func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	trimmed := ""
	for _, ch := range s {
		if rune(ch) >= 'a' && rune(ch) <= 'z' {
			trimmed += string(ch)
		} else if rune(ch) >= '0' && rune(ch) <= '9' {
			trimmed += string(ch)
		}
	}
	fmt.Println(trimmed)

	halfPoint := len(trimmed) / 2

	for i := 0; i <= halfPoint-1; i++ {
		if trimmed[i] != trimmed[len(trimmed)-1-i] {
			return false
		}
	}

	return true
}