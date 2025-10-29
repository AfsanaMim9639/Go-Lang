package main
import "fmt"

// Interface
type Speaker interface {
    Speak()
}

// Struct 1
type Person struct {
    name string
}

// Struct 2
type Robot struct {
    id int
}

// Implement method for Person
func (p Person) Speak() {
    fmt.Println(p.name, "বলছে: হ্যালো!")
}

// Implement method for Robot
func (r Robot) Speak() {
    fmt.Println("রোবট", r.id, "বলছে: বীপ বীপ!")
}

func main() {
    var s Speaker

    s = Person{"Afsana"}
    s.Speak()

    s = Robot{101}
    s.Speak()
}
