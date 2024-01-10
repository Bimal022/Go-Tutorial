package main

import (
	"fmt"
)

func main() {
	// Boolean type
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)

	c := a && b // Logical AND operation
	fmt.Println("c:", c)

	d := a || b // Logical OR operation
	fmt.Println("d:", d)

	// Signed integers
	var intVar int = 89
	shortIntVar := 95
	fmt.Println("value of intVar is", intVar, "and shortIntVar is", shortIntVar)

	// Size and type of intVar
	fmt.Printf("type of intVar is %T, size of intVar is %d\n", intVar, unsafe.Sizeof(intVar))

	// Unsigned integers
	var uintVar uint = 100
	uintShortVar := uint(200)
	fmt.Println("value of uintVar is", uintVar, "and uintShortVar is", uintShortVar)

	// Floating point types
	floatVar1, floatVar2 := 5.67, 8.97
	fmt.Printf("type of floatVar1 %T, floatVar2 %T\n", floatVar1, floatVar2)

	sum := floatVar1 + floatVar2
	diff := floatVar1 - floatVar2
	fmt.Println("sum", sum, "diff", diff)

	intVar1, intVar2 := 56, 89
	fmt.Println("sum", intVar1+intVar2, "diff", intVar1-intVar2)

	// Complex types
	complexVar1 := complex(5, 7)
	complexVar2 := 8 + 27i
	complexSum := complexVar1 + complexVar2
	complexProduct := complexVar1 * complexVar2
	fmt.Println("sum:", complexSum)
	fmt.Println("product:", complexProduct)

	// Other numeric types
	var byteVar byte = 'A'
	var runeVar rune = '♥'

	fmt.Printf("byteVar: %c, runeVar: %c\n", byteVar, runeVar)

	// String type
	firstName := "Naveen"
	lastName := "Ramanathan"
	fullName := firstName + " " + lastName
	fmt.Println("My name is", fullName)

	// Type conversion
	i := 55      // int
	j := 67.8    // float64

	// Explicit conversion from float64 to int
	sumConverted := i + int(j)
	fmt.Println("Sum after conversion:", sumConverted)

	// Explicit conversion from int to float64
	var k float64 = float64(i)
	fmt.Println("k (after conversion):", k)
}
