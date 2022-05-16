package main

import (
	"fmt"

	consequence "github.com/QQuinn03/test/consequence"
	do "github.com/QQuinn03/test/person"
)

type newWoman struct {
	name string
}

func run() {
	fmt.Println("setting up the test")
	//p:=person.person{"QQ"}
	p := consequence.NewPerson("QQ")

	personBehavior := do.NewBhavior(p)
	personBehavior.Eat()
	personBehavior.Drink()
	personBehavior.Sleep()
}
func main() {
	fmt.Println("hello I am testing how to import packages")
	run()

}
