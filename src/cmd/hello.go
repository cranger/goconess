package main

import (
	"fmt"
	"clock"
)

func main() {
	fmt.Println("Hello! There")
	if clock.IsAm() {
		fmt.Print("hey you")
	} else {
		fmt.Print(" with you")
	}
}