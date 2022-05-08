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

	j := leng - 1
	for i := 0; i < leng/2; i++ {
		if str[i] != str[j] {
			return false
		}
		j--
	}
	return true
}
