package main

import "fmt"

type Passport struct {
	Photo []byte
	Name string
	Surname string
	DateOfBirth string
}

type Vehicle struct {
	numberOfWheels int
}

type Truck struct {
	Vehicle
	CarryWeight int
}

func main() {
	t := new(Truck)
	fmt.Println(t.numberOfWheels)
}
