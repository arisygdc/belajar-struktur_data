package longgestsubstr

type LonggestSubStr struct {
	mp       map[byte]bool
	str      string
	longgest int
}

func NewLonggestSubStr(str string) LonggestSubStr {
	return LonggestSubStr{
		mp:  make(map[byte]bool),
		str: str,
	}
}

func (lss *LonggestSubStr) Count() {
	if len(lss.str) <= 1 {
		lss.longgest = len(lss.str)
		return
	}

	for i := 0; i < len(lss.str)-1; i++ {
		for j := i; j < len(lss.str) && !lss.mp[lss.str[j]]; j++ {
			lss.mp[lss.str[j]] = true
		}
		lss.reset()
	}
}

func (lss *LonggestSubStr) reset() {
	if len(lss.mp) > lss.longgest {
		lss.longgest = len(lss.mp)
	}
	lss.mp = make(map[byte]bool)
}

func (lss LonggestSubStr) GetLoggest() int {
	return lss.longgest
}

func LengthOfLongestSubstring(s string) int {
	lsstr := NewLonggestSubStr(s)
	lsstr.Count()
	return lsstr.longgest
}
