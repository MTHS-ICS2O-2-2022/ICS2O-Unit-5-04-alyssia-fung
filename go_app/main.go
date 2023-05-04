package main

import "fmt"

func main() {
	var dayOfWeek string
	var age int
	var isDiscountDay bool

	// Prompt user for input
	fmt.Print("Enter day of week (e.g. Monday): ")
	fmt.Scanln(&dayOfWeek)
	fmt.Print("Enter age: ")
	fmt.Scanln(&age)

	// Determine if it's a discount day
	isDiscountDay = dayOfWeek == "Tuesday" || dayOfWeek == "Thursday"

	// Check if customer qualifies for discount
	if isDiscountDay || age >= 13 {
		fmt.Println("You get a discount!")
	} else {
		fmt.Println("Sorry, you don't qualify for a discount.")
	}
	fmt.Println("\nDone")
}
