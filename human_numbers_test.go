package humannumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
	}{
		{"empty string", ""},
		{"unknown word", "blah"},
		{"unknown in number", "five blah three"},
		{"decimal in parse", "five point five"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := Parse(test.input)
			assert.Error(t, err)
		})
	}
}

// nolint:funlen
func TestSliceToInt(t *testing.T) {
	t.Parallel()

	var result = parseIntString("three")
	assert.Equal(t, 3, result)

	result = parseIntString("zero")
	assert.Equal(t, 0, result)

	result = parseIntString("twelve thousand one hundred twenty")
	assert.Equal(t, 12_120, result)

	// Pure hundreds
	result = parseIntString("ten")
	assert.Equal(t, 10, result)

	result = parseIntString("nineteen")
	assert.Equal(t, 19, result)

	result = parseIntString("twenty three")
	assert.Equal(t, 23, result)

	result = parseIntString("ninety")
	assert.Equal(t, 90, result)

	result = parseIntString("one hundred")
	assert.Equal(t, 100, result)

	result = parseIntString("nine hundred ninety nine")
	assert.Equal(t, 999, result)

	// Hundreds immediately before large magnitudes
	result = parseIntString("five hundred thousand")
	assert.Equal(t, 500_000, result)

	result = parseIntString("seven thousand one hundred twenty three")
	assert.Equal(t, 7_123, result)

	result = parseIntString("one million eight hundred fifty four thousand three hundred eighty six")
	assert.Equal(t, 1_854_386, result)

	result = parseIntString("nine hundred ninety nine million")
	assert.Equal(t, 999_000_000, result)

	// Gaps in magnitude groups
	result = parseIntString("one thousand one")
	assert.Equal(t, 1_001, result)

	result = parseIntString("one million one thousand")
	assert.Equal(t, 1_001_000, result)

	result = parseIntString("one million one")
	assert.Equal(t, 1_000_001, result)

	result = parseIntString("one billion one thousand one")
	assert.Equal(t, 1_000_001_001, result)

	result = parseIntString("eighteen billion four million seven thousand one hundred twenty three")
	assert.Equal(t, 18_004_007_123, result)

	// Trillion magnitude
	result = parseIntString("one trillion")
	assert.Equal(t, 1_000_000_000_000, result)

	result = parseIntString("one trillion two hundred thirty four billion five hundred sixty seven million eight hundred ninety thousand one hundred twenty three")
	assert.Equal(t, 1_234_567_890_123, result)
}

func TestSliceToIntLargeMagnitudes(t *testing.T) {
	t.Parallel()

	var result = parseIntString("one quadrillion")
	assert.Equal(t, 1_000_000_000_000_000, result)

	result = parseIntString("two quadrillion five hundred trillion")
	assert.Equal(t, 2_500_000_000_000_000, result)

	result = parseIntString("one quintillion")
	assert.Equal(t, 1_000_000_000_000_000_000, result)

	result = parseIntString("eight quintillion one hundred quadrillion")
	assert.Equal(t, 8_100_000_000_000_000_000, result)
}
