package easy

import (
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for start, end := 0, len(s)-1; start < end; start, end = start+1, end-1 {
		for start < len(s) && !isAlphanumeric(s[start]) {
			start++
		}
		for end > 0 && !isAlphanumeric(s[end]) {
			end--
		}
		if start < len(s) && end > 0 && s[start] != s[end] {
			return false
		}
	}
	return true
}

func isAlphanumeric(b byte) bool {
    return ('a' <= b && b <= 'z') || 
           ('A' <= b && b <= 'Z') || 
           ('0' <= b && b <= '9')
}
