package main

import "fmt"

type Shape interface {
	area() float64
}

type rectangle struct {
	length, breadth float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}

func printLength(r rectangle) { // you can only Pass r
	fmt.Printf("Length: %.2f\n", r.length)
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func calculate(s Shape) float64 { // You can Pass R , c or s
	return s.area()
}

/*
 **********

            generics or empty Interfaces are Needed For Creating mixed Slices,reading UnPredictable Json data,
            or in Writing generics Functions

 ***********
   // 1 ) mixed Slices
func main() {
	values := []interface{}{123, "hello", true, 45.6}

	for _, v := range values {
		fmt.Printf("Value: %v, Type: %T\n", v, v)
	}
}

 // 2 ) Reading unpredictable Json Data
	jsonData := []byte(`{"name": "Nishant", "age": 25, "active": true}`)

	var result map[string]interface{} // 👈 empty interface used here


 // 3 )  Writing Generic Functions

func PrintAnything(value interface{}) {
	fmt.Println("You passed:", value)
}

func main() {
	PrintAnything(42)
	PrintAnything("Hello")
	PrintAnything(true)
}

*/

func CheckTypeOfx(x interface{}) { // type assertion method // simple and useful
	if _, ok := x.(rectangle); ok {
		fmt.Println("rectangle")
	} else {
		fmt.Println("No Rectangle")
	}
}

func main() {
	var x interface{}
	// type y interface{}
	x = rectangle{10, 40}
	r := rectangle{10, 20}
	c := circle{2.0}
	var s Shape

	s = r

	fmt.Println(x)
	//fmt.Println(y)
	fmt.Println(s.area())

	fmt.Println(calculate(r))
	fmt.Println(calculate(c))

	CheckTypeOfx(x) // this is known as Type Assertion when we are using an empty interface{} , and we need
	// to check which data type exactly the var is holding right now  so we do this

	printLength(r)
}
