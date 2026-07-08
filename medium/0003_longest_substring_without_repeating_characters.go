package medium

func lengthOfLongestSubstring(s string) int {
	var symbolConsist [128]bool
	totalCount := 0
	for left, right := 0, 0; right < len(s); right++ {
		if symbolConsist[s[right]] {
			for s[right] != s[left] {
				symbolConsist[s[left]] = false
				left++
			}
			left++
			continue
		}
		symbolConsist[s[right]] = true
		totalCount = max(totalCount, right+1-left)
	}
	return totalCount
}
