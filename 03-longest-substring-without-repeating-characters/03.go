package _3_longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	length := 0
	left := 0
	recorder := make(map[string]int)
	for right := range s {
		rightString := string(s[right])
		recorder[rightString]++

		for recorder[rightString] > 1 {
			recorder[string(s[left])]--
			left++
		}
		length = max(length, right-left+1)
		right++
	}
	return length
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
