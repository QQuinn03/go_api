package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx" //you can think  jmoiron/sqlx privode ways for sql query, jmoiron always
	//go with lib/pq if data is json
	_ "github.com/lib/pq" //A pure Go postgres driver for Go's database/sql package,
	/*Notice that we’re loading the driver anonymously, aliasing its package qualifier
	 to _ so none of its exported names are visible to our code. Under the hood, the driver
	 registers itself as being available to the database/sql package,
	but in general nothing else happens with the exception that the init function is run.*/)

type Database struct {
	Client *sqlx.DB
}

func NewDatabase() (*Database, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_TABLE"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)

	dbConn, err := sqlx.Connect("postgres", connectionString) //sqlx connect(driver name, string) and return a dbbase
	if err != nil {
		return &Database{}, err
	}
	return &Database{Client: dbConn}, nil
}

func (d *Database) Ping(ctx context.Context) error {
	return d.Client.DB.PingContext(ctx)
}

//docker exec -it postgres id bash
//psql -U postgres
//check table \d+ table name
//6a2b46dd-5ca3-4a00-9d4e-162f900828d0
