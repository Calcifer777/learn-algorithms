package divandconq

// Sort-and-Count( L )
// If the list has one element then
// there are no inversions
// Else
// Divide the list into two halves:
// A contains the first n/2 elements
// B contains the remaining n/2 elements
// ( r A , A ) = Sort-and-Count( A )
// ( r B , B ) = Sort-and-Count( B )
// ( r , L ) = Merge-and-Count( A, B )
// Endif
// Return r = r A + r B + r , and the sorted list L

func SortAndCount(arr []int) (sortedArr []int, numInversions int) {
	if len(arr) <= 1 {
		sortedArr = arr
		numInversions = 0
	} else {
		h1, h2 := arr[:len(arr)/2], arr[len(arr)/2:]
		h1Sorted, h1Inv := SortAndCount(h1)
		h2Sorted, h2Inv := SortAndCount(h2)
		merged, mergeInv := MergeAndCount(h1Sorted, h2Sorted)
		sortedArr = merged
		numInversions = h1Inv + h2Inv + mergeInv
	}
	return
}

// Merge-and-Count( A , B )
// Maintain a Current pointer into each list, initialized to point to the front elements
// Maintain a variable Count for the number of inversions, initialized to 0
// While both lists are nonempty:
//
//	Let a i and b j be the elements pointed to by the Current pointer
//	Append the smaller of these two to the output list
//	If b j is the smaller element then
//	  Increment Count by the number of elements remaining in A
//	Endif
//	Advance the Current pointer in the list from which the
//	smaller element was selected.
//
// EndWhile
// Once one list is empty, append the remainder of the other list to the output
// Return Count and the merged list
func MergeAndCount(xs, ys []int) (merged []int, invCount int) {
	var looper func(xs, ys, mergedIn []int, invCountIn int) (mergedOut []int, invCountOut int)
	looper = func(xs, ys, mergedIn []int, invCountIn int) (mergedOut []int, invCountOut int) {
		if len(xs) == 0 {
			mergedOut = append(mergedIn, ys...)
			invCountOut = invCountIn
		} else if len(ys) == 0 {
			mergedOut = append(mergedIn, xs...)
			invCountOut = invCountIn
		} else {
			if xs[0] < ys[0] {
				mergedOut, invCountOut = looper(xs[1:], ys, append(mergedIn, xs[0]), invCountIn)
			} else {
				mergedOut, invCountOut = looper(xs, ys[1:], append(mergedIn, ys[0]), invCountIn+len(xs))
			}
		}
		return
	}
	merged, invCount = looper(xs, ys, make([]int, 0), 0)
	return
}
