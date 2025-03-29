package main

//ReverseSlice returns reversed slice of provided items.
func ReverseSlice(items []int) []int {
	reversedSlice := []int{}
  
	if len(items) == 0 {
	  return items
	}
  
	for i := 0; i < len(items); i++{
	  reversedSlice = append(reversedSlice, items[len(items)-i-1])
	}
  
	return reversedSlice
  }
  
  