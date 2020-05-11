package main

import "fmt"

func main() {
	// START OMIT
	sites := []string{"google", "facebook", "thoughtworks"}

	for _, site := range sites {
		if site == "google" {
			fmt.Printf("Ok %s \n", site)
		} else if site == "facebook" {
			fmt.Println(site)
		} else {
			fmt.Printf("According to %s this is a very bad code", site)
		}
	}
	// END OMIT

}
