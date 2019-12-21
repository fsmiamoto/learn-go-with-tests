package main

// Sum return the sum of a given slice of integers
func Sum(numbers []int) int {
	sum := 0

	for _, n := range numbers {
		sum += n
	}

	return sum
}

// SumAll returns the a slice with the sum of each index of the
// given slices
func SumAll(numbers ...[]int) []int {
	var sums []int

	for _, s := range numbers {
		sums = append(sums, Sum(s))
	}

	return sums
}

// SumAllFaster has the same functionality as SumAll
// but a faster (and harder to read) implementation
func SumAllFaster(numbers ...[]int) []int {
	length := len(numbers)
	sums := make([]int, length)

	for i := range numbers {
		sums[i] = Sum(numbers[i])
	}

	return sums
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int

	for _, s := range numbers {
		if len(s) == 0 {
			sums = append(sums, 0)
			continue
		}

		tail := s[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
