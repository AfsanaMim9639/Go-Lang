package main
import "fmt"

func main() {
    counter := func() func() int {
        x := 0
        return func() int {
            x++
            return x
        }
    }()

    fmt.Println(counter()) // 1
    fmt.Println(counter()) // 2
    fmt.Println(counter()) // 3
}
