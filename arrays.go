package main

import "fmt"

func main() {
	var ArrayDeclarations [3]int = [3]int {1, 2, 3}
				//OR
	var EasyArrayDeclaration = [3]int {2, 3, 4}

	fmt.Println(ArrayDeclarations, len(ArrayDeclarations))
	fmt.Println(EasyArrayDeclaration, len(EasyArrayDeclaration))


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
	var scores = []int{1, 2, 3, 4, 5}
	scores[2] = 10
	scores = append(scores, 25)
	fmt.Println("Slice of scores: ",scores)
	// Slice ranges
	rangeone := scores[1:3]
	rangetwo := scores[1:]
	rangethree := scores[:4]
	fmt.Printf("Range One : %v \nRange two : %v \nRange three : %v\n", rangeone, rangetwo, rangethree)
	c := a[1:3]                 // Create a slice from array a, covering elements 1 to 2
	fmt.Println(c)             // Output: [78 50]


	//Slices are arrays with non definite size
	d := []int{6, 7, 8}         // Create a slice directly
	fmt.Println(d) 

}