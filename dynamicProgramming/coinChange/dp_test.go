package coinchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoinChangeDP(t *testing.T) {

	assert := assert.New(t)
	input := []int{1, 2, 3}
	exp := 4
	o := findWaysDP(input, 4)

	assert.Equal(exp, o)
}
