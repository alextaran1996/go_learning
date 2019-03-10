package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
)

func main() {
	config := flag.String("c", "config.yml", "Path to to config.yml file")
	// newid := flag.String("n", "", "Create new record.Use -n Hleb")
	// updateid := flag.String("u", "", "Update existing value. Use -u 2 Ann")
	// deleteid := flag.String("d", "", "Delete specified id.Use -d 1")
	// getusers := flag.String("g", "", "Get all records.Use -g")
	flag.Parse()
	db, err := readconfig(*config)
	if err != nil {
		log.Fatal(err)
	}
	dbparam := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db.User, db.Passwd, db.Serveraddr, db.Port, db.DB)
	conn, err := sql.Open("mysql", dbparam)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	name := User{Name: "German"}
	createuser(conn, db, &name)
	deleteID(conn, db, 1)
	// updateID(conn, &db, 2, "Zoom")
	getusers(conn, db)
}
