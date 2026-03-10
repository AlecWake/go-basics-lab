package main

import "fmt"

type Counter struct {
	Value int
}

func IncrementValue(c Counter) {
	c.Value++
}

func IncrementPointer(c *Counter) {
	c.Value++
}

func main() {

	c := Counter{Value: 5}

	fmt.Println("Initial value:", c.Value)

	IncrementValue(c)
	fmt.Println("After IncrementValue:", c.Value)

	IncrementPointer(&c)
	fmt.Println("After IncrementPointer:", c.Value)
}
