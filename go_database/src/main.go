package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	db := readconfig("../config.yml")
	// db := DBParams{Serveraddr: "172.17.0.2", Port: 3306, User: "root", Passwd: "root", DB: "golang", Table: "users"}
	dbparam := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db.User, db.Passwd, db.Serveraddr, db.Port, db.DB)
	conn, err := sql.Open("mysql", dbparam)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// alex := User{Name: "Aliaksandr"}
	// createuser(&db, &alex)
	// deleteID(&db, 3)
	// updateID(conn, &db, 2, "Zoom")
	getusers(conn, db)
}
