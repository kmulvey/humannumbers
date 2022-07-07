package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	var total, err = Parse("three")
	assert.NoError(t, err)
	assert.Equal(t, 3, total)

	total, err = Parse("seventeen")
	assert.NoError(t, err)
	assert.Equal(t, 17, total)

	total, err = Parse("forty four")
	assert.NoError(t, err)
	assert.Equal(t, 44, total)

	total, err = Parse("seven hundred and forty three")
	assert.NoError(t, err)
	assert.Equal(t, 743, total)

	total, err = Parse("two thousand three hundred and seven")
	assert.NoError(t, err)
	assert.Equal(t, 2307, total)
}

func TestCompressNumberSliceToInt(t *testing.T) {
	assert.Equal(t, 247_624, compressNumberSliceToInt([]int{2, 100, 40, 7, 1000, 6, 100, 20, 4}))
	assert.Equal(t, 247, compressNumberSliceToInt([]int{2, 100, 40, 7}))
	assert.Equal(t, 2, compressNumberSliceToInt([]int{2}))
}
