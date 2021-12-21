package main

import (
	"Library/handlers"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main(){

	db, err := sql.Open("mysql", "mysql:@tcp(127.0.0.1:3306)/library")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	handlers.Handlers()
}
