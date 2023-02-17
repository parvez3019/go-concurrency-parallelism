package main

import "fmt"

func main() {
	// START OMIT
	// Mutiple ways to initialise a variable
	sadnessLevel := 2020 // or var sadnessLevel int
	fmt.Printf("On a scale of 0-10, how sad are you today? \nME - %d\n", sadnessLevel)

	// Arrays
	randomArray := [4]string{"This", "year", "is", "amazing"}   // Intialized with values
	randomArray[3] = "awful" 
	fmt.Println(randomArray)

	// END OMIT

}
