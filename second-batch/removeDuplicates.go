package main

// RemoveDuplicates returns provided slice of strings with removed duplicates.
func RemoveDuplicates(items []string) []string {
	tidyItems := []string{}
	
  
  
	// Looping though items
	for _, item := range items{
	  var isDuplicate bool = false
	  if len(tidyItems) == 0{
		  tidyItems = append(tidyItems, item)
		  continue
	  }
  
	  for _, tidyItem := range tidyItems{
  
		  if item == tidyItem {
			  isDuplicate = true
			  continue
		  } 
	  } 
	  if !isDuplicate {
		  tidyItems = append(tidyItems, item)
	  }
	}
  
	return tidyItems
  }
  