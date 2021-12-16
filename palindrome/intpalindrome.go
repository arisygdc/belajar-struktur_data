package palindrome

import (
	"strconv"
	"strings"
)

func IsIntPalindrome(nums int) bool {
	str := strconv.FormatInt(int64(nums), 10)
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
