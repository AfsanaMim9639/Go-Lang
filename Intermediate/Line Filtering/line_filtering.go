package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    filename := "example.txt"

    // -----------------------------
    // 1️⃣ Create & Write to file
    initialContent := "Hello World\nGo is awesome\nI love programming\nGoLang is fun\nBye World"
    err := os.WriteFile(filename, []byte(initialContent), 0644)
    if err != nil {
        fmt.Println("Error writing file:", err)
        return
    }
    fmt.Println("File created and initial content written.\n")

    // -----------------------------
    // 2️⃣ Append a new line
    file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Error opening file for append:", err)
        return
    }
    defer file.Close()

    newLine := "\nAppending a new Go line!"
    _, err = file.WriteString(newLine)
    if err != nil {
        fmt.Println("Error appending to file:", err)
        return
    }
    fmt.Println("New line appended.\n")

    // -----------------------------
    // 3️⃣ Read file line by line & filter lines containing "Go"
    fmt.Println("Filtered lines containing 'Go':")
    file, err = os.Open(filename)
    if err != nil {
        fmt.Println("Error opening file for reading:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, "Go") {
            fmt.Println(line)
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error scanning file:", err)
    }
}
