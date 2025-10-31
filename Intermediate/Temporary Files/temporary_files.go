package main

import (
    "fmt"
    "os"
)

func main() {
    // Creates a temporary file in OS default temp folder
    tempFile, err := os.CreateTemp("", "mytempfile_*.txt")
    if err != nil {
        fmt.Println("Error creating temp file:", err)
        return
    }
    defer os.Remove(tempFile.Name()) // delete file after program ends
    defer tempFile.Close()

    fmt.Println("Temp file created:", tempFile.Name())

    // Write something to temp file
    tempFile.WriteString("This is temporary content.\n")
}
