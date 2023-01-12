package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommon(t *testing.T) {
	assert.Equal(t, 2, max(1, 2))
	assert.Equal(t, 1, min(1, 2))
	assert.Equal(t, -1, minAbs(-1, 2))
	assert.Equal(t, 1, absIntCompare(-4, 3))
	assert.Equal(t, 2, max(2, -1))
}
