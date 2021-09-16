func generateParenthesis(n int) []string {
	// only add ( when open < n
	// only add ) if closed < open
	// valid if open == closed == n

	stack := ""
	var res []string
	backtrack(n, 0, 0, stack, &res)

	return res

}

func backtrack(n int, open int, closed int, stack string, res *[]string) {
	if open == closed && open == n {
		*res = append(*res, stack)
		return
	}

	if open < n {
		stack += "("
		backtrack(n, open+1, closed, stack, res)
		stack = stack[:len(stack)-1]
	}

	if closed < open {
		stack = stack + ")"
		backtrack(n, open, closed+1, stack, res)
		stack = stack[:len(stack)-1]
	}
}