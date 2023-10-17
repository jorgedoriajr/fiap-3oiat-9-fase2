package slice

// Contains checks if the slice contains the given value
func Contains[T comparable](c []T, s T) bool {
	for _, v := range c {
		if v == s {
			return true
		}
	}
	return false
}

// Any checks if predicate is true for any element
func Any[T any](c []T, predicate func(T) bool) bool {
	for _, v := range c {
		if predicate(v) {
			return true
		}
	}
	return false
}

/*
Map returns a new slice with transformed items

Usage:

	numbers := []string{1, 2, 3, 4}
	numbersInStr := Map(numbers, func(n int) string {
		return string(n)
	})
*/
func Map[T any, M any](oldSlice []T, mapFunc func(T) M) []M {
	newSlice := make([]M, len(oldSlice))
	for i, e := range oldSlice {
		newSlice[i] = mapFunc(e)
	}
	return newSlice
}

/*
ToMap returns a map[MK]MV using the given slice and mapping functions. MK must be a comparable.

Usage:

	input := []string{"value", "another value"}

	m := slice.ToMap(
		input,
		func(s string) string { return s },
		func(s string) int { return len(s) },
	)

output:

	m[value] -> 5
	m[another value] -> 13
*/
func ToMap[T any, MK comparable, MV any](inputSlice []T, keyfunc func(T) MK, valueFunc func(T) MV) map[MK]MV {
	outputMap := make(map[MK]MV, 0)
	for _, v := range inputSlice {
		outputMap[keyfunc(v)] = valueFunc(v)
	}

	return outputMap
}
