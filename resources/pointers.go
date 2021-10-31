package main

import "fmt"

func changeValue(str *string) {
	*str = "changed!"
}

func changeValue2(str string) {
	str = "changed!"
}

func main() {
	fmt.Println("Working with variables")
	// Simple Operations
	x := 7
	y := &x
	fmt.Println(x, y)
	*y = 8
	fmt.Println(x, y)

	fmt.Println("Working with functions")
	toChange := "hello"
	fmt.Println("Before: ", toChange)
	changeValue2(toChange)
	changeValue(&toChange)
	fmt.Println("After Chnage by Reference: ", toChange)

}
