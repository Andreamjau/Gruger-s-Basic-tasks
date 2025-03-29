package main

// JoinWords joins a list of words into a sentence
// Example:
//   JoinWords([]string{"i", "love", "cheese"}) => "i love cheese"
func JoinWords(words []string) string {
    var sentence string

    if len(words) == 0 {
        return ""
    }

    for _, word := range words{
        if sentence == ""{
            sentence += word
            continue
        }
        sentence += " " + word
    }
    return sentence
}
