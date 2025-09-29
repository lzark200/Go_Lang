package main

import "fmt"

type car struct {
	Model       string
	Height      int
	width       int
	max_speed   float64
	milage_kmpl float64
	FrontWheel  wheel
	backWheel   wheel
	frontlights headlight
	backlights  headlight
}

// Nested struct inside the struct:
type wheel struct {
	Radius   float64
	Material string
}

type headlight struct {
	Category    string
	Color       string
	FieldOfView int
	Strength    float64
}

func test(m car) {
	fmt.Println(m.Model)
	fmt.Println(m.Height)
	fmt.Println(m.width)
	fmt.Println(m.max_speed)
	fmt.Println(m.milage_kmpl)
	fmt.Println(m.FrontWheel.Radius)
	fmt.Println(m.FrontWheel.Material)
	fmt.Println(m.frontlights.Category)
	fmt.Println(m.backlights.Category)
}

func main() {
	test(car{
		Model:       "tesla",
		Height:      20,
		width:       40,
		max_speed:   170.23,
		milage_kmpl: 10,
		FrontWheel: wheel{
			Radius:   45,
			Material: "plastic",
		},
		backWheel: wheel{
			Radius:   45,
			Material: "plastic",
		},
		frontlights: headlight{
			Category:    "small",
			Color:       "yellow",
			FieldOfView: 80,
			Strength:    65.32,
		},
		backlights: headlight{
			Category:    "small",
			Color:       "yellow",
			FieldOfView: 80,
			Strength:    65.32,
		},
	})

	fmt.Println("-----------------------------------------------------------------------------")
	
	test(car{
		Model:       "Maruti",
		Height:      10,
		width:       10,
		max_speed:   120,
		milage_kmpl: 30,
		FrontWheel: wheel{},
		backWheel: wheel{},
		frontlights: headlight{},
		backlights: headlight{},
	})

	
}



