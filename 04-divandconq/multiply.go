package divandconq

import (
	"fmt"
	"strconv"
)

// Karatsuba(x,y):
//
//	Write x = x1 · 2 n/2 + x0
//	y = y1 · 2 n/2 + y0
//	Compute x1 + x 0 and y 1 + y 0
//	p = Karatsuba( x 1 + x 0 , y 1 + y 0 )
//	x1 y1 = Karatsuba( x1 , y1 )
//	x0 y0 = Karatsuba( x0 , y0 )
//	Return x1y1 · 2n + (p − x1y1 − x0y0 ) · 2 n/2 + x0y0
func Karatsuba(x, y int) (multiplication int) {
	if x < 10 || y < 10 {
		multiplication = x * y
	} else if x < y {
		multiplication = Karatsuba(y, x)
	} else {
		// Convert to bits, pad to same length
		xs := fmt.Sprintf("%b", x)
		yTmpl := fmt.Sprintf("%%0%db", len(xs))
		ys := fmt.Sprintf(yTmpl, y)
		splitBit := len(xs) / 2
		x1s, x0s := xs[:splitBit], xs[splitBit:]
		y1s, y0s := ys[:splitBit], ys[splitBit:]
		x0, _ := strconv.ParseInt(x0s, 2, 64)
		x1, _ := strconv.ParseInt(x1s, 2, 64)
		y0, _ := strconv.ParseInt(y0s, 2, 64)
		y1, _ := strconv.ParseInt(y1s, 2, 64)

		p := Karatsuba(int(x0+x1), int(y0+y1))
		x1y1 := Karatsuba(int(x1), int(y1))
		x0y0 := Karatsuba(int(x0), int(y0))
		multiplication = (x1y1 << len(xs)) + ((p - x1y1 - x0y0) << splitBit) + x0y0
	}
	return
}

func SplitBits(x int, idx int) (int, int) {
	// bitCount := len(fmt.Sprintf("%b", x))
	l := x >> idx
	r := ((1 << idx) - 1) & x
	return l, r
}
