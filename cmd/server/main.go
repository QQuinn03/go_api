package main

import (
	"context"
	"fmt"

	"github.com/QQuinn03/api_toy/internal/db"
)

func run() error {
	db, err := db.NewDatabase()
	if err != nil {
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("database connects successfully")
	return nil
}
func main() {
	fmt.Println("hello I am testing how to import packages")
	run()

}
