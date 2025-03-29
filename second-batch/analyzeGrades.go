package main

// AnalyzeGrades takes a slice of test scores (0–100) and returns:
// - the average score as a float64
// - the highest score as an int
// - the letter grade corresponding to the average score
//   (A: 100-90, B: 90-80, C: 80-70, D: 70-60, F: 60-0)
//
// Example:
//   AnalyzeGrades([]int{85, 92, 78, 90}) => (86.25, 92, "B")
func AnalyzeGrades(scores []int) (float64, int, string) {
    var total int 
    var averageGrade string
    bestScore := 0
    
    // Finding average and best score
    for _, score := range scores{
        total += score

        if score > bestScore{
            bestScore = score
        }
    }
    average := float64(total)/float64(len(scores))

    switch{
    case average >= 90:
        averageGrade = "A"
    case average >= 80:
        averageGrade = "B"
    case average >= 70:
        averageGrade = "C"
    case average >= 60:
        averageGrade = "D"
    default:
        averageGrade = "F"
    }

    return average, bestScore, averageGrade
}
