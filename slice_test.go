package collections

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueUsingMap(t *testing.T) {
	x := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		x = append(x, i%10)
	}

	y := UniqueUsingMap(x)
	assert.Len(t, y, 10)
}

func TestUniqueInPlace(t *testing.T) {
	x := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		x = append(x, i%10)
	}

	y := UniqueInPlace(x)
	assert.Len(t, y, 10)
}

func TestUniqueInline(t *testing.T) {
	x := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		x = append(x, i%10)
	}

	y := UniqueInline(x)
	assert.Len(t, y, 10)
}

func TestUniqueFilter(t *testing.T) {
	x := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		x = append(x, i%10)
	}

	y := UniqueFilter(x)
	assert.Len(t, y, 10)
}

func BenchmarkUnique(b *testing.B) {
	for i := 0; i < 10; i++ {
		numEls := 100 * (i + 1)

		x := make([]int, 0, numEls)
		for i := 0; i < numEls; i++ {
			r := rand.Intn(numEls / 2)
			x = append(x, r)
		}

		b.Run(fmt.Sprintf("For %d items", numEls), func(b *testing.B) {
			b.Run("Unique", func(b *testing.B) {
				y := make([]int, numEls)
				copy(y, x)

				for b.Loop() {
					UniqueUsingMap(y)
				}
			})

			b.Run("UniqueInPlace", func(b *testing.B) {
				y := make([]int, numEls)
				copy(y, x)

				for b.Loop() {
					UniqueInPlace(y)
				}
			})

			b.Run("UniqueInline", func(b *testing.B) {
				y := make([]int, numEls)
				copy(y, x)

				for b.Loop() {
					UniqueInline(y)
				}
			})

			b.Run("UniqueFilter", func(b *testing.B) {
				y := make([]int, numEls)
				copy(y, x)

				for b.Loop() {
					UniqueFilter(y)
				}
			})
		})
	}
}
