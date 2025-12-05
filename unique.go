package collections

// UniqueUsingMap returns a new slice with only unique items using the map
// function
func UniqueUsingMap[T comparable](source []T) []T {
	return MapUnique(source, func(s T) T { return s })
}

// UniqueInline inlines the code from the mpa function
func UniqueInline[T comparable](source []T) []T {
	m := make(map[T]struct{})
	n := make([]T, 0, len(source))
	for _, v := range source {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			n = append(n, v)
		}
	}
	return n
}

// UniqueInPlace returns a new slice with only unique items, but modifies the source
// slice in place. This is kept here for reference, since it is slower than the
// Unique implementation.
func UniqueInPlace[T comparable](source []T) []T {
	// Number of unique items
	l := 0
	m := make(map[T]struct{})

	for _, v := range source {
		if _, ok := m[v]; !ok {
			source[l] = v
			l++
			m[v] = struct{}{}
		}
	}
	return source[:l]
}

// UniqueFilter returns a new slice with only unique items using the filter
// function
func UniqueFilter[S comparable](source []S) []S {
	return Filter(source, func(agg []S, s S) bool {
		return !Contains(s, agg)
	})
}
