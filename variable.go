package main

import "fmt"

func main() {
	// we can declare the same variable in the same line.
	x, y := 5, 6
	fmt.Println(x, y)

	// printing the type of the variable in the go:
	fmt.Printf("%T", y)
	fmt.Println()

	// typecasting in the golang :
	rate := 5.6398
	fmt.Println(rate)
	newRate := int(rate)
	fmt.Println(newRate)

	// constants in the go lang :
	const premiumPlan string = "jio-hotstar-1month-scheme"
	const age int8 = 100
	// we cannot do the short declaration in the constants : const premiumPlan:="abcdef"  illegal it is not allowed in the cases of the const❌❌
	fmt.Printf("%T\n", premiumPlan)
	fmt.Println(age)
}
