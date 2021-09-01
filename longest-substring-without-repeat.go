func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	max := -1
	start := 0
	end := 0
	enc := make(map[byte]int)
	for end < len(s) {
		if idx, ok := enc[s[end]]; !ok { // Not Encountered
			enc[s[end]] = end
		} else { // already encountered > reset start = idx+1
			if idx >= start {
				start = idx + 1 // idx of encountered.
			}
			if s[start] == s[end] {
				start = end
			}
			enc[s[end]] = end
		}
		end++
		if end-start > max {
			max = (end + 1) - (start + 1) // added 1 to start and end to shift index from 0..n to 1..n
		}

	}
	return max
}
