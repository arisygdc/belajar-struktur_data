package palindrome

import (
	"strings"
)

func IsPalindrome(str string) bool {
	leng := len(str)
	if leng == 1 {
		return true
	}

	str = strings.ToLower(str)

	for i, j := 0, leng-1; i < leng/2; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
