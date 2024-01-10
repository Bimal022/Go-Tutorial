package main

import (
	"fmt"
	"reflect"
	"math"
)

func main() {
	/**************Declaring Single Variables********************/
	var a int
	var b = 10
	c:=11 //Shorthand Method

	fmt.Printf("a is of type %v, b is of type %v, c is of type %v", reflect.TypeOf(a), 
	reflect.TypeOf(b), reflect.TypeOf(c))
	fmt.Println()

	/******************Declaring Multiple Variables******************/

	var (
		name   = "Kumar Bimal"
		enrollment = "E21CSEU0860"
		age = 23
	)

	fmt.Println("My name is", name, ", Enrollment Number is", enrollment, "and age is", age)


	/**********************Comparing variables******************/


	x:= 10.52
	y:= 15.2321
	z:= math.Min(x, y)
	fmt.Println("Minimum of ", x, " and ", y, " is", z)

	x, y = 120, 140
	z = math.Max(x, y)
	fmt.Println("Maximum of ", x, " and ", y, " is", z)
}
