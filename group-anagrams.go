func groupAnagrams(strs []string) [][]string {

	anagrams := make(map[string][]string)

	for _, str := range strs {
		var nums []int
		for _, ch := range str {
			nums = append(nums, int(ch))
		}
		sort.Ints(nums)
		sorted := ""
		for _, n := range nums {
			sorted += string(n)
		}

		if _, ok := anagrams[sorted]; !ok {
			anagrams[sorted] = []string{str}
		} else {
			anagrams[sorted] = append(anagrams[sorted], str)
		}
	}

	var res [][]string

	for _, group := range anagrams {
		res = append(res, group)
	}

	return res
}


// TOn SOn 
func groupAnagrams(strs []string) [][]string {
    keymap := make(map[[26]int][]string)

    for _, str := range strs {
        var count [26]int 

        for _, r := range str {
            count[r-'a']++ 
        }

        keymap[count] = append(keymap[count], str)
    }

    var res [][]string 
    for _, val := range keymap {
        res = append(res, val)
    }

    return res 
}