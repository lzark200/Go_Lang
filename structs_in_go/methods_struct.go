// Methods in the go has two things : 
// receiver variable r rect
// function area()

package main 
import (

	"fmt"
)

type rect struct{
	width float64
	height float64
}

// area has a receiver of (r rect)
func (r rect) area() float64{
	return r.width*r.height
}

func main(){
	r1 := rect{
		width : 5.363 ,
		height: 9.235 ,
	}
	r2 := rect{
		width :10.2356 , 
		height : 78.3659,
	}
	fmt.Println(r1.area() , r2.area())
}