package main

import "fmt"

func main() {
    x := 42
    f := 3.14159
    s := "Go"

    fmt.Printf("x = %v, f = %.2f, s = %q\n", x, f, s)
    fmt.Printf("Binary of x: %b\n", x)
    fmt.Printf("Hex of x: %#x\n", x)
}
