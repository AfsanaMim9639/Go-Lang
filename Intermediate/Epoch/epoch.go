//🔁 Epoch থেকে Time এ রূপান্তর করা
package main
import (
    "fmt"
    "time"
)

func main() {
    epoch := int64(1761609725)
    t := time.Unix(epoch, 0)
    fmt.Println("Time from Epoch:", t)
}