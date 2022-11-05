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
		assert.Equal(want, got, "got %d want %d given, %v", got, want, numbers)
	})

}

func TestSumAll(t *testing.T) {
	t.Run("", func(t *testing.T) {
		assert := assert.New(t)
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		assert.ElementsMatch(want, got, "got %v want %v", got, want)
	})

}

func TestSumTailsTable(t *testing.T) {
	tests := []struct {
		name      string
		nums      [][]int
		expect    []int
		assertion assert.ComparisonAssertionFunc
	}{
		{"3 arrays 2 length", [][]int{{1, 2}, {3, 4}, {5, 6}}, []int{2, 4, 6}, assert.Equal},
		{"2 arrays 3 length", [][]int{{1, 2, 3}, {3, 4, 5}}, []int{5, 9}, assert.Equal},
		{"Empty slices", [][]int{{}, {}}, []int{0, 0}, assert.Equal},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, tt.expect, SumAllTails(tt.nums...))
		})
	}
}
