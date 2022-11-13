package generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumIntsTable(t *testing.T) {
	tests := []struct {
		name      string
		nums      map[string]int64
		expect    int64
		assertion assert.ComparisonAssertionFunc
	}{
		{"SumInts", map[string]int64{"one": 1, "two": 2}, 3, assert.Equal},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, tt.expect, SumInts(tt.nums))
		})
	}
}

func TestSumFloatsTable(t *testing.T) {
	tests := []struct {
		name      string
		nums      map[string]float64
		expect    float64
		assertion assert.ComparisonAssertionFunc
	}{
		{"SumFloats", map[string]float64{"one": 10.0, "two": 20.0}, 30.0, assert.Equal},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, tt.expect, SumFloats(tt.nums))
		})
	}
}

func TestSumGenericsTable(t *testing.T) {
	t.Run("SumFloats", func(t *testing.T) {
		assert.Equal(t, 30.0, SumIntsOrFloats(map[string]float64{"one": 10.0, "two": 20.0}))
	})
	t.Run("SumInts", func(t *testing.T) {
		assert.Equal(t, int64(15), SumIntsOrFloats(map[string]int64{"one": 5, "two": 10}))
	})
}
