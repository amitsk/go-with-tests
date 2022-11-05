package arrayslices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumTable(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		expect    int
		assertion assert.ComparisonAssertionFunc
	}{
		{"[1,2]", []int{1, 2, 3, 4, 5}, 15, assert.Equal},
		{"[1,2,3,4,5]", []int{1, 2}, 3, assert.Equal},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, tt.expect, Sum(tt.nums))
		})
	}
}

func TestSum(t *testing.T) {
	t.Run("", func(t *testing.T) {
		assert := assert.New(t)
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assert.Equal(got, want, "got %d want %d given, %v", got, want, numbers)
	})

}
