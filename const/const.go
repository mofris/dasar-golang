package main

import "fmt"

func main() {
	const name string = "Faris"
	const age = 20
  
	fmt.Println(name)
	fmt.Println(age)

	// deklarasi multiple
	const (
		firstname = "Mochammad"
		lastname = "Faris"
	)

	fmt.Println(firstname)
	fmt.Println(lastname)
}