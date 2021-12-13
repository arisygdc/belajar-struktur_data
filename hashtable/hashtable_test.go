package hashtable

import (
	"log"
	"testing"
)

func TestHashValue(t *testing.T) {
	testTable := []struct {
		value string
	}{
		{value: "a"}, {value: "b"}, {value: "c"}, {value: "d"}, {value: "e"}, {value: "f"},
		{value: "g"}, {value: "h"}, {value: "i"}, {value: "j"}, {value: "k"}, {value: "l"},
		{value: "m"}, {value: "n"}, {value: "o"}, {value: "p"}, {value: "q"}, {value: "r"},
		{value: "s"}, {value: "t"}, {value: "u"}, {value: "v"}, {value: "w"}, {value: "x"},
		{value: "y"}, {value: "z"}, {value: "aa"}, {value: "ab"}, {value: "ac"}, {value: "ad"},
		{value: "ag"}, {value: "ah"}, {value: "ai"}, {value: "aj"}, {value: "ak"}, {value: "al"},
		{value: "am"}, {value: "an"}, {value: "ao"}, {value: "ap"}, {value: "aq"}, {value: "ar"},
		{value: "as"}, {value: "at"}, {value: "au"}, {value: "av"}, {value: "aw"}, {value: "ax"},
		{value: "ay"}, {value: "az"}, {value: "ae"}, {value: "af"}, {value: "kbfweakbjjewbf00"},
		{value: "bkjwaehlofewa"}, {value: "vvv"}, {value: "zvz"}, {value: "kbfweakbjjewbf10"},
		{value: "kaafweakbjjewbf00"},
	}
	del := "bkjwaehlofewa"
	ht := NewHashTable()
	for _, a := range testTable {
		ht.Put(a.value)
	}
	for _, a := range testTable {
		val := ht.Get(a.value)
		log.Printf("mendapatkan %v, expect %v", val, a.value)
		if val != a.value {
			t.FailNow()
		}
	}
	ht.Remove(del)
	if ht.Get(del) != "" {
		log.Println("Gagal ketika penghapusan")
		t.FailNow()
	}
}
