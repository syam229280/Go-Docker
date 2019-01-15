package main

import (
	"fmt"
	"math"
)

func priceCalculator() {
	var age int = 10
	fmt.Println("Hello World")
	fmt.Println("Age is", age)

	var (
		name = "syam"
		year = 1987
	)

	fmt.Println("User", name, "is born in year", year)

	name, year1 := "syam", 1988 // atleast one variable should be new
	fmt.Println("User", name, "is born in year", year1)

	val1, val2 := 23.00, 43.00
	minVal := math.Min(val1, val2)
	fmt.Println(minVal)
	fmt.Println("Price is : ", calcPrice(10, 5))
	getArray()
}

func calcPrice(unit_price int, quantity int) int {
	var totalPrice = unit_price * quantity
	return totalPrice
}

func getArray() {
	var a [2]int
	a[0] = 2
	a[1] = 3
	fmt.Println(a)
	b := [2]int{2, 3}
	fmt.Println(b)
}
