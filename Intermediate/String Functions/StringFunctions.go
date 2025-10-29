package main
import (
    "fmt"
    "strings"
)

func main() {
    str := "GoLang"
    fmt.Println(strings.ToUpper(str))
    fmt.Println(strings.ToLower(str))

	fmt.Println(strings.Contains("Bangladesh", "desh"))  // true
	fmt.Println(strings.HasPrefix("Golang", "Go"))       // true
	fmt.Println(strings.HasSuffix("hello.txt", ".txt"))  // true

	words := strings.Split("apple,banana,mango", ",")
	fmt.Println(words)  // ["apple" "banana" "mango"]

	joined := strings.Join(words, "-")
	fmt.Println(joined) // apple-banana-mango

}
