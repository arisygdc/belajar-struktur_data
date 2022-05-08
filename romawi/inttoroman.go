package romawi

func IntToRoman(num int) string {
	var str string
	if h := num / 1000; h > 0 {
		str = str + "M" + IntToRoman(num-1000)
	} else if num/900 > 0 {
		str = str + "CM" + IntToRoman(num-900)
	} else if num/500 > 0 {
		str = str + "D" + IntToRoman(num-500)
	} else if num/400 > 0 {
		str = str + "CD" + IntToRoman(num-400)
	} else if num/100 > 0 {
		str = str + "C" + IntToRoman(num-100)
	} else if num/90 > 0 {
		str = str + "XC" + IntToRoman(num-90)
	} else if num/50 > 0 {
		str = str + "L" + IntToRoman(num-50)
	} else if num/40 > 0 {
		str = str + "XL" + IntToRoman(num-40)
	} else if num/10 > 0 {
		str = str + "X" + IntToRoman(num-10)
	} else if num/9 > 0 {
		str = str + "IX" + IntToRoman(num-9)
	} else if num/5 > 0 {
		str = str + "V" + IntToRoman(num-5)
	} else if num/4 > 0 {
		str = str + "IV" + IntToRoman(num-4)
	} else if num/1 > 0 {
		str = str + "I" + IntToRoman(num-1)
	}
	return str
}
