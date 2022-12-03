package collections

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
