package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		assert := assert.New(t)
		rectangle := Rectangle{Width: 10.0, Length: 10.0}
		got := rectangle.Perimeter()
		want := 40.0
		assert.Equal(want, got, "got %.2f want %.2f", got, want)
	})

	t.Run("circle", func(t *testing.T) {
		assert := assert.New(t)
		circle := Circle{Radius: 10.0}
		got := circle.Perimeter()
		want := 62.83
		assert.InDelta(want, got, 0.01, "got %.2f want %.2f", got, want)
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{"Rectangle", Rectangle{Width: 12, Length: 6}, 72.0},
		{"Circle", Circle{Radius: 10}, 314.16},
		{"Triangle", Triangle{Base: 12, Height: 6}, 36},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			got := tt.shape.Area()
			assert.InDelta(tt.expected, got, 0.01, "%#v  got %.2f want %.2f", tt.shape, got, tt.expected)
		})
	}
}
