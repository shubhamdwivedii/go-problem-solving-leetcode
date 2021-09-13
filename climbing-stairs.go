func climbStairs(n int) int {
	memo := make([]int, n+1) // length n+1
	memo[0] = 1              // 1 way to climb zero steps (ie: not to climb)
	memo[1] = 1              // 1 way to climb 1 step (ie: 1 step)

	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}

	return memo[n]
}
