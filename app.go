package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "mydb"
)

var (
	db *sql.DB
)

func dbQuery(w http.ResponseWriter, r *http.Request) {
	sqlStatement := `SELECT id, name FROM test`
	var name string
	var id int
	row := db.QueryRow(sqlStatement)
	switch err := row.Scan(&id, &name); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		w.Write([]byte("DB: " + name))
	default:
		panic(err)
	}
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	now := time.Now()
	for err != nil && now.Add(time.Duration(30)*time.Second).After(time.Now()) {
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

	sql = "insert into test values (1, 'hallo')"
	_, err = db.Exec(sql)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created Table!")

	http.HandleFunc("/", dbQuery)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
