package collections

import (
	"net/url"
	"strconv"

	"golang.org/x/exp/constraints"
)

// ToStringSliceOf converts a slice from one type to another. Both types must have an underlying type of `string. This
// is useful when converting between enums of different types.
//
// TODO: Is there a way to write this generically for slices that do not have an underlying type of string?
func ToStringSliceOf[T ~string, S ~string](in []S) []T {
	m := make([]T, 0, len(in))
	for _, i := range in {
		m = append(m, T(i))
	}
	return m
}

// SelectFromUrlValues provides a way to retrieve a slice of strings from a map[string][]string in a certain type whose
// underlying type is string.
func SelectFromUrlValues[T ~string](vals url.Values, key string) []T {
	return ToStringSliceOf[T](vals[key])
}

// GetFromUrlValues returns the *first* value from a map[string][]string in a certain type whose underlying type is a
// string
func GetFromUrlValues[T ~string](vals url.Values, key string) (T, bool) {
	x, ok := vals[key]
	if !ok || len(x) == 0 {
		var t T
		return t, false
	}
	return T(x[0]), true
}

// GetIntFromUrlValues returns a single integer from a map[string][]string by parsing the list of entries in order
// and returning the first integer entry.
func GetIntFromUrlValues[T constraints.Integer](vals url.Values, key string) (T, bool) {
	strs := SelectFromUrlValues[string](vals, key)
	for _, s := range strs {
		i, err := strconv.ParseInt(s, 10, 64)
		if err == nil {
			return T(i), true
		}
	}
	var t T
	return t, false
}
