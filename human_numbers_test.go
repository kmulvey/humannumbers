package humannumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Parallel()

	var total, err = Parse("three")
	assert.NoError(t, err)
	assert.InEpsilon(t, float64(3), total, 0.0001)

	total, err = Parse("seventeen")
	assert.NoError(t, err)
	assert.InEpsilon(t, float64(17), total, 0.0001)

	total, err = Parse("forty four")
	assert.NoError(t, err)
	assert.InEpsilon(t, float64(44), total, 0.0001)

	total, err = Parse("seven hundred and forty three")
	assert.NoError(t, err)
	assert.InEpsilon(t, float64(743), total, 0.0001)

	total, err = Parse("two thousand three hundred and seven")
	assert.NoError(t, err)
	assert.InEpsilon(t, float64(2307), total, 0.0001)

	total, err = Parse("negative two million")
	assert.NoError(t, err)
	assert.InEpsilon(t, float64(-2e6), total, 0.0001)

	total, err = Parse("three million eight hundred and ninety four thousand seven hundred and sixty five")
	assert.NoError(t, err)
	assert.InEpsilon(t, float64(3_894_765), total, 0.0001)
}

func TestConvertHumanStringToNumberSlice(t *testing.T) {
	t.Parallel()

	var arr, err = convertHumanStringToNumberSlice("three million eight hundred ninety four thousand seven hundred five") // the word 'and' would have been removed by Parse()
	assert.NoError(t, err)
	assert.Equal(t, []int{3, 1e6, 8, 100, 90, 4, 1000, 7, 100, 5}, arr)
}

func TestCompressNumberSliceToInt(t *testing.T) {
	t.Parallel()

	var result, err = compressNumberSliceToInt([]int{2})
	assert.NoError(t, err)
	assert.InEpsilon(t, float64(2), result, 0.0001)

	result, err = compressNumberSliceToInt([]int{17})
	assert.NoError(t, err)
	assert.InEpsilon(t, float64(17), result, 0.0001)

	result, err = compressNumberSliceToInt([]int{20})
	assert.NoError(t, err)
	assert.InEpsilon(t, float64(20), result, 0.0001)

	result, err = compressNumberSliceToInt([]int{90, 9})
	assert.NoError(t, err)
	assert.InEpsilon(t, float64(99), result, 0.0001)

	result, err = compressNumberSliceToInt([]int{100, 7})
	assert.NoError(t, err)
	assert.InEpsilon(t, float64(107), result, 0.0001)

	result, err = compressNumberSliceToInt([]int{100, 40})
	assert.NoError(t, err)
	assert.InEpsilon(t, float64(140), result, 0.0001)

	result, err = compressNumberSliceToInt([]int{2, 100, 40, 7})
	assert.NoError(t, err)
	assert.InEpsilon(t, float64(247), result, 0.0001)

	result, err = compressNumberSliceToInt([]int{7, 1000, 6})
	assert.NoError(t, err)
	assert.InEpsilon(t, float64(7006), result, 0.0001)

	result, err = compressNumberSliceToInt([]int{7, 1000, 60})
	assert.NoError(t, err)
	assert.InEpsilon(t, float64(7060), result, 0.0001)

	result, err = compressNumberSliceToInt([]int{7, 1000, 50, 5})
	assert.NoError(t, err)
	assert.InEpsilon(t, float64(7055), result, 0.0001)

	result, err = compressNumberSliceToInt([]int{2, 100, 40, 7, 1000, 6, 100, 20, 4})
	assert.NoError(t, err)
	assert.InEpsilon(t, float64(247_624), result, 0.0001)

	result, err = compressNumberSliceToInt([]int{3, 1e6, 8, 100, 90, 4, 1000, 7, 100, 60, 5})
	assert.NoError(t, err)
	assert.InEpsilon(t, float64(3_894_765), result, 0.0001)

	result, err = compressNumberSliceToInt([]int{3, 1e6, 8})
	assert.NoError(t, err)
	assert.InEpsilon(t, float64(3_000_008), result, 0.0001)

	result, err = compressNumberSliceToInt([]int{3, 100, 1e6, 8})
	assert.NoError(t, err)
	assert.InEpsilon(t, float64(300_000_008), result, 0.0001)
}

func TestFloatToSlice(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "two", floatToString(2.0))
	assert.Equal(t, "forty five", floatToString(45))
	assert.Equal(t, "one hundred twenty three dot four five six", floatToString(123.456))
	assert.Equal(t, "seven thousand one hundred twenty three dot four five six", floatToString(7_123.456))
	assert.Equal(t, "fifty seven thousand one hundred twenty three", floatToString(57_123))
	assert.Equal(t, "nine hundred eighty seven million six hundred fifty four thousand three hundred twenty one", floatToString(987_654_321))
	assert.Equal(t, "one hundred twenty three billion nine hundred eighty seven million six hundred fifty four thousand three hundred twenty one", floatToString(123_987_654_321))
	assert.Equal(t, "one hundred twenty three trillion four hundred fifty six billion nine hundred eighty seven million six hundred fifty four thousand three hundred twenty one", floatToString(123_456_987_654_321))
}

func TestSliceToInt(t *testing.T) {
	t.Parallel()

	var result = parseIntString("three")
	assert.Equal(t, 3, result)

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
