package humannumbers

import "strings"

// Parse takes an english string containing numbers in the form
// of words and converts it to float64.
// Examples: forty three, two hundred and forty six thousand three hundred and eighty seven.
func ParseFloat(humanString string) (float64, error) {
	// some linting
	humanString = strings.ToLower(humanString)
	humanString = strings.ReplaceAll(humanString, " and ", " ")

	if err := validateInput(humanString); err != nil {
		return 0, err
	}

	// Extract negative flag early
	var negative = strings.Contains(humanString, "negative")
	humanString = strings.ReplaceAll(humanString, "negative", " ")

	var integerPart string
	var fractionalPart string

	// Find decimal point
	fields := strings.Fields(humanString)
	foundDecimal := false
	for i, word := range fields {
		if word == "point" || word == "dot" {
			// Join fields before decimal
			integerPart = strings.Join(fields[:i], " ")
			// Join fields after decimal
			if i+1 < len(fields) {
				fractionalPart = strings.Join(fields[i+1:], " ")
			}
			foundDecimal = true
			break
		}
	}

	// If no decimal point, the whole string is the integer part
	if !foundDecimal {
		integerPart = strings.Join(fields, " ")
	}

	var total = float64(parseIntString(integerPart))

	if fractionalPart != "" {
		decimals, err := handleDecimals(fractionalPart)
		if err != nil {
			return 0, err
		}
		total += decimals
	}

	if negative {
		total *= -1
	}

	return total, nil
}

// handleDecimals converts the fractional part of a human-readable number string
func handleDecimals(humanString string) (float64, error) {
	var total float64
	var multiplier = 0.1
	var humanArr = strings.Fields(humanString)

	for _, word := range humanArr {
		if num, has := baseNumbers[word]; has {
			total += float64(num) * multiplier // Always add, even when num is 0
			multiplier *= .10
		}
	}

	return total, nil
}
