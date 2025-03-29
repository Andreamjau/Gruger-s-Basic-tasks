package main

// Fibonacci generates the first n numbers of the Fibonacci
// sequence and returns them as a slice of integers.
// The sequence starts with 0, 1, and each number after
// that is the sum of the two before.
//
// Example:
//
//	Fibonacci(5) => []int{0, 1, 1, 2, 3}
//	Fibonacci(1) => []int{0}
//	Fibonacci(0) => []int{}
func Fibonacci(n int) []int {

    if n == 0 {
        return []int{}
    }

    if n == 1 {
        return []int{0}
    }
    sequence := make([]int, n)

	sequence[0] = 0
	sequence[1] = 1

	for i := 2; i < n; i++ {
		sequence[i] = sequence[i-1] + sequence[i-2]
	}
    return sequence
}