package main

import (
    "fmt"
    "os"
)

func main() {
    files, err := os.ReadDir(".") // current directory
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    for _, file := range files {
        if file.IsDir() {
            fmt.Println("Directory:", file.Name())
        } else {
            fmt.Println("File:", file.Name())
        }
    }
}
