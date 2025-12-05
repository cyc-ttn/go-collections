package collections

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Unique2 returns a new slice with only unique items, but modifies the source
// slice in place. This is kept here for reference, since it is slower than the
// Unique implementation.
func Unique2[T comparable](source []T) []T {
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

func TestUnique(t *testing.T) {
	x := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		x = append(x, i%10)
	}

	y := Unique(x)
	assert.Len(t, y, 10)
}

func TestUnique2(t *testing.T) {
	x := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		x = append(x, i%10)
	}

	y := Unique2(x)
	assert.Len(t, y, 10)
}

func BenchmarkUnique(b *testing.B) {
	x := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		r := rand.Intn(10)
		x = append(x, r)
	}

	b.Run("Unique", func(b *testing.B) {
		y := make([]int, 100)
		copy(y, x)

		for b.Loop() {
			Unique(y)
		}
	})

	b.Run("Unique2", func(b *testing.B) {
		y := make([]int, 100)
		copy(y, x)

		for b.Loop() {
			Unique2(y)
		}
	})
}
