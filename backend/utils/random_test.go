package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewRandom(t *testing.T) {
	rand1 := NewRandom()
	rand2 := NewRandom()

	num1 := rand1.Number(1, 100)
	num2 := rand2.Number(1, 100)

	assert.NotEqual(t, num1, num2, "random number should be different")

	str1 := rand1.String(10)
	str2 := rand2.String(10)

	assert.NotEqual(t, str1, str2, "random string should be different")
}
