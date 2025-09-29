/*Interfaces are collections of method signatures.A type  "implements" an
interface if it has all of the methods of given interface defined in it.*/
package main

import (
	"fmt"
)

type shape interface{
	area() float64
	perimeter() float64
}
type rect struct{
	width , height float64
}
type circle struct{
	radius float64
}
func (r rect) area() float64{
	return r.width*r.height
}
func (r rect) perimeter() float64{
	return 2*r.width+2*r.height
}

func (c circle) area() float64{
	return c.radius*c.radius
}

func main(){
	r1:=rect{
		width : 5.413,
		height : 9.32,
	}
	r2:=rect{
		width : 0,
		height : 1.235,
	}
	c1:= circle{
		radius : 5.36,
	}

	fmt.Println(r1.area() , r2.perimeter() , c1.area())
}