### Pointers (&) and Dereference Operators
Pointers are reference to actual variables in a RAM. When you change the content of the memory location, you will also change variables. 
- & -- gives address of the variable
- use dereference operator(*) to work with the pointers

### Interfaces

### Closures
```go
package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
```
Output example
```bash
➜  resources git:(main) ✗ go run closure.go
1
2
3
1
```


### Anonymous Functions
### Maps
- Dicitionary
```go
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
```

### References
1. https://www.youtube.com/watch?v=yJE2RC37BF4&list=PLzMcBGfZo4-mtY_SE3HuzQJzuj4VlUG0q&index=15
1. https://gobyexample.com/closures
