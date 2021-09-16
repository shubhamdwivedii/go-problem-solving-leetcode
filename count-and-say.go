func countAndSay(n int) string {
	prevSay := "1"
	for i := 2; i <= n; i++ {
		currStr := ""
		curr := string(prevSay[0])
		currCount := 1

		for i := 1; i < len(prevSay); i++ {
			if string(prevSay[i]) != curr {
				currStr += fmt.Sprintf("%v%v", currCount, curr)
				curr = string(prevSay[i])
				currCount = 1
			} else {
				currCount++
			}
		}
		currStr += fmt.Sprintf("%v%v", currCount, curr)
		prevSay = currStr
	}
	return prevSay
}
