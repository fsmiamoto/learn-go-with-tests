package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 20, 10}

		got := Sum(numbers)
		expect := 36

		if got != expect {
			t.Errorf("got %d expected %d given %v", got, expect, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// DeepEqual is not type safe!
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	// Adding a helper wraps the type unsafety of DeepEqual
	check := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		check(t, got, want)

	})

	t.Run("safely sum a empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		check(t, got, want)
	})
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2, 3, 4, 5, 6}, []int{0, 9, 3, 5, 2, 8})
	}
}

func BenchmarkSumAllFaster(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllFaster([]int{1, 2, 3, 4, 5, 6}, []int{0, 9, 3, 5, 2, 8})
	}
}
