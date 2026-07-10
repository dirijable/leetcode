package medium

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var (
		freq1, freq2 [26]int
	)
	for i := range len(s1) {
		freq1[s1[i]-'a']++
		freq2[s2[i]-'a']++
	}
	if freq1 == freq2 {
		return true
	}
	for left, right := 0, len(s1); right < len(s2); right++ {
		freq2[s2[right] - 'a']++
		freq2[s2[left] - 'a']--
		left++
		if freq1 == freq2 {
			return true
		}
	}
	return false
}
