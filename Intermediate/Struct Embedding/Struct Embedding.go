package main
import "fmt"

type Person struct {
    name string
    age  int
}

type Employee struct {
    Person      // Embedded struct
    company string
}

func main() {
    e := Employee{
        Person:  Person{"Afsana", 22},
        company: "Google",
    }

    fmt.Println(e.name)   // Access directly (no need for e.Person.name)
    fmt.Println(e.age)
    fmt.Println(e.company)
}
