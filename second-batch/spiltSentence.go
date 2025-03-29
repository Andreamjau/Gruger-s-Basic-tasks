package main

// SplitSentence splits a given sentence by spaces
// and returns the resulting slice of words.
// Example:
//   SplitSentence("aaaaa hello there") => []string{"aaaaa", "hello", "there"}
func SplitSentence(sentence string) []string {
    var words []string
    var word string
    
    if len(sentence) <= 2{
        return append(words, sentence)
    }

    for _, letter := range sentence{
        if string(letter) == " "{
            if word != ""{
                words = append(words, word)
                word = ""
            }
            continue
        }

        word += string(letter)
        
    }

    if word != ""{
        words = append(words, word)
    }
    return words
}