package humannumbers

import (
	"errors"
	"fmt"
	"strings"
)

var (
	errUnknownWord           = errors.New("unknown word")
	errNumberArrayNotReduced = errors.New("number array was not fully reduced")
)

// Parse takes a string containing numbers in the form
// of words, currently only English, and converts it
// to float64. Examples:
// forty three
// two hundred and forty six thousand three hundred and eighty seven.
func Parse(humanString string) (int, error) {
	// some linting
	humanString = strings.ToLower(humanString)
	humanString = strings.ReplaceAll(humanString, " and ", " ")

	if err := validateInput(humanString); err != nil {
		return 0, err
	}

	if strings.Contains(humanString, "point") || strings.Contains(humanString, "dot") {
		return 0, errors.New("decimal numbers not supported in Parse, use ParseFloat instead")
	}

	return parseIntString(humanString), nil
}

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

func validateInput(humanString string) error {
	if humanString == "" {
		return fmt.Errorf("input string is empty")
	}

	for _, word := range strings.Fields(humanString) {
		if _, has := baseNumbers[word]; !has {
			if _, has := largeMagnitudes[word]; !has {
				if word != hundred && word != "point" && word != "dot" && word != "negative" {
					return fmt.Errorf("invalid word in input: %s", word)
				}
			}
		}
	}

	return nil
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
