package main

import "fmt"

func main() {
	// START OMIT
	// Mutiple ways to initialise a variable
	sadnessLevel := 2020 // or var sadnessLevel int
	fmt.Printf("On a scale of 0-10, how sad are you today? \nME - %d\n", sadnessLevel)

	// Arrays
	randomArray := [5]int{10, 20, 30, 40, 50}   // Intialized with values
	randomArray[1] = 21 
	fmt.Println("Priting Array : ", randomArray)

	// Slices
	slice := []string{}
	fmt.Println("Printing Slice : ", slice)
	slice = append(slice, "FEELINGS")
	fmt.Printf("This slice have %+v in it",slice)

	// END OMIT

}
