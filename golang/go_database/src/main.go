package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	config := flag.String("c", "config.yml", "Path to to config.yml file")
	newid := flag.String("n", "", "Create new record.Use -n Hleb")
	updateid := flag.String("u", "", "Update existing value. Use -u 2,Ann")
	deleteid := flag.Int("d", 0, "Delete specified id.Use -d 1")
	getuser := flag.String("g", "", "Get all records.Use -g")
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
	if *deleteid != 0 {
		deleteID(conn, db, *deleteid)
	}
	if *newid != "" {
		name := User{Name: *newid}
		createuser(conn, db, &name)
	}
	if *updateid != "" {
		fmt.Println(*updateid)
		id := strings.Split(*updateid, ",")
		idval, err := strconv.Atoi(id[0])
		if err != nil {
			log.Fatal("Can not convert string.See examles")
		}
		updateID(conn, db, idval, id[1])
	}
	if *getuser == "yes" {
		getusers(conn, db)
	}

}
