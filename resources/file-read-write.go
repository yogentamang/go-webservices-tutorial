package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("localdata.data")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
