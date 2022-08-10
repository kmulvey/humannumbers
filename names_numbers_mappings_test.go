package humannumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargeMagToString(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "", largeMagToString(1))
	assert.Equal(t, "hundred", largeMagToString(100))
	assert.Equal(t, "thousand", largeMagToString(1020))
	assert.Equal(t, "million", largeMagToString(1340020))
	assert.Equal(t, "billion", largeMagToString(1340020234))
	assert.Equal(t, "trillion", largeMagToString(1340020233944))
	assert.Equal(t, "quadrillion", largeMagToString(5671340020233944))
}
