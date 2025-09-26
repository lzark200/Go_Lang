package main

import "fmt"

func concat(s1 string , s2 string)string {
	return s1+s2
}

// we can write function like this : 
func concat2(s1 , s2 string)string {
	return s1+s2
}
/*When the arguments are of the same datatype then there is no requirement 
to mention the datatype for the each parameter instead you can simply do one thing
just mention the parameter in the last variable*/

// another ways if two pairs are of the variable datatype are existing : 
func concat3(s1 , s2 string , x , y int , isCompleted , isTrue , isFalse bool){
	fmt.Println(s1+s2)
	fmt.Println(x , y)
	fmt.Println(isCompleted , isTrue , isFalse)
}


func main(){
	fmt.Println(concat("hello" , " shivi"))
	fmt.Println(concat("ph : " , " 9563214589"))
	fmt.Println(concat("56" , "+89"))
	fmt.Println("hello world")

	concat3("hello" , "ryan" , 15 , 16 , true , false , true)
}