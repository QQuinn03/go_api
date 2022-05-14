package main

import (
	"fmt"

	ppl "github.com/QQuinn03/test/person1"
)

func main() {
	/* fmt.Println("hello I am testing how to import packages")

	type qq struct{
	    QQ string;
	}

	func (*qq) eat() {
		fmt.Println("she eats")
	}
	func (*qq) drink() {
		fmt.Println("she drinks")
	}

	func (*qq) drink() {
		fmt.Println("she drinks")
	}

	newWoman := NewPerson(&man)*/
	c := ppl.Hello()
	fmt.Println(c)
}
