package main
import "fmt" 

type Address struct {
    city, country string
}

type Employee struct {
    name    string
    salary  float64
    address Address
}

func main() {
    emp := Employee{
        name:   "Rafi",
        salary: 50000,
        address: Address{
            city:    "Dhaka",
            country: "Bangladesh",
        },
    }

    fmt.Println(emp.address.city)
}
