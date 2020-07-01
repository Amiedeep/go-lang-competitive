package geeksforgeeks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutiveSubsequence(t *testing.T) {

	assert := assert.New(t)
	q := []int{2, 6, 1, 9, 4, 5, 3}
	exp := 6
	o := longestConsecutiveSubsequence(q)

	assert.Equal(exp, o)

	q = []int{1, 9, 3, 10, 4, 20, 2}
	exp = 4

	o = longestConsecutiveSubsequence(q)

	assert.Equal(exp, o)
}
