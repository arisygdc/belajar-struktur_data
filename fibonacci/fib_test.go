package fibonacci

import (
	"log"
	"testing"
)

func TestFib(t *testing.T) {
	testTable := []struct {
		in     int
		expect float64
	}{
		{in: 1, expect: 1},
		{in: 3, expect: 2},
		{in: 5, expect: 5},
		{in: 16, expect: 987},
		{in: 59, expect: 956722026041},
	}

	for _, v := range testTable {
		res := Count(v.in)
		log.Printf("get %v, expect %v\n", res, v.expect)

		if res != v.expect {
			t.FailNow()
		}
	}
}
