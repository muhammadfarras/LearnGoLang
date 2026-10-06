package main

import "fmt"

var nilai1 int8 = 10
var nilai2 float64 = 17.54321

func learnTypeConversion() {
	var sum1 float64 = float64(nilai1) + nilai2
	var sum2 uint8 = uint8(nilai1) + uint8(nilai2)
	var sum3 uint8 = sum2 + uint8(nilai1)

	fmt.Printf("Sum1 : %f\nSum2 : %d\nSum3 : %d\n", sum1, sum2, sum3)
}

func learnAdditionLiteral() {
	var sum4 float32 = 20.1 + 2
	fmt.Printf("Sum 4 : %f\n", sum4)
}
