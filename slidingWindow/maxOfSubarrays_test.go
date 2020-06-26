package slidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintMaxOfSubarrays(t *testing.T) {

	assert := assert.New(t)
	q := []int{1, 2, 3, 1, 4, 5, 2, 3, 6}
	exp := []int{3, 3, 4, 5, 5, 5, 6}
	o := printMaxOfSubarrays(q, 3)

	assert.Equal(exp, o)

	q = []int{8, 5, 10, 7, 9, 4, 15, 12, 90, 13}
	exp = []int{10, 10, 10, 15, 15, 90, 90}

	o = printMaxOfSubarrays(q, 4)

	assert.Equal(exp, o)
}
