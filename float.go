package humannumbers

import "strings"

// Parse takes a string containing numbers in the form
// of words, currently only English, and converts it
// to float64. Examples:
// forty three
// two hundred and forty six thousand three hundred and eighty seven.
func ParseFloat(humanString string) (float64, error) {
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
