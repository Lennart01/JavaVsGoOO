package main

import "fmt"

// Define a Vehicle interface that is implemented by Car and all its subtypes
type Vehicle interface {
	VehicleType()
	Brand()
}

// Car struct represents a generic car
type Car struct{}

// Car now does not fully implement Vehicle interface (missing Brand method)
func (c Car) VehicleType() {
	fmt.Println("Vehicle Type: Car")
}

// Porsche struct embeds Car and provides a specific implementation for Brand
type Porsche struct {
	Car
}

// Brand method for Porsche
func (p Porsche) Brand() {
	fmt.Println("Brand: Porsche")
}

// Speed method specific to Porsche
func (p Porsche) Speed() {
	fmt.Println("Max: 200kmph")
}

// Audi struct embeds Car and provides a specific implementation for Brand
type Audi struct {
	Car
}

// Brand method for Audi
func (a Audi) Brand() {
	fmt.Println("Brand: Audi")
}

// Speed method specific to Audi
func (a Audi) Speed() {
	fmt.Println("Max: 180kmph")
}

// interfaces allow generic functions
func DescribeVehicle(v Vehicle) {
	v.VehicleType()
	v.Brand()
}

func main() {
	p1 := Porsche{}
	p1.Speed()
	DescribeVehicle(p1)

	a1 := Audi{}
	a1.Speed()
	DescribeVehicle(a1)

	c1 := Car{}
	c1.VehicleType()
}
