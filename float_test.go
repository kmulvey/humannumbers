package humannumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleDecimals(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{"single_digit", "one", 0.1},
		{"multiple_digits", "one two three", 0.123},
		{"with_zeros", "zero one zero three", 0.0103},
		{"zeros_only", "zero zero zero", 0.0},
		{"four_five_six", "four five six", 0.456},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result := handleDecimals(test.input)

			if test.expected == 0.0 {
			assert.Zero(t, result)
			} else {
				assert.InEpsilon(t, test.expected, result, 0.0001)
			}
		})
	}
}

func TestParseFloat(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		// Simple integers (no decimals)
		{"single digit", "five", 5.0},
		{"double digit", "forty three", 43.0},
		{"hundred", "two hundred", 200.0},
		{"hundred with units", "two hundred five", 205.0},
		{"thousand", "one thousand", 1000.0},
		{"complex integer", "three thousand four hundred fifty two", 3452.0},

		// With "and" separator
		{"hundred with and", "one hundred and five", 105.0},
		{"thousand with and", "two thousand and five", 2005.0},
		{"complex with and", "five thousand six hundred and twenty three", 5623.0},

		// Simple decimals
		{"decimal half", "five point five", 5.5},
		{"decimal tenth", "ten point one", 10.1},
		{"decimal many digits", "forty three point one two three", 43.123},

		// Decimals with "dot"
		{"dot keyword", "twelve dot five six seven", 12.567},

		// Large numbers with decimals
		{"million with decimal", "one million five hundred thousand point five", 1500000.5},
		{"complex large decimal", "three thousand four hundred fifty two point six seven eight", 3452.678},

		// Negative numbers
		{"negative single", "negative five", -5.0},
		{"negative double", "negative forty three", -43.0},
		{"negative hundred", "negative two hundred", -200.0},
		{"negative thousand", "negative one thousand", -1000.0},

		// Negative with decimals
		{"negative decimal", "negative five point five", -5.5},
		{"negative complex decimal", "negative twelve point three four five", -12.345},
		{"negative large decimal", "negative one million point five", -1000000.5},

		// Edge cases
		{"zero", "zero", 0.0},
		{"zero decimal", "zero point five", 0.5},
		{"ten", "ten", 10.0},
		{"nineteen", "nineteen", 19.0},
		{"ninety", "ninety", 90.0},

		// Case insensitivity
		{"uppercase", "FORTY THREE", 43.0},
		{"mixed case", "FoRtY ThReE", 43.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result, err := ParseFloat(test.input)
			assert.NoError(t, err)

			if test.expected == 0.0 {
			assert.Zero(t, result)
			} else {
				assert.InEpsilon(t, test.expected, result, 0.0001)
			}
		})
	}
}

func TestParseFloatErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
	}{
		{"unknown word", "blah"},
		{"unknown in number", "five blah three"},
		{"unknown in decimal", "five point blah"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := ParseFloat(test.input)
			assert.Error(t, err)
		})
	}
}
