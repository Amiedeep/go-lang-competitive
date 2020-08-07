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
	assert.Equal("aaabdehorshatwy", o)

	o = findSmallestCombination("bhjwbvejbwdj", "jv")
	assert.Equal("bbbdehjjjvww", o)

	o = findSmallestCombination("ghjwerguyinxcvagfweuj", "grf")
	assert.Equal("aceegggrfhijjnuuvwwxy", o)

	o = findSmallestCombination("ghjwerguyinxcvagfweuj", "fgg")
	assert.Equal("aceefggghijjnruuvwwxy", o)

	o = findSmallestCombination("ghjwerguyinxcvagfweuj", "geijr")
	assert.Equal("acefgeijrgghjnuuvwwxy", o)

	o = findSmallestCombination("ghjwerguyinxcvagfweuj", "ue")
	assert.Equal("acefggghijjnrueuvwwxy", o)

	o = findSmallestCombination("abcdeeuuuef", "uef")
	assert.Equal("abcdeeuefuu", o)

	o = findSmallestCombination("abcdeeuuuvvef", "uve")
	assert.Equal("abcdeefuuuvev", o)
}
