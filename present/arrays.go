package main

import "fmt"

func main() {
	// START OMIT
	// Arrays
	x := [5]int{10, 20, 30, 40, 50}   // Intialized with values
	var y [5]int = [5]int{10, 20, 30} // Partial assignment

	fmt.Println(x, y)
	
	// Slices
	a := []int{}
	fmt.Println(a)
	a = append(a, 5)
	fmt.Println(a)
	// END OMIT
	
}
