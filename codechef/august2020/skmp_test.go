package august2020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSKMP(t *testing.T) {
	assert := assert.New(t)

	o := findSmallestCombination("akramkeeanany", "aka")
	assert.Equal("aaakaeekmnnry", o)

	o = findSmallestCombination("supahotboy", "bohoty")
	assert.Equal("abohotypsu", o)

	o = findSmallestCombination("daehabshatorawy", "badawy")
	assert.Equal("aabadawyehhorst", o)

	o = findSmallestCombination("daehabshatorawy", "sha")
	assert.Equal("aaashabdehortwy", o)
}
