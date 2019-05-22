package main

import "fmt"

type Passport struct {
	Photo []byte
	Name string
	Surname string
	DateOfBirth string
}

func (p Passport) GetFullName() string {
	return fmt.Sprintf("%s %s", p.Name, p.Surname)
}

func (p *Passport) ChangeFullName(name string, surname string) {
	p.Name = name
	p.Surname = surname
}

func main() {
	p1 := new(Passport)
	p1.ChangeFullName("Go", "Is Awesome")
	fmt.Println(p1.GetFullName())
}
