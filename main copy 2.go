package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Method using pointer to modify struct
func (p *Person) Birthday() {
	p.Age++ // Increase age by 1
}

func main() {
	p := Person{Name: "Leo", Age: 20}
	p.Birthday()
	fmt.Println(p.Age) // Output: 21
}
