package greedy

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHuffman(t *testing.T) {
	freqs := map[string]int{
		"a": 32,
		"b": 25,
		"c": 20,
		"d": 18,
		"e": 05,
	}
	out := Huffman(freqs)
	exp := `(
		v: cedba, 
		l: (
			v: ced, 
			l: (
				v: c, 
				l: nil, 
				r: nil
			), 
			r: (
				v: ed, 
				l: (
					v: e, 
					l: nil, 
					r: nil
				), 
				r: (
					v: d, 
					l: nil, 
					r: nil
				)
			)
		), 
		r: (
			v: ba, 
			l: (
				v: b, 
				l: nil, 
				r: nil
			), 
			r: (
				v: a, 
				l: nil, 
				r: nil
			)
		)
	)`
	exp = strings.ReplaceAll(exp, "\n", "")
	exp = strings.ReplaceAll(exp, "\t", "")
	assert.Equal(t, exp, out.String())
}
