package main
import "fmt"

func main() {
    name := "Afsana"
    age := 22

	
	var name1 string
    var age2 int


    fmt.Print("Name: ", name)
    fmt.Println(" Age:", age)
    fmt.Printf("Hello %s, you are %d years old.\n", name, age)


    fmt.Print("Enter your name: ")
    fmt.Scan(&name1)

    fmt.Print("Enter your age: ")
    fmt.Scan(&age2)

    fmt.Printf("Name: %s, Age: %d\n", name1, age2)
}
