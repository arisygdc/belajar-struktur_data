package palindrome

import "strconv"

func IsIntPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.FormatInt(int64(x), 10)

	j := len(str) - 1
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[j] {
			return false
		}
		j--
	}
	return true
}
