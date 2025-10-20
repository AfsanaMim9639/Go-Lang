package main
import "fmt"

func main() {
    a := 10
    b := 3
	x := 10
	x += 5   // 15
	x *= 2   // 30

    fmt.Println("Addition:", a+b)
    fmt.Println("Subtraction:", a-b)
    fmt.Println("Multiplication:", a*b)
    fmt.Println("Division:", a/b)
    fmt.Println("Modulus:", a%b)
	result := float64(a) / float64(b)
	fmt.Println(result)
	
	fmt.Println(x)

}
