package person

import "fmt"

type newPerson struct {
	name string
}

func (p *newPerson) drink() {
	fmt.Println("a person drinks")
}
func (p *newPerson) eat() {
	fmt.Println("a person eats")
}
func (p *newPerson) sleep() {
	fmt.Println("a person sleeps")
}
