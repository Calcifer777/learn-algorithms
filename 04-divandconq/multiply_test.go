package divandconq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKaratsuba(t *testing.T) {
	x := 36
	y := 11
	exp := 396
	out := Karatsuba(x, y)
	assert.Equal(t, exp, out)
}
