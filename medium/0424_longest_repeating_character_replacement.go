package medium

func characterReplacement(s string, k int) int {
	var freq [26]int
	start, end, count := 0, 0, 0
	for ; end < len(s); end++ {
		freq[s[end]-'A']++
		count = max(count, freq[s[end]-'A'])
		if end-start+1-count > k {
			freq[s[start]-'A']--
			start++
		}
	}
	return end - start +1
}
