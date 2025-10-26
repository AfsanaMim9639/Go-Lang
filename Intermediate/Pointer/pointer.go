package main
import "fmt"

func main() {
    var x int = 10
    var p *int = &x // & gives the memory address of x

    fmt.Println("x =", x)
    fmt.Println("Address of x =", &x)
    fmt.Println("Pointer p =", p)


	y := 10
    p1 := &y

    *p1 = 20 // change the value of x via pointer
    fmt.Println(y) // prints 20
}
