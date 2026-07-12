package humannumbers

import (
	"errors"
	"fmt"
	"strings"
)

var (
	errDecimalNotSupported = errors.New("decimal numbers not supported in Parse, use ParseFloat instead")
	errEmptyInput          = errors.New("input string is empty")
	errInvalidWord         = errors.New("invalid word in input")
)

// Parse takes an english string containing numbers in the form of words and converts it to an integer.
// Examples: forty three, two hundred and forty six thousand three hundred and eighty seven.
func Parse(humanString string) (int, error) {
	// some linting
	humanString = strings.ToLower(humanString)
	humanString = strings.ReplaceAll(humanString, " and ", " ")

	if err := validateInput(humanString); err != nil {
		return 0, err
	}

	if strings.Contains(humanString, "point") || strings.Contains(humanString, "dot") {
		return 0, errDecimalNotSupported
	}

	return parseIntString(humanString), nil
}

// parseIntString converts a human-readable number string into an integer.
func parseIntString(humanString string) int {
	var total int
	var section int
	var humanArr = strings.Fields(humanString)

	var largeMagnitudesIndicies = make(map[int]struct{})
	for i, word := range humanArr {
		if _, has := largeMagnitudes[word]; has {
			largeMagnitudesIndicies[i] = struct{}{}
		}
	}

	for i, word := range humanArr {
		if num, has := baseNumbers[word]; has {
			section += num
		} else if word == hundred {
			section *= 100
		} else if num, has := largeMagnitudes[word]; has {
			section *= num
			if _, has := largeMagnitudesIndicies[i]; has {
				total += section
				section = 0
			}
		}
	}
	total += section

	if strings.HasPrefix(humanString, "negative") {
		total *= -1
	}
	return total
}

// validateInput checks if the input string contains only valid number words.
func validateInput(humanString string) error {
	if humanString == "" {
		return errEmptyInput
	}

	for word := range strings.FieldsSeq(humanString) {
		if _, has := baseNumbers[word]; !has {
			if _, has := largeMagnitudes[word]; !has {
				if word != hundred && word != "point" && word != "dot" && word != "negative" {
					return fmt.Errorf("%w: %s", errInvalidWord, word)
				}
			}
		}
	}

	return nil
}
