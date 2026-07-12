package humannumbers

// nolint: gochecknoglobals
var baseNumbers = map[string]int{
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
	"eighteen":  18,
	"nineteen":  19,
	"twenty":    20,
	"thirty":    30,
	"forty":     40,
	"fifty":     50,
	"sixty":     60,
	"seventy":   70,
	"eighty":    80,
	"ninety":    90,
}

var hundred = "hundred"

// nolint: gochecknoglobals
var largeMagnitudes = map[string]int{
	"hundred":     100,
	"thousand":    1000,
	"million":     1e6,
	"billion":     1e9,
	"trillion":    1e12,
	"quadrillion": 1e15,
	"quintillion": 1e18,
}

// largeMagToString is a convience func to work the above map
// maybe we no longer need the map?
// nolint: gochecknoglobals
func largeMagToString(number int) string {
	switch {
	case number >= 100 && number < 1000:
		return "hundred"
	case number >= 1000 && number < 1e6:
		return "thousand"
	case number >= 1e6 && number < 1e9:
		return "million"
	case number >= 1e9 && number < 1e12:
		return "billion"
	case number >= 1e12 && number < 1e15:
		return "trillion"
	case number >= 1e15 && number < 1e18:
		return "quadrillion"
	}
	return ""
}
