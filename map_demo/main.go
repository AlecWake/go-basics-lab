package main

import "fmt"

func modifyMap(m map[string]int) {
	m["grape"] = 50
}

func main() {

	// create map
	fruit := make(map[string]int)

	// insert values
	fruit["apple"] = 5
	fruit["banana"] = 10

	fmt.Println("initial map:", fruit)

	// read existing key
	fmt.Println("apple count:", fruit["apple"])

	// read missing key
	fmt.Println("orange count:", fruit["orange"])

	// comma-ok syntax
	value, ok := fruit["orange"]

	if ok {
		fmt.Println("orange exists:", value)
	} else {
		fmt.Println("orange key not found")
	}

	// demonstrate reference behavior
	modifyMap(fruit)

	fmt.Println("after modifyMap:", fruit)
}
