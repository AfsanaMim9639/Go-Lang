package main
import "fmt"

func main() {
    // Step 1: Declare and initialize an array
    numbers := [5]int{10, 20, 30, 40, 50}

    // Step 2: Print all elements with index and value
    fmt.Println("All elements with index and value:")
    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

    fmt.Println()

    // Step 3: Use blank identifier (_) to ignore index
    fmt.Println("Only values (ignore index):")
    for _, value := range numbers {
        fmt.Println("Value:", value)
    }

    fmt.Println()

    // Step 4: Use blank identifier to ignore function return value
    quotient, _ := divide(10, 3)
    fmt.Println("Result of divide(10, 3): Quotient =", quotient)
}

// Function that returns two values
func divide(a, b int) (int, int) {
    return a / b, a % b
}
