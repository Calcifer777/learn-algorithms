package divandconq

import "github.com/stretchr/testify/assert"
import "testing"

func TestMergeSort(t *testing.T) {
	arr := []int{1, 3, 5, 2, 6, 3, 7}
	less := func(i1, i2 int) bool { return i1 < i2 }
	out := MergeSort(arr, less)
	expected := []int{1, 2, 3, 3, 5, 6, 7}
	assert.Equal(t, expected, out)
}
