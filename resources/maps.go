package main

import "fmt"

func main() {
	// Using map operators
	var mp map[string]int = map[string]int{
		"apple":   1,
		"orange":  0,
		"bananas": 5,
	}

	// Accessing
	fmt.Println("Apples: ", mp["apple"])

	//Changing values
	mp["apple"] = 900
	fmt.Println("Apples: ", mp["apple"])

	fmt.Println(mp)

	// Adding new
	mp["peer"] = 8
	fmt.Println(mp)

	// Delete
	delete(mp, "orange")
	fmt.Println("After deleting orange: ", mp)

	// empty map with make

	mp2 := make(map[string]int)
	fmt.Println(mp2)

}
