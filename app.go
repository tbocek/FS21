package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
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
	db  *sql.DB
	rnd int
)

func dbQuery(w http.ResponseWriter, r *http.Request) {
	sqlStatement := `SELECT id, name FROM test`
	var name string
	var id int
	if rnd == 0 {
		rand.Seed(time.Now().Unix())
		rnd = rand.Int()
	}
	row := db.QueryRow(sqlStatement)
	log.Printf("rnd: %v", rnd)
	switch err := row.Scan(&id, &name); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		w.Write([]byte("DB: " + name + " / " + strconv.Itoa(rnd)))
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

	server := &http.Server{Addr: ":8080", Handler: http.HandlerFunc(dbQuery)}
	//server.SetKeepAlivesEnabled(false)
	server.ListenAndServe()

}
