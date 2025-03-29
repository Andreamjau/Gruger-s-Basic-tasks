package main

// ReplaceWord replaces all occurrences of a
// substring `old`` with `new` in the given `text`
// Example:
//   ReplaceWord("foo bar foo", "foo", "baz") => "baz bar baz"
func ReplaceWord(text, old, new string) string {
    var currentWord string
    var words []string
    
    // Storing words in a slice
    for _, letter := range text{
        if string(letter) == " "{
            if currentWord != ""{
                words = append(words, currentWord)
                currentWord = ""
            }
        } else {
            currentWord += string(letter)
        }
        
    }
    if currentWord != ""{
        words = append(words, currentWord)
    }

    // Looping though words and replacing old word with new 
    for i := 0; i < len(words); i++{
        if words[i] == old {
            words[i] = new
        }
    }

    // Creating a string from slice
    return JoinWords(words)
}