package main

import "fmt"

func main() {
	//Print in single line
	fmt.Print("Hello World")
	fmt.Print("Hello Bimal")

	//Pring in new line
	fmt.Println("Hello World")
	fmt.Println("Hello Bimal")

	//Formatted strings(Printf)
	//				%_(format specifier)
	var age int = 23
	vat name string = "Kumar Bimal"
	fmt.Printf("My name is %v and my age is %v", name, age)
	
	//Sprintf: Save the returned string in a variable
	var str := fmt.Sprintf("Hello my name is %v and my age is %v", name, age)
	fmt.Println("Saved string is: %v", str)
} 