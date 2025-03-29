package main

import "fmt"

// FindIncreasingSequences checks input slice and returns all contiguous
// strictly increasing sub-sequences (length >= 2) in the order they appear.
//
// Example:
//   FindIncreasingSequences([]int{1, 2, 2, 3, 4, 1}) => [][]int{{1, 2}, {2, 3, 4}}
func FindIncreasingSequences(nums []int) [][]int {
    // Assuming I only accept posetive values i assigned prevvalue to -2
    previousValue := -2
    var sequences [][]int
    var currentSequence []int
    
    if len(nums) < 1 {
        return sequences
    }

    for _, num := range nums{
        fmt.Println(num, previousValue)
        if num == previousValue + 1{
            if len(currentSequence) == 0{
                currentSequence = append(currentSequence, previousValue)
            }
             currentSequence = append(currentSequence, num)
             
        } else if len(currentSequence) > 0 {
            sequences = append(sequences, currentSequence)
            currentSequence = []int{}
        }
        previousValue = num
    }
    if len(currentSequence) > 0 {
        sequences = append(sequences, currentSequence)
    }
    return sequences
}