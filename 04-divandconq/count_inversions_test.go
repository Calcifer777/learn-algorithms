package divandconq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeAndCount(t *testing.T) {
	arr1 := []int{1, 4, 7}
	arr2 := []int{2, 5, 8}
	outMerged, invCountOut := MergeAndCount(arr1, arr2)
	expMerged := []int{1, 2, 4, 5, 7, 8}
	expInvCount := 3
	assert.Equal(t, expMerged, outMerged)
	assert.Equal(t, expInvCount, invCountOut)
}

func TestSortAndCount(t *testing.T) {
	// The input should not contain any duplicate values
	subTests := []struct {
		input    []int
		expected int
	}{
		{[]int{3, 2, 1}, 3},
		{[]int{8, 4, 2, 1}, 6},
		{[]int{6, 3, 5, 2}, 5},
	}
	for _, subTest := range subTests {
		_, numInv := SortAndCount(subTest.input)
		assert.Equal(t, subTest.expected, numInv)
	}
}
