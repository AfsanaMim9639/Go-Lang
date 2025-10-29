package main
import "fmt"

type Rectangle struct {
    width, height float64
}

// Method
func (r Rectangle) Area() float64 {
    return r.width * r.height
}


type Student struct {
    name  string
    marks int
}

// Pointer receiver
func (s *Student) UpdateMarks(m int) {
    s.marks = m
}

func main() {
	r := Rectangle{10, 5}
    fmt.Println("Area:", r.Area())

    st := Student{"Afsana", 85}
    st.UpdateMarks(95)
    fmt.Println(st)
}