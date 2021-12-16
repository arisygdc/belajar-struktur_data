package longgestsubstr

import (
	"log"
	"testing"
)

func TestLong(t *testing.T) {
	testTable := []struct {
		str    string
		expect int
	}{
		{str: "abcabcbb", expect: 3},
		{str: "bbbb", expect: 1},
		{str: "pwwkew", expect: 3},
		{str: "", expect: 0},
		{str: " ", expect: 1},
		{str: "aab", expect: 2},
		{str: "dvdf", expect: 3},
	}

	for _, v := range testTable {
		res := LengthOfLongestSubstring(v.str)
		if res != v.expect {
			t.Fail()
		}

		log.Printf("result %v expect %v", res, v.expect)
	}
}
