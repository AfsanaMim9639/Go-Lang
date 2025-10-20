package main

import "fmt"

func main() {
	temperature := 25
	if temperature >= 30 {
		fmt.Println("It's hot outside.")
	}else{
		fmt.Println("It's cool outside.")
	}





	score := 85  // you can change this value

	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else {
		fmt.Println("D")
	}
}
