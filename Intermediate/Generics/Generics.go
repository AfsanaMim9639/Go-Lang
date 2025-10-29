package main
import "fmt"

// Generic function
func Add[T int | float64](a, b T) T {
    return a + b
}

func main() {
    fmt.Println(Add(10, 20))        // int
    fmt.Println(Add(3.5, 2.5))      // float64
}
