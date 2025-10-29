package main
import (
    "fmt"
    "regexp"
)

func main() {
    re := regexp.MustCompile(`go`)
    text := "I love golang and go programming"

    match := re.MatchString(text)
    fmt.Println(match) // true
	re2 := regexp.MustCompile(`[a-zA-Z]+`)
    text2 := "Go 123 Python 456 Java"

    words := re2.FindAllString(text2, -1)
    fmt.Println(words)
}
