package main

// CountWordOccurences takes a slice of `words`, and a `word` string.
// Return a number of occurences that `word` takes place in `words`.
func CountWordOccurences(words []string, word string) int {
	result := 0
	for _, wordfromslice := range words{
		if wordfromslice == word{
			result++
		}
	}
	
	return result
}