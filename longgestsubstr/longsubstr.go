package longgestsubstr

func LengthOfLongestSubstring(s string) int {
	var (
		mp       map[rune]rune = make(map[rune]rune)
		longgest int
	)

	for _, v := range s {
		if _, ok := mp[v]; !ok {
			mp[v] = v
			if len(mp) > longgest {
				longgest = len(mp)
			}
		} else {
			mp = make(map[rune]rune)
			mp[v] = v
			if len(mp) > longgest {
				longgest = len(mp)
			}
		}

	}
	return longgest
}
