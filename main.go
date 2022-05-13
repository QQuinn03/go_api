package main

import (
	"fmt"
	"person"
)


func main() {
	fmt.Println("hello I am testing how to import packages")

	type QQ struct{	
       
	}

	func (*QQ) eat() {
		fmt.Println("she eats")
	}
	func (*QQ) drink() {
		fmt.Println("she drinks")
	}

	func (*QQ) drink() {
		fmt.Println("she drinks")
	}




	newWoman := NewPerson(QQ)

}


