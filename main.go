package main

import (
	"fmt"
)

func sum(a, b int) int {
	return a + b
}

func main() {
	myCar := Car{"BMW", 150}
	myCar.brand = "Ferrari"
	fmt.Println(myCar.brand, myCar.maxSpeed)
}

type Car struct {
	brand    string
	maxSpeed int
}
