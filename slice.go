package collections

// Contains returns true if needle can be found in haystack for any type of slice.
func Contains[T comparable](needle T, haystack []T) bool {
	return Index(needle, haystack) > -1
}

// Index returns the index at which needle can be found in haystack for all types of slices. In the case that it cannot
// be found, it will return -1.
func Index[T comparable](needle T, haystack []T) int {
	for i, h := range haystack {
		if h == needle {
			return i
		}
	}
	return -1
}

// IndexWhere returns the first index at which the provided [where] function returns true. If it never returns true,
// -1 is returned.
func IndexWhere[T any](haystack []T, where func(t T) bool) int {
	for i, h := range haystack {
		if where(h) {
			return i
		}
	}
	return -1
}

// Map returns a new slice including the items returned by the provided function, usually a modification or a part of
// the provided source. For example, this can be used to return a list of Ids from a struct which contains an Id field.
//
// To omit the object entirely from the new list, return false in the second parameter. Thus, Map can be used to
// filter at the same time.
func Map[T any, S any](source []S, fn func(agg []T, s S) (T, bool)) []T {
	m := make([]T, 0, len(source))
	for _, s := range source {
		mapped, ok := fn(m, s)
		if ok {
			m = append(m, mapped)
		}
	}
	return m
}

// Filter returns a new slice including only items where the provided filter function returns true.
func Filter[S any](source []S, fn func(agg []S, s S) bool) []S {
	filtered := make([]S, 0, len(source))
	for _, s := range source {
		if fn(filtered, s) {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

// MapUnique is the same as Map but checks first if the converted item is unique. Note that by using MapUnique, there
// is no way to filter (the defined filter is the 'Unique' function. Unique is implemented by using the Contains
// function.
func MapUnique[T comparable, S any](source []S, fn func(s S) T) []T {
	return Map(source, func(agg []T, s S) (T, bool) {
		converted := fn(s)
		contains := Contains(converted, agg)
		return converted, !contains
	})
}
