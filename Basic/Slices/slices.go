package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Go Slices Comprehensive Guide ===\n")

	// 1. Creating Slices
	demonstrateCreation()

	// 2. Slicing Operations
	demonstrateSlicing()

	// 3. Appending to Slices
	demonstrateAppend()

	// 4. Copying Slices
	demonstrateCopy()

	// 5. Slice Internals (length vs capacity)
	demonstrateInternals()

	// 6. Modifying Slices
	demonstrateModification()

	// 7. Iterating Over Slices
	demonstrateIteration()

	// 8. Common Operations
	demonstrateCommonOps()

	// 9. Nil vs Empty Slices
	demonstrateNilVsEmpty()

	// 10. Multi-dimensional Slices
	demonstrateMultiDimensional()
}

// demonstrateCreation shows different ways to create slices
func demonstrateCreation() {
	fmt.Println("1. Creating Slices")
	fmt.Println("------------------")

	// Slice literal
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice literal: %v\n", s1)

	// Using make() - length and capacity
	s2 := make([]int, 5)      // length 5, capacity 5
	s3 := make([]int, 3, 10)  // length 3, capacity 10
	fmt.Printf("make([]int, 5): %v, len=%d, cap=%d\n", s2, len(s2), cap(s2))
	fmt.Printf("make([]int, 3, 10): %v, len=%d, cap=%d\n", s3, len(s3), cap(s3))

	// From an existing array
	arr := [5]int{10, 20, 30, 40, 50}
	s4 := arr[1:4] // slice from index 1 to 3
	fmt.Printf("Slice from array arr[1:4]: %v\n", s4)

	// Nil slice
	var s5 []int
	fmt.Printf("Nil slice: %v, len=%d, cap=%d, is nil=%t\n\n", s5, len(s5), cap(s5), s5 == nil)
}

// demonstrateSlicing shows slicing operations
func demonstrateSlicing() {
	fmt.Println("2. Slicing Operations")
	fmt.Println("---------------------")

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Original slice: %v\n", s)
	fmt.Printf("s[2:5]: %v\n", s[2:5])
	fmt.Printf("s[:4]: %v\n", s[:4])
	fmt.Printf("s[6:]: %v\n", s[6:])
	fmt.Printf("s[:]: %v\n\n", s[:])
}

// demonstrateAppend shows how append works
func demonstrateAppend() {
	fmt.Println("3. Appending to Slices")
	fmt.Println("----------------------")

	s := []int{1, 2, 3}
	fmt.Printf("Initial: %v, len=%d, cap=%d\n", s, len(s), cap(s))

	s = append(s, 4)
	fmt.Printf("After append(s, 4): %v, len=%d, cap=%d\n", s, len(s), cap(s))

	s = append(s, 5, 6, 7)
	fmt.Printf("After append(s, 5, 6, 7): %v, len=%d, cap=%d\n", s, len(s), cap(s))

	// Append another slice
	s2 := []int{8, 9, 10}
	s = append(s, s2...)
	fmt.Printf("After append(s, s2...): %v, len=%d, cap=%d\n\n", s, len(s), cap(s))
}

// demonstrateCopy shows how to copy slices
func demonstrateCopy() {
	fmt.Println("4. Copying Slices")
	fmt.Println("-----------------")

	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	n := copy(dst, src)
	fmt.Printf("Source: %v\n", src)
	fmt.Printf("Destination: %v\n", dst)
	fmt.Printf("Elements copied: %d\n", n)

	// Modify destination to show they're independent
	dst[0] = 99
	fmt.Printf("After modifying dst[0]: src=%v, dst=%v\n\n", src, dst)
}

// demonstrateInternals shows length vs capacity
func demonstrateInternals() {
	fmt.Println("5. Slice Internals (Length vs Capacity)")
	fmt.Println("----------------------------------------")

	s := make([]int, 3, 5)
	fmt.Printf("Initial: %v, len=%d, cap=%d\n", s, len(s), cap(s))

	// Fill initial values
	s[0], s[1], s[2] = 1, 2, 3
	fmt.Printf("After assignment: %v, len=%d, cap=%d\n", s, len(s), cap(s))

	// Append within capacity
	s = append(s, 4)
	fmt.Printf("After append(s, 4): %v, len=%d, cap=%d\n", s, len(s), cap(s))

	// Append beyond capacity (triggers reallocation)
	s = append(s, 5, 6)
	fmt.Printf("After append(s, 5, 6): %v, len=%d, cap=%d\n\n", s, len(s), cap(s))
}

