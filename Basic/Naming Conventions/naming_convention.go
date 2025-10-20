package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase
	//Eg. CalculateArea, UserInfo, NewHTTPRequest
	// Structs, interfaces, enums

	//snake_case
	// Eg. user_id, first_name,http_request

	//uppercase
	//use case is constants
	//mixedcase
	//Eg. js, htmlDocument, isvalid

	const MAXRETRIES = 5

	var employeeId = 1001
	fmt.Println(employeeId)

}