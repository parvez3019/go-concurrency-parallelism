package main

import "fmt"

func main() {
	// START OMIT
	// Mutiple ways to initialise a variable
	sadnessLevel := 2020 // or var a int
	fmt.Printf("On a scale of 0-10, how sad are you today? \nMe - %d", sadnessLevel)

	// Arrays
	randomArray := [5]int{10, 20, 30, 40, 50}   // Intialized with values
	var anotherRandomArray [5]int = [5]int{10, 20, 30} // Partial assignment

	fmt.Println(randomArray, anotherRandomArray)

	// Slices
	slice := []int{}
	fmt.Println(slice)
	slice = append(slice, 5)
	fmt.Println(slice)

	// END OMIT

}
