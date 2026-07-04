package medium


func groupAnagrams(strs []string) [][]string {
	groupsByFreq := make(map[[26]int8][]string)
	for _, word := range strs {
		freq := calculateFreq(word)
		groupsByFreq[freq] = append(groupsByFreq[freq], word)
	}
	result := make([][]string, 0, len(groupsByFreq))
	for _, wordFromGroup := range groupsByFreq {
		result = append(result, wordFromGroup)
	}
	return result
}

func calculateFreq(str string) [26]int8 {
	var freq [26]int8
	for _, symbol := range str {
		freq[symbol-'a']++
	}
	return freq
}
//--------------------------------------------------------------------
// func groupAnagrams(strs []string) [][]string {
// 	added := make([]bool, len(strs))
// 	groups := make([][]string, 0)
// 	var groupsCount int
// 	for i := range strs {
// 		if added[i] {
// 			continue
// 		}
// 		groups = append(groups, make([]string, 0))
// 		groupsCount++
// 		str := strs[i]
// 		actualFreq := calculateFreq(str)
// 		for j := i; j < len(strs); j++ {
// 			candidateFreq := calculateFreq(strs[j])
// 			if candidateFreq == actualFreq {
// 				groups[groupsCount] = append(groups[groupsCount], strs[j])
// 				added[j] = true
// 			}
// 		}
// 	}
// 	return groups
// }
