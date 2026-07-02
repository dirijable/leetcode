package easy

func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
    }
	var freqActual [26]int
    var freqCandidate [26]int
 	for idx := range s {
		freqActual[s[idx] - 'a']++
		freqCandidate[t[idx] - 'a']++
	}
	return freqActual == freqCandidate
}