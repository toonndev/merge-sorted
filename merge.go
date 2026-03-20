package mergesorted

// reverseSlice reverses a slice in-place.
func reverseSlice(s []int) []int {
	result := make([]int, len(s))
	for i, v := range s {
		result[len(s)-1-i] = v
	}
	return result
}

// mergeTwoAscending merges two ascending-sorted slices into one ascending-sorted slice
// using two-pointer technique without any built-in sort.
func mergeTwoAscending(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}

// Merge takes three slices:
//   - collection1: descending order
//   - collection2: ascending order
//   - collection3: ascending order
//
// and returns a single slice sorted ascending.
func Merge(collection1 []int, collection2 []int, collection3 []int) []int {
	asc1 := reverseSlice(collection1)
	merged := mergeTwoAscending(asc1, collection2)
	merged = mergeTwoAscending(merged, collection3)
	return merged
}
