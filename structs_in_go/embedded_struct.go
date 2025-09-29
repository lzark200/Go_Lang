/*Go is not an object oriented language. However , embedded structs provide
a kind of data-only inheritance that can be useful at times.Keep in mind
Go doesn't support classes or inheritance in the complete sense.
Embedded structsa are juat a way to elevate and share fields between struct
definitions*/

package main 
import ("fmt")

type Car struct {
	make string 
	model string
	truck_info  struct{
		chesis int 
		owner string
	}
}

type Truck struct{
	Car   // here we can see that We are doing kind of the inheritance but since GO is not OOPS we cannot say it as inheritance is just reusability
	LoadingCapacity int
}

func main(){
	truck := Truck{}
	truck.make = "TATA"
	truck.model = "Highway"
	truck.LoadingCapacity = 1200
	truck.truck_info.chesis = 452368
	truck.truck_info.owner = "Aditya Kshatriya"

	fmt.Println(
		truck.make , 
		truck.model , 
		truck.LoadingCapacity , 
		truck.truck_info.chesis , 
		truck.truck_info.owner ,
	)
}

 

