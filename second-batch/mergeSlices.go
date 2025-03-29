package main

// MergeSlices takes two slices of integers called `x` and `y`.
// Returns a slice of integers consisting of both `x` and `y`.
func MergeSlices(x []int, y []int) []int {
	z := []int{}

	z = append(z, x...)
	z = append(z, y...)
	return z
}
