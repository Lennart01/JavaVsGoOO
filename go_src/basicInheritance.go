package main

import "fmt"

// Define a Vehicle interface that is implemented by Car and all its subtypes
type Vehicle interface {
	VehicleType()
	Brand()
}

// Car struct represents a generic car
type Car struct{}

// implement Vehicle interface for Car
func (c Car) VehicleType() {
	fmt.Println("Vehicle Type: Car")
}

func (c Car) Brand() {
	fmt.Println("Brand: Generic")
}

// Porsche struct embeds Car
type Porsche struct {
	Car
}

// Brand method for Porsche overrides Car's Brand method
func (p Porsche) Brand() {
	fmt.Println("Brand: Porsche")
}

// Speed method specific to Porsche (not present in Car)
func (p Porsche) Speed() {
	fmt.Println("Max: 200kmph")
}

// Audi struct embeds Car
type Audi struct {
	Car
}

// Brand method for Audi overrides Car's Brand method
func (a Audi) Brand() {
	fmt.Println("Brand: Audi")
}

// Speed method specific to Audi (not present in Car)
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
}
