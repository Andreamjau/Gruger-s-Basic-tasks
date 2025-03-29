package main

// SplitEvenOdd takes a slice of integers and returns two slices:
// - the first contains all even numbers
// - the second contains all odd numbers
// Order is preserved from the original slice.
//
// Example:
//   SplitEvenOdd([]int{1, 2, 3, 4, 5}) => ([]int{2, 4}, []int{1, 3, 5})
func SplitEvenOdd(nums []int) ([]int, []int) {
    var odds []int
    var evens []int

    for _, num := range nums{
        if num%2 == 0{
            evens = append(evens, num)
            
        } else {
            odds = append(odds, num)
        }
    }
    return evens, odds
}
