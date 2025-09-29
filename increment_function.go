package main

import (
	"fmt" 
	"errors"
)

func main(){
	sendsSoFar:=45
	sendsToAdd:=100
	currentMessage:=incrementSends(sendsSoFar , sendsToAdd)
	fmt.Println("list of the Total messages are: " , currentMessage)
	myname , mywife , goodcouple , doing := getNames()
	fmt.Println(myname , mywife , goodcouple , doing) 
	
	// y := getNames() Illegal go does not allows you to ignore the return value
	
	// So how to ignore the return values in the go : 
	_ ,z, m , _ :=getNames() 
	fmt.Println(z)
	fmt.Println(m)
	// note : we cannot live the unused variable in the go it will throw an error.

	fmt.Println(getCoords())

	a1 , a2 , a3 := getAge(65)
	fmt.Println(a1 , a2 , a3)

	ans , err := divide(5,0)
	fmt.Println(ans , err)
}
func incrementSends(sendsSoFar , sendsToAdd int)int{
	return sendsSoFar+sendsToAdd 
}
// go can return multiple values at the same time : 
func getNames()(string , string , bool , int){
	return "Aditya", "Shonali", true , 369
}

// Naked return functions : This functions should only be used inside the  short functions. As it harms the redability of the longer functions.

func getCoords()(x , y int){
	// x and y are intialized with zero values
	return // automatically returns the x and y 
}

// the naked function can be written as this also : 
func getAge(age int)(int , int , int){
	x , y , z := 45 , 12 , 36
	return age-x , age-y, age-z
}

func divide(dividend , divisor int)(int , error){
	if divisor == 0{
		return 0 , errors.New("Can't divide any number by zero.")
	}
	return dividend/divisor , nil
}





