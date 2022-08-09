package humannumbers

var base = map[string]int{
	"zero":      0,
	"one":       1,
	"two":       2,
	"three":     3,
	"four":      4,
	"five":      5,
	"six":       6,
	"seven":     7,
	"eight":     8,
	"nine":      9,
	"ten":       10,
	"eleven":    11,
	"twelve":    12,
	"thirteen":  13,
	"fourteen":  14,
	"fifteen":   15,
	"sixteen":   16,
	"seventeen": 17,
	"eightteen": 18,
	"nineteen":  19,
}

var baseReverse = map[int]string{
	0:  "zero",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eightteen",
	19: "nineteen",
}

var decades = map[string]int{
	"twenty":  20,
	"thirty":  30,
	"forty":   40,
	"fifty":   50,
	"sixty":   60,
	"seventy": 70,
	"eighty":  80,
	"ninety":  90,
}

var decadesReverse = map[int]string{
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

var largeMagnitudes = map[string]int{
	"thousand":    1000,
	"million":     1e6,
	"billion":     1e9,
	"trillion":    1e12,
	"quadrillion": 1e15,
	"quintillion": 1e18,
}

var largeMagnitudesReverse = map[int]string{
	100:  "hundred",
	1000: "thousand",
	1e6:  "million",
	1e9:  "billion",
	1e12: "trillion",
	1e15: "quadrillion",
	1e18: "quintillion",
}

// largeMagToString is a convience func to work the above map
// maybe we no longer need the map?
func largeMagToString(number int) string {
	switch {
	case number >= 100 && number < 1000:
		return largeMagnitudesReverse[100]
	case number >= 1000 && number < 1e6:
		return largeMagnitudesReverse[1000]
	case number >= 1e6 && number < 1e9:
		return largeMagnitudesReverse[1e6]
	case number >= 1e9 && number < 1e12:
		return largeMagnitudesReverse[1e9]
	case number >= 1e12 && number < 1e15:
		return largeMagnitudesReverse[1e12]
	case number >= 1e15 && number < 1e18:
		return largeMagnitudesReverse[1e15]
	}
	return ""
}
