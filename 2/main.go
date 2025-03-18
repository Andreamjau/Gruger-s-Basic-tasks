package main

import (
	"fmt"
)

// write a program that checks if character is a vowels or consonant (return true for (a,e,i,o,u))

func main(){
	vowels := [...]string{"a", "e", "i", "o", "u", "y", "æ", "ø", "å"}
	var input string; 
	
	fmt.Println("Please enter one letter")
	fmt.Scan(&input)

	for _, vowel := range vowels{
		if input == vowel{
			fmt.Println("Letter is a vowel")
			return
		}
	}
	fmt.Println("Letter is a constant")
}