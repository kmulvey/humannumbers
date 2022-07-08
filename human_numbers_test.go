package humannumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
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
	assert.Equal(t, float64(247_624), compressNumberSliceToInt([]int{2, 100, 40, 7, 1000, 6, 100, 20, 4}))
	assert.Equal(t, float64(247), compressNumberSliceToInt([]int{2, 100, 40, 7}))
	assert.Equal(t, float64(2), compressNumberSliceToInt([]int{2}))
	assert.Equal(t, float64(3_894_765), compressNumberSliceToInt([]int{3, 1e6, 8, 100, 90, 4, 1000, 7, 100, 60, 5}))
}
