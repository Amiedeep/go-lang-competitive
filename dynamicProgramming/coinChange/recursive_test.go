package coinchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoinChangeRecursive(t *testing.T) {

	assert := assert.New(t)
	input := []int{1, 2, 3}
	exp := 4
	o := findWaysRecursive(input, 4, 0)

	assert.Equal(exp, o)
}
