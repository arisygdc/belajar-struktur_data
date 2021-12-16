package palindrome

import (
	"log"
	"testing"
)

func TestPalindrome(t *testing.T) {
	testTabble := []struct {
		str    string
		expect bool
	}{
		{str: "aba", expect: true},
		{str: "Abba", expect: true},
		{str: "AbCba", expect: true},
		{str: "asdfdasf", expect: false},
		{str: "aaaaaaaaaaaaaaaaavaaaaaaaaaaaaaaaaa", expect: true},
	}
	for _, v := range testTabble {
		result := IsPalindrome(v.str)
		if result != v.expect {
			t.Fail()
		}
		log.Printf("%v result %v expected %v", v.str, result, v.expect)
	}
}
