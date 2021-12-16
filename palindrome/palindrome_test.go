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

func TestIntPalindrome(t *testing.T) {
	testTabble := []struct {
		num    int
		expect bool
	}{
		{num: 121, expect: true},
		{num: -121, expect: false},
		{num: 556323, expect: false},
		{num: 373, expect: true},
		{num: 6776, expect: true},
		{num: 10, expect: false},
		{num: -101, expect: false},
	}
	for _, v := range testTabble {
		result := IsIntPalindrome(v.num)
		if result != v.expect {
			t.Fail()
		}
		log.Printf("%v result %v expected %v", v.num, result, v.expect)
	}
}
