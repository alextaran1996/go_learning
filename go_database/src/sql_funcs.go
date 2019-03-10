package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DBParams is params for workign with database
type DBParams struct {
	Serveraddr string `yaml:"Server_address"`
	Port       int    `yaml:"Port"`
	DB         string `yaml:"Database_name"`
	Table      string `yaml:"Table_name"`
	User       string `yaml:"User"`
	Passwd     string `yaml:"Password"`
}

// User is a schema of existing databaase
type User struct {
	ID   int
	Name string
}

// Create user
func createuser(conn *sql.DB, db *DBParams, user *User) {
	query := fmt.Sprintf("INSERT into %s.%s values(%d,\"%s\")", db.DB, db.Table, user.ID, user.Name)
	res, err := conn.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)

}

// Get all users
func getusers(conn *sql.DB, db *DBParams) {
	query := fmt.Sprintf("select * from %s.%s", db.DB, db.Table)
	res, err := conn.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()
	selectusr := []User{}
	for res.Next() {
		var usr User
		err := res.Scan(&usr.ID, &usr.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		selectusr = append(selectusr, usr)
	}
	for _, i := range selectusr {
		fmt.Println(i.ID, " ", i.Name)
	}

}

// Delete value with specified ID
func deleteID(conn *sql.DB, db *DBParams, id int) {
	query := fmt.Sprintf("DELETE FROM %s.%s where id = %d", db.DB, db.Table, id)
	_, err := conn.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("id%d was removed\n", id)
}

// Update value with specified ID
func updateID(conn *sql.DB, db *DBParams, id int, name string) {

	query := fmt.Sprintf("update %s.%s set name = \"%s\" where id = %d", db.DB, db.Table, name, id)
	_, err := conn.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("id%d was changed\n", id)
}
