package main

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	tests := []struct {
		name     string
		shape    Shape
		hasPerim float64
	}{
		{name: "Rectangle", shape: &Rectangle{Width: 12, Height: 6}, hasPerim: 36},
		{name: "Circle", shape: &Circle{Radius: 10}, hasPerim: 2 * math.Pi * 10},
		{name: "Triangle", shape: &Triangle{Base: 4, Height: 3}, hasPerim: 12},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Perimeter()
			if got != test.hasPerim {
				t.Errorf("%#v got %.2f want %.2f", test.shape, got, test.hasPerim)
			}
		})
	}
}

func TestArea(t *testing.T) {
	tests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: &Rectangle{Width: 12, Height: 6}, hasArea: 72},
		{name: "Circle", shape: &Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: &Triangle{Base: 12, Height: 6}, hasArea: 36},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.hasArea {
				t.Errorf("%#v got %.2f want %.2f", test.shape, got, test.hasArea)
			}
		})
	}
}

// Benchmark the two different implementations of Area

func BenchmarkAreaPow(b *testing.B) {
	c := Circle{5.0}
	for i := 0; i < b.N; i++ {
		c.AreaPow()
	}
}

func BenchmarkArea(b *testing.B) {
	c := Circle{5.0}
	for i := 0; i < b.N; i++ {
		c.Area()
	}
}
