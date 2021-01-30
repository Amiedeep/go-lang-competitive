package jan2021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	assert := assert.New(t)

	chars := []byte{ 'a', 'b','c','d','e','f','g','h','i','j','k','l','m','n','o','p' }

	exp := "a"
	o := decode("0000", chars)
	assert.Equal(exp, o)

	
	exp = "ap"
	o = decode("00001111", chars)
	assert.Equal(exp, o)

	exp = "j"
	o = decode("1001", chars)
	assert.Equal(exp, o)
}