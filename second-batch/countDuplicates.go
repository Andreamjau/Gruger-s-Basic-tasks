package main

// CountDuplicates takes a slice of integers and returns a map[int]int
// containing only the elements that appear more than once,
// along with how many times they appear.
//
// Example:
//   CountDuplicates([]int{1, 2, 2, 3, 4, 4, 4}) => map[2:2 4:3]
func CountDuplicates(nums []int) map[int]int {
    var results = map[int]int{}

    for _, num := range nums {
        results[num]++  
    }
    /*
    var duplicates = map[int]int{}

    for key, value := range results{
        if value > 1 {
            duplicates[key] = value
        }
    } */
    for key, value := range results {
        if value == 1 {
          delete(results, key) // another first-class function, it takes any `map` type and key that we want to delete from it.
        }
     }
    return results
}