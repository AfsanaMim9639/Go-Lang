package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println("Pi:", math.Pi)
    fmt.Println("E:", math.E)

    x := -5.7
    fmt.Println("Abs:", math.Abs(x))
    fmt.Println("Sqrt(16):", math.Sqrt(16))
    fmt.Println("2^3:", math.Pow(2, 3))
    fmt.Println("Log(E):", math.Log(math.E))
    fmt.Println("Log10(100):", math.Log10(100))
    fmt.Println("Sin(Pi/2):", math.Sin(math.Pi/2))
    fmt.Println("Cos(0):", math.Cos(0))
    fmt.Println("Tan(Pi/4):", math.Tan(math.Pi/4))
    fmt.Println("Ceil(3.2):", math.Ceil(3.2))
    fmt.Println("Floor(3.7):", math.Floor(3.7))
    fmt.Println("Round(3.5):", math.Round(3.5))
    fmt.Println("Mod(5,2):", math.Mod(5, 2))
    fmt.Println("Max(5,10):", math.Max(5, 10))
    fmt.Println("Min(5,10):", math.Min(5, 10))
}
