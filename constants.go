package main

import (
	"fmt"
	"math"
)

func main() {
	/******Constants  cannot be reassigned again to any other value.*******/
	const a = 50
	fmt.Println(a)
	// a = 100 //ERROR
	// fmt.Println(a)

	/*********The value of a constant should be known at compile time*******/
	var b = math.Sqrt(4)   //ALLOWED
	// const b = math.Sqrt(4) //NOT ALLOWED
	fmt.Println(b)


}
