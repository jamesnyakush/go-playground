package main

import "fmt"

func main() {

	var a [5]int

	fmt.Println("Empty array",a)

	a[4] = 50

	fmt.Println("Set",a)
	fmt.Println("Get",a[4])

	b := [5]int{1,2,3,4,5}

	fmt.Println("Long array",b)
}
