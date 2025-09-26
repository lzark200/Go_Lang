// string interpolation :
/*fmt.Sprintf : it returns the string*/
package main

import (
	"fmt"
)

func main() {
	name, age := "aditya", 25

	result := fmt.Sprintf("My name is %s and I am %d years old", name, age)
	fmt.Println(result)

}
