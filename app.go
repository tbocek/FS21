package main

import (
	"database/sql"
	"fmt"
	"time"

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

	err = db.Ping()
	now := time.Now()
	for err != nil && now.Add(time.Duration(10)*time.Second).After(time.Now()) {
		// check the connection
		err = db.Ping()
		time.Sleep(time.Second)
	}
	if err != nil {
		panic(err)
	}

	sql := "create table if not exists test (id int, name varchar(255))"
	_, err = db.Exec(sql)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created Table!")
}
