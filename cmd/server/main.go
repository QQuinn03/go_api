package main

import (
	"fmt"

	"github.com/QQuinn03/go_api/internal/comment"
	"github.com/QQuinn03/go_api/internal/db"
	httpHandler "github.com/QQuinn03/go_api/internal/transport/http"
)

func Run() error {
	db, err := db.NewDatabase()
	if err != nil {
		return err
	}
	if err := db.MigrateDB(); err != nil {
		return err
	}
	//fmt.Println("database migrates successfully")
	//fmt.Println("database connects successfully")
	commentService := comment.NewService(db)
	handler := httpHandler.NewHandler(commentService)
	handler.MapRoutes()
	if err := handler.Serve(); err != nil {
		return err
	}

	return nil
}
func main() {
	fmt.Println("running up our application")
	if err := Run(); err != nil {
		fmt.Println(err)
	}

}
