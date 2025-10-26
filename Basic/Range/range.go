//Range with Arrays and Slices
package main
import "fmt"

func main() {
    nums := []int{10, 20, 30}

    for index, value := range nums {
        fmt.Println(index, value)
    }

	text := "GoLang"

for index, runeValue := range text {
    fmt.Printf("%d: %c\n", index, runeValue)
}


ch := make(chan int)

go func() {
    for i := 1; i <= 3; i++ {
        ch <- i
    }
    close(ch)
}()

for value := range ch {
    fmt.Println(value)
}

}


