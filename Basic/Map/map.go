package main

import "fmt"

func main() {
	// Adding/Updating
	ages := make(map[string]int)
	ages["Alice"] = 25
	ages["Bob"] = 30

	// Reading
	age := ages["Alice"] // 25
	fmt.Println("Alice's age:", age)

	// Checking if key exists (comma ok idiom)
	charlieAge, ok := ages["Charlie"]
	if ok {
		fmt.Println("Charlie is", charlieAge)
	} else {
		fmt.Println("Charlie not found")
	}

	// Deleting
	delete(ages, "Bob")

	// Length
	fmt.Println("Number of people:", len(ages))

	// Show final map
	fmt.Println("Final ages map:", ages)
}