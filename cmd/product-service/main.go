package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/tarkour/product-service/config"
)

func main() {

	connstr := config.Get_connStr("./config/database_connection.yaml")

	db, err := sql.Open("postgres", connstr)

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

}
