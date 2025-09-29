package main 
import "fmt"

type Person struct {
	Name string
	Age int
	IsGay bool
	Height int

	// declaring the nested struct
	Working worklife
	
	// declaring the nested anonymous struct
	Parents struct {
		father string
		mother string 
	}
}
type worklife struct { 
	Company string
	WorkingHrs int
}

// anonymous struct : A struct without the type and name like the old syntax
var myCar = struct{
	Make string
	Model string
}{
	Make : "tesla",
	Model: "model 3",
}

func main(){
	var p1 Person
	p1.Name = "aditya"
	p1.Age = 25
    p1.Working.Company = "Google"
	fmt.Println(p1.Name , p1.Age , p1.IsGay , p1.Height , p1.Working.Company , p1.Working.WorkingHrs)

	// we can also call and make struct object in this way:
	p2:=Person{}
	p2.Name = "Aditya"
	p2.Working.Company = "Google"
	fmt.Println(p2.Name , p2.Age , p2.IsGay , p2.Height , p2.Working.Company , p2.Working.WorkingHrs)
	fmt.Println(myCar.Make , myCar.Model)
	p2.Parents.father = "kundan"
	p2.Parents.mother = "babita"

	fmt.Println(p2.Parents.father , p2.Parents.mother)
}


