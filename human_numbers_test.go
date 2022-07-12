package humannumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Parallel()

	var total, err = Parse("three")
	assert.NoError(t, err)
	assert.Equal(t, float64(3), total)

	total, err = Parse("seventeen")
	assert.NoError(t, err)
	assert.Equal(t, float64(17), total)

	total, err = Parse("forty four")
	assert.NoError(t, err)
	assert.Equal(t, float64(44), total)

	total, err = Parse("seven hundred and forty three")
	assert.NoError(t, err)
	assert.Equal(t, float64(743), total)

	total, err = Parse("two thousand three hundred and seven")
	assert.NoError(t, err)
	assert.Equal(t, float64(2307), total)

	total, err = Parse("negative two million")
	assert.NoError(t, err)
	assert.Equal(t, float64(-2e6), total)

	total, err = Parse("three million eight hundred and ninty four thousand seven hundred and sixty five")
	assert.NoError(t, err)
	assert.Equal(t, float64(3_894_765), total)

	total, err = Parse("negative two million six hundred thousand and five point six three five eight")
	assert.NoError(t, err)
	assert.Equal(t, -2600005.6358, total)
}

func TestConvertHumanStringToNumberSlice(t *testing.T) {
	t.Parallel()

	var arr, err = convertHumanStringToNumberSlice("three million eight hundred ninty four thousand seven hundred five") // the word 'and' would have been removed by Parse()
	assert.NoError(t, err)
	assert.Equal(t, []int{3, 1e6, 8, 100, 90, 4, 1000, 7, 100, 5}, arr)
}

func TestCompressNumberSliceToInt(t *testing.T) {
	t.Parallel()

	var result, err = compressNumberSliceToInt([]int{2})
	assert.NoError(t, err)
	assert.Equal(t, float64(2), result)

	result, err = compressNumberSliceToInt([]int{17})
	assert.NoError(t, err)
	assert.Equal(t, float64(17), result)

	result, err = compressNumberSliceToInt([]int{20})
	assert.NoError(t, err)
	assert.Equal(t, float64(20), result)

	result, err = compressNumberSliceToInt([]int{90, 9})
	assert.NoError(t, err)
	assert.Equal(t, float64(99), result)

	result, err = compressNumberSliceToInt([]int{100, 7})
	assert.NoError(t, err)
	assert.Equal(t, float64(107), result)

	result, err = compressNumberSliceToInt([]int{100, 40})
	assert.NoError(t, err)
	assert.Equal(t, float64(140), result)

	result, err = compressNumberSliceToInt([]int{2, 100, 40, 7})
	assert.NoError(t, err)
	assert.Equal(t, float64(247), result)

	result, err = compressNumberSliceToInt([]int{7, 1000, 6})
	assert.NoError(t, err)
	assert.Equal(t, float64(7006), result)

	result, err = compressNumberSliceToInt([]int{7, 1000, 60})
	assert.NoError(t, err)
	assert.Equal(t, float64(7060), result)

	result, err = compressNumberSliceToInt([]int{7, 1000, 50, 5})
	assert.NoError(t, err)
	assert.Equal(t, float64(7055), result)

	result, err = compressNumberSliceToInt([]int{2, 100, 40, 7, 1000, 6, 100, 20, 4})
	assert.NoError(t, err)
	assert.Equal(t, float64(247_624), result)

	result, err = compressNumberSliceToInt([]int{3, 1e6, 8, 100, 90, 4, 1000, 7, 100, 60, 5})
	assert.NoError(t, err)
	assert.Equal(t, float64(3_894_765), result)

	result, err = compressNumberSliceToInt([]int{3, 1e6, 8})
	assert.NoError(t, err)
	assert.Equal(t, float64(3_000_008), result)

	result, err = compressNumberSliceToInt([]int{3, 100, 1e6, 8})
	assert.NoError(t, err)
	assert.Equal(t, float64(300_000_008), result)
}

func TestFloatToSlice(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "two", floatToString(2.0))
	assert.Equal(t, "forty five", floatToString(45))
	assert.Equal(t, "one hundred twenty three dot four five six", floatToString(123.456))
	assert.Equal(t, "seven thousand one hundred twenty three dot four five six", floatToString(7123.456))
	//assert.Equal(t, "seven thousand one hundred twenty three dot four five six", floatToString(57123))
}
