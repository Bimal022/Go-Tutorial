package main

import (
	"fmt"
)
/*********************Single Return Type********************/

// Function to calculate the total price
func calculateBill(price, no int) int {
	totalPrice := price * no
	return totalPrice
}
/*********************Multiple Return Type********************/


// Function to calculate area and perimeter of a rectangle
func rectProps(length, width float64) (float64, float64) {
	area := length * width
	perimeter := (length + width) * 2
	return area, perimeter
}
//Named Return Value
/*
func rectProps(length, width float64)(area, perimeter float64) {  
    area = length * width
    perimeter = (length + width) * 2
    return // no explicit return value, as area and perimeter are named return values
}
*/
func main() {
	// Calling the calculateBill function
	price, no := 90, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)

	// Calling the rectProps function
	rectLength, rectWidth := 10.8, 5.6
	area, perimeter := rectProps(rectLength, rectWidth)
	fmt.Printf("Rectangle Area: %f, Perimeter: %f\n", area, perimeter)


	// Using the blank identifier to discard a value
	areaOnly, _ := rectProps(12.5, 7.3)
	fmt.Printf("Discarding Perimeter - Area: %f\n", areaOnly)
	//we use only the area and the _ identifier is used to discard the perimeter.
}
