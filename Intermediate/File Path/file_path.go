package main

import (
    "fmt"
    "os"
)

func main() {
    cwd, err := os.Getwd()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Current Working Directory:", cwd)
}
