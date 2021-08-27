package main

import "fmt"

func main() {
	fmt.Printf("Name: ")
	var name string
	fmt.Scanln(&name)

	fmt.Printf("Age: ")
	var age int
	fmt.Scanf("%d", &age)

	fmt.Printf("Hello, %s! You're at least %d days old!\n", name, age * 365)
}