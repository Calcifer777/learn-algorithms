package divandconq

func MergeSort[T any](arr []T, less func(t1, t2 T) bool) []T {
	var looper func(arr []T) []T
	looper = func(arr []T) []T {
		if len(arr) <= 1 {
			return arr
		} else {
			ref := arr[0]
			larger := make([]T, 0)
			smaller := make([]T, 0)
			for _, t := range arr[1:] {
				if less(t, ref) {
					smaller = append(smaller, t)
				} else {
					larger = append(larger, t)
				}
			}
			return append(append(looper(smaller), ref), looper(larger)...)
		}
	}
	return looper(arr)
}
