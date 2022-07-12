package humannumbers

import (
	"fmt"
	"strings"
)

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
func Parse(humanString string) (float64, error) {
	// some linting
	humanString = strings.ToLower(humanString)
	humanString = strings.ReplaceAll(humanString, " and ", " ")
	// handle negatives
	var negative = strings.Contains(humanString, "negative")
	humanString = strings.ReplaceAll(humanString, "negative", " ")

	// handle decimals
	var base = humanString
	var decimal float64
	var err error
	if strings.Contains(humanString, "point") {
		var arr = strings.Split(base, "point")
		base = arr[0]
		decimal, err = handleDecimals(arr[1])
		if err != nil {
			return 0, err
		}
	}

	baseArr, err := convertHumanStringToNumberSlice(base)
	if err != nil {
		return 0, err
	}

	baseTotal, err := compressNumberSliceToInt(baseArr)
	if err != nil {
		return 0, err
	}

	if decimal != 0.0 {
		baseTotal += decimal
	}
	if negative {
		baseTotal *= -1
	}

	return baseTotal, nil
}

// handleDecimals is pretty simple, due to the language, it just
// smashes the digits together behind the decimal point
func handleDecimals(humanString string) (float64, error) {
	var decimalArr, err = convertHumanStringToNumberSlice(humanString)
	if err != nil {
		return 0, err
	}

	var total float64
	var multiplier = 0.1
	for _, digit := range decimalArr {
		total += float64(digit) * multiplier
		multiplier *= .10
	}
	return total, nil
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
			return nil, fmt.Errorf("unknown word '%s'", word)
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
func compressNumberSliceToInt(numbers []int) (float64, error) {

	if len(numbers) == 1 {
		return float64(numbers[0]), nil
	}

	// calculate decades
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] >= 20 && numbers[i] <= 90 {
			if numbers[i+1] > 0 && numbers[i+1] < 10 {
				numbers[i] = numbers[i] + numbers[i+1]
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
				i -= 1
			}
		}
		if i < len(numbers)-1 && numbers[i] >= 100 && numbers[i] < 1000 {
			if numbers[i+1] > 0 && numbers[i+1] < 100 {
				numbers[i] = numbers[i] + numbers[i+1]
				numbers = remove(numbers, i+1)
			}
		}
	}

	// calculate large multiples i.e. the numbers before the large one
	for i := 0; i < len(numbers); i++ {
		if i > 0 && numbers[i] >= 1000 && numbers[i-1] < 1000 {
			numbers[i] = numbers[i] * numbers[i-1]
			numbers = remove(numbers, i-1)
			i -= 1
		}
	}

	// add on larger number specificity i.e. the numbers after the large one
	for i := len(numbers) - 1; i > 0; i-- {
		if numbers[i] < numbers[i-1] {
			numbers[i-1] += numbers[i]
			numbers = remove(numbers, i)
		}
	}

	if len(numbers) != 1 {
		return 0.0, fmt.Errorf("number array was no fully reduced: %+v", numbers)
	}

	return float64(numbers[0]), nil
}

// floatToString is a work in progress, its intention is to turn floats into human text
func floatToString(number float64) string {
	var numStr = fmt.Sprint(number)
	var decimalIndex int // this helps us know when we have gotten to the base and thus need to start multiplying
	var multiple = 10
	var wordsArr []string

	// in case there is no '.', we need to get ourselves into the
	// first if() in the loop
	if !strings.Contains(numStr, ".") {
		decimalIndex = len(numStr)
	}

	for i := len(numStr) - 1; i >= 0; i-- {
		if i < decimalIndex-1 {
			decade, has := decadesReverse[int(numStr[i]-'0')*multiple]
			if has {
				wordsArr = append([]string{decade}, wordsArr...)
			}
			var largeMag = largeMagToString(int(numStr[i]-'0') * multiple)
			if largeMag != "" {
				wordsArr = append([]string{baseReverse[int(numStr[i]-'0')], largeMag}, wordsArr...)
			}
			multiple *= 10
		} else if numStr[i] != 46 { // ascii: .
			wordsArr = append([]string{baseReverse[int(numStr[i]-'0')]}, wordsArr...)
		} else {
			decimalIndex = i
			wordsArr = append([]string{"dot"}, wordsArr...)
		}
	}

	return strings.Join(wordsArr, " ")
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
