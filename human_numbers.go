package humannumbers

import (
	"errors"
	"strings"
)

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

var decades = map[string]int{
	"twenty":  20,
	"thirty":  30,
	"forty":   40,
	"fifty":   50,
	"sixty":   60,
	"seventy": 70,
	"eighty":  80,
	"ninty":   90,
}

var largeMagnitudes = map[string]int{
	"thousand":    1000,
	"million":     1e6,
	"billion":     1e9,
	"trillion":    1e12,
	"quadrillion": 1e15,
	"quintillion": 1e18,
}

// Parse takes a string containing numbers in the form
// of words, currently only English, and converts it
// to an int. Examples:
// two
// forty three
// eight thousand
// eigth hundred and six
// one thousand six hundred and forty
// two thousand three hundred and eighty seven
// two hundred and forty six thousand three hundred and eighty seven
func Parse(humanString string) (int, error) {
	humanString = strings.ToLower(humanString)
	humanString = strings.ReplaceAll(humanString, " and ", " ")

	var numberArr, err = convertHumanStringToNumberSlice(humanString)
	if err != nil {
		return 0, err
	}

	return compressNumberSliceToInt(numberArr), nil
}

// convertHumanStringToNumberSlice loops through the give string and places the
// numeric equivelant to each word in an array of ints
// e.g. input: "two hundred and forty seven thousand six hundred and twenty four
// 		output: []int{2, 100, 40, 7, 1000, 6, 100, 20, 4}
func convertHumanStringToNumberSlice(humanString string) ([]int, error) {
	var humanArr = strings.Fields(humanString)
	var numbers = make([]int, len(humanArr))

	for i, word := range humanArr {
		if num, has := base[word]; has {
			numbers[i] = num
		} else if num, has := decades[word]; has {
			numbers[i] = num
		} else if word == "hundred" {
			numbers[i] = 100
		} else if num, has := largeMagnitudes[word]; has {
			numbers[i] = num
		} else {
			return nil, errors.New("unknown word " + word)
		}
	}
	return numbers, nil
}

// compressNumberSliceToInt takes an int slice of numbers
// and either adds or multiplies them depending on their
// placement in the slice until there is only one
// element in the slice.
// e.g. input: []int{2, 100, 40, 7, 1000, 6, 100, 20, 4}
// 		output: 247624
func compressNumberSliceToInt(numbers []int) int {
	// calculate decades
	for i, num := range numbers {
		if num >= 20 && num <= 90 {
			if numbers[i+1] > 0 && numbers[i+1] < 10 {
				numbers[i] = num + numbers[i+1]
				numbers = remove(numbers, i+1)
			}
		}
	}

	// calculate hundreds
	for i, num := range numbers {
		if num == 100 && i > 0 {
			if numbers[i-1] > 0 && numbers[i-1] < 10 {
				numbers[i] = num * numbers[i-1]
				numbers = remove(numbers, i-1)
			}
		}
	}

	// compress the array into one number
	for len(numbers) > 1 {
		if numbers[0] > numbers[1] {
			numbers[0] += numbers[1]
			numbers = remove(numbers, 1)
		} else {
			numbers[0] *= numbers[1]
			numbers = remove(numbers, 1)
		}
	}
	return numbers[0]
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
