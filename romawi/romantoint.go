package romawi

var mapRoman map[byte]int = map[byte]int{
	73: 1,
	86: 5,
	88: 10,
	76: 50,
	67: 100,
	68: 500,
	77: 1000,
}

func RomanToInt(s string) int {
	res := 0
	var tmp int
	for i := len(s) - 1; i >= 0; i-- {
		if mapRoman[s[i]] < tmp {
			res -= mapRoman[s[i]]
		} else {
			res += mapRoman[s[i]]
		}
		tmp = mapRoman[s[i]]
	}
	return res
}
