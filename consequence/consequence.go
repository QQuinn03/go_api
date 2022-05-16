package consequence

import (
	"fmt"
)

type person struct {
	name string
}

func NewPerson(n string) *person {
	return &person{
		name: n,
	}
}

func (p *person) Eat() {
	fmt.Println("so he has energies")
}
func (p *person) Drink() {
	fmt.Println("so he is not thirsty")
}

func (p *person) Sleep() {
	fmt.Println("so he is not tired")
}
