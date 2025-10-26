//Factorial of a number (n!) = n × (n−1) × (n−2) × … × 1

package main
import "fmt"

func factorial(n int) int {
    if n == 0 {
        return 1 // base case
    }
    return n * factorial(n-1) // recursive case
}


//Example 2: Fibonacci Sequence

func fibonacci(n int) int {
    if n <= 1 {
        return n // base case
    }
    return fibonacci(n-1) + fibonacci(n-2)
}


func main() {
    fmt.Println(factorial(5))

	for i := 0; i < 10; i++ {
        fmt.Print(fibonacci(i), " ")
    }
}




