package main

import "fmt"

func main() {
	var a [3]int         // Declare an integer array of length 3 with values 0
	fmt.Println(a)
	
	// Assign values to array elements

	a[0] = 12             
	a[1] = 78
	a[2] = 50
	fmt.Println(a)

	b := [3]int{12, 78, 50}     // Short-hand declaration to create array
	fmt.Println(b)

	/*****************Slices*****************/
	c := a[1:3]                 // Create a slice from array a, covering elements 1 to 2
	fmt.Println(c)             // Output: [78 50]

	d := []int{6, 7, 8}         // Create a slice directly
	fmt.Println(d) 

}