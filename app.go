package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "172.17.0.1"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "mydb"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sql := "create table if not exists test (id int, name varchar(255))"
	_, err = db.Exec(sql)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created Table!")
}
