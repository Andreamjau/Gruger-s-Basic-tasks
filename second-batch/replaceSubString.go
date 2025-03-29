package main

// I have renamed the function to be more truthful
// text = "Hello World"
//ReplaceSubString(text, "Hello", "Ciao") -> "Ciao World"
//ReplaceSubString(text, "l", ":)") -> "He:):)o Wor:)d"
//ReplaceSubString(text, "World", "World, im sleepy!") -> "Hello World, im sleepy!"

func ReplaceSubString(text, old, new string) string{
    i := 0 
    var result string
    for i < len(text){
        if len(old) + i <= len(text){
            if old == text[i:len(old) + i]{
                result += new      
                i += len(old)   
            } else {
                result += string(text[i])
                i++
            }
        } else {
            result += string(text[i])
            i++
        }
    }
    return result
}