// demonstrateModification shows shared underlying array behavior
func demonstrateModification() {
	fmt.Println("6. Modifying Slices (Shared Underlying Array)")
	fmt.Println("----------------------------------------------")

	original := []int{1, 2, 3, 4, 5}
	slice1 := original[1:4] // [2, 3, 4]
	slice2 := original[2:5] // [3, 4, 5]

	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Slice1 [1:4]: %v\n", slice1)
	fmt.Printf("Slice2 [2:5]: %v\n", slice2)

	// Modify slice1
	slice1[0] = 99
	fmt.Printf("\nAfter slice1[0] = 99:\n")
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Slice1: %v\n", slice1)
	fmt.Printf("Slice2: %v (also affected!)\n\n", slice2)
}

// demonstrateIteration shows different ways to iterate
func demonstrateIteration() {
	fmt.Println("7. Iterating Over Slices")
	fmt.Println("------------------------")

	fruits := []string{"apple", "banana", "cherry", "date"}

	// Method 1: Using range with index and value
	fmt.Println("Using range (index, value):")
	for i, fruit := range fruits {
		fmt.Printf("  [%d] = %s\n", i, fruit)
	}

	// Method 2: Using range with only value
	fmt.Println("Using range (value only):")
	for _, fruit := range fruits {
		fmt.Printf("  - %s\n", fruit)
	}

	// Method 3: Traditional for loop
	fmt.Println("Using traditional for loop:")
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("  fruits[%d] = %s\n", i, fruits[i])
	}
	fmt.Println()
}

// demonstrateCommonOps shows common slice operations
func demonstrateCommonOps() {
	fmt.Println("8. Common Slice Operations")
	fmt.Println("--------------------------")

	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original: %v\n", s)

	// Insert element at index 2
	s = insert(s, 2, 99)
	fmt.Printf("After insert(2, 99): %v\n", s)

	// Delete element at index 3
	s = remove(s, 3)
	fmt.Printf("After remove(3): %v\n", s)

	// Filter even numbers
	s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := filter(s, func(n int) bool { return n%2 == 0 })
	fmt.Printf("Filter evens from %v: %v\n", s, evens)

	// Reverse slice
	s = []int{1, 2, 3, 4, 5}
	reverse(s)
	fmt.Printf("Reversed: %v\n\n", s)
}

// demonstrateNilVsEmpty shows the difference between nil and empty slices
func demonstrateNilVsEmpty() {
	fmt.Println("9. Nil vs Empty Slices")
	fmt.Println("----------------------")

	var nilSlice []int
	emptySlice := []int{}
	sliceWithMake := make([]int, 0)

	fmt.Printf("Nil slice: %v, len=%d, cap=%d, is nil=%t\n",
		nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
	fmt.Printf("Empty slice (literal): %v, len=%d, cap=%d, is nil=%t\n",
		emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)
	fmt.Printf("Empty slice (make): %v, len=%d, cap=%d, is nil=%t\n\n",
		sliceWithMake, len(sliceWithMake), cap(sliceWithMake), sliceWithMake == nil)
}

// demonstrateMultiDimensional shows 2D slices
func demonstrateMultiDimensional() {
	fmt.Println("10. Multi-dimensional Slices")
	fmt.Println("----------------------------")

	// Create a 2D slice (slice of slices)
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("2D Slice (matrix):")
	for i, row := range matrix {
		fmt.Printf("  Row %d: %v\n", i, row)
	}

	// Create jagged slice
	jagged := make([][]int, 3)
	jagged[0] = []int{1}
	jagged[1] = []int{2, 3}
	jagged[2] = []int{4, 5, 6}

	fmt.Println("\nJagged Slice:")
	for i, row := range jagged {
		fmt.Printf("  Row %d: %v\n", i, row)
	}
	fmt.Println()
}

// Helper Functions

// insert adds an element at the specified index
func insert(s []int, index, value int) []int {
	s = append(s, 0)              // Expand slice by 1
	copy(s[index+1:], s[index:])  // Shift elements right
	s[index] = value              // Insert value
	return s
}

// remove deletes an element at the specified index
func remove(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

// filter returns a new slice containing elements that satisfy the predicate
func filter(s []int, fn func(int) bool) []int {
	result := []int{}
	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// reverse reverses the slice in place
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}