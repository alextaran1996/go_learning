package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DBParams is params for workign with database
type DBParams struct {
	Serveraddr string
	Port       int
	DB         string
	Table      string
	User       string
	Passwd     string
}

// User is a schema of existing databaase
type User struct {
	ID   int
	Name string
}

func createuser(db *DBParams, user *User) {
	dbparam := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db.User, db.Passwd, db.Serveraddr, db.Port, db.DB)
	conn, err := sql.Open("mysql", dbparam)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	query := fmt.Sprintf("INSERT into %s.%s values(%d,\"%s\")", db.DB, db.Table, user.ID, user.Name)
	fmt.Println(query)
	res, err := conn.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

}

func main() {
	db := DBParams{Serveraddr: "172.17.0.2", Port: 3306, User: "root", Passwd: "root", DB: "golang", Table: "users"}
	alex := User{ID: 4, Name: "Aliaksandr"}
	createuser(&db, &alex)
}
