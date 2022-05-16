package person

import "fmt"

type newConsequence interface {
	Drink()
	Eat()
	Sleep()
}
type behaviorConsequence struct {
	behConse newConsequence
}

func NewBhavior(nc newConsequence) *behaviorConsequence {
	return &behaviorConsequence{
		behConse: nc,
	}
}
func (bc *behaviorConsequence) Drink() {
	fmt.Println("a person drinks")
	bc.behConse.Drink()
}
func (bc *behaviorConsequence) Eat() {
	fmt.Println("a person eats")
	bc.behConse.Eat()
}
func (bc *behaviorConsequence) Sleep() {
	fmt.Println("a person sleeps")
	bc.behConse.Sleep()
}
