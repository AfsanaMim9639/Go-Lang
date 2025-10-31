package main

import (
    "fmt"
    "os"
)

func main() {
    // 1. Create a new file (overwrites if exists)
    file, err := os.Create("example.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    // 2. Make sure to close the file at the end
    defer file.Close()

    // 3. Write string to file
    _, err = file.Write([]byte("Hello, Go File Writing!"))
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }

    fmt.Println("File written successfully!")
}
