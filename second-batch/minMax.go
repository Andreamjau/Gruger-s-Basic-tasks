package main

// MinMax takes a slice of integers and returns four values:
// - the minimum value
// - the maximum value
// - the second minimum value
// - the second maximum value
//
// Example:
//   MinMax([]int{12, 5, 7, 19, 3}) => (3, 19, 5, 12)

func MinMax(nums []int) (int, int, int, int) {
    if len(nums) == 0 {
        return 0, 0, 0, 0
    }

    if len(nums) == 1 {
        return nums[0], nums[0], nums[0], nums[0]
    }

    min := 99999
    nextmin := 99999
    max := -99999
    nextmax := -99999

    for _, num := range nums {
        if num < min{
            nextmin = min
            min = num
        } else if num < nextmin && num != min{
            nextmin = num 
        }

        if num > max {
            nextmax = max
            max = num 
        } else if num > nextmax && num != max{
            nextmax = max
        }
    }
    return min, max, nextmin, nextmax
}

/* this one is flawed 
func MinMax(nums []int) (int, int, int, int) {
    min := nums[0]
    max := nums[0]
    nextmin := nums[0]
    nextmax := nums[0]

    for _, num := range nums{
        switch{
        case num > max:
            nextmax = max
            max = num

        case num > nextmax && num != max || max == nextmax:
            nextmax = num

        case num < nextmin && num > min || min == nextmin:
            nextmin = num

        case num < min:
            nextmin = min
            min = num
        }
    }
    return min, max, nextmin, nextmax
}
    */
