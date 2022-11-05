package iteration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {

	t.Run("2 repetitions", func(t *testing.T) {
		assert := assert.New(t)
		repeated := Repeat("a", 2)
		assert.Equal(repeated, "aa")

	})

	t.Run("default repetitions", func(t *testing.T) {
		assert := assert.New(t)
		repeated := Repeat("a", 0)
		assert.Equal(repeated, "aaaaa")

	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}
