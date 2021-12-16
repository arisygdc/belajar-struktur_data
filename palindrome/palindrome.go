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
	devider := leng / 2

	var head, tail string
	head = str[:devider]

	for i := devider; i > 0; i-- {
		tail = string(str[leng-i]) + tail
	}
	return head == tail
}
