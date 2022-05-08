package longgestsubstr

func LengthOfLongestSubstring(s string) int {
	var (
		mp       = make(map[byte]bool)
		longgest int
	)
	if len(s) <= 1 {
		return len(s)
	}

	for i := 0; i < len(s)-1; i++ {
		for j := i; j < len(s) && !mp[s[j]]; j++ {
			mp[s[j]] = true
		}

		if len(mp) > longgest {
			longgest = len(mp)
		}
		mp = make(map[byte]bool)
	}

	return longgest
}
