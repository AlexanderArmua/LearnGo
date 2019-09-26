package main

import "fmt"

func main() {
	defer fmt.Printf("World")

	fmt.Printf("Hello ")
}
