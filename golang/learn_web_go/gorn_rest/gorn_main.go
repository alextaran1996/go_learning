package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Worker defenition of the worker
type Worker struct {
	gorm.Model
	Name string
	Age  int
}

// Migration will connect to database and migrate structure in it
func Migration() {
	dbparam := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", "root", "root", "172.17.0.2", 3306, "gorm")
	db, err := gorm.Open("mysql", dbparam) // Open doesn't connect ot database it's only create new connection with specified credentials
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}
	defer db.Close()
	db.AutoMigrate(&Worker{}) // Reference to struture that will be migrated to mysql
}

// HandleRequestsServ will handle all requests and communicate with DB
func HandleRequestsServ() {
	MyRouter := mux.NewRouter().StrictSlash(true) // StrictSlash() allow server to percive /example and /example/ as request to the same resource
	MyRouter.HandleFunc("/users", allusers).Methods("GET")
	MyRouter.HandleFunc("/user/{name}", deleteuser).Methods("DELETE") // {name} will be interpreted as a variable until next slash
	MyRouter.HandleFunc("/user/{name}/{age}", createuser).Methods("POST")
	MyRouter.HandleFunc("/user/{name}/{age}", updateuser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", MyRouter))
}

func allusers(w http.ResponseWriter, r *http.Request) {
	dbparam := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", "root", "root", "172.17.0.2", 3306, "gorm")
	db, err := gorm.Open("mysql", dbparam) // Change databse url
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}
	defer db.Close()
	var workers []Worker
	db.Find(&workers)
	json.NewEncoder(w).Encode(workers)
}

func deleteuser(w http.ResponseWriter, r *http.Request) {
	dbparam := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", "root", "root", "172.17.0.2", 3306, "gorm")
	db, err := gorm.Open("mysql", dbparam) // Change databse url
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}
	defer db.Close()
	vars := mux.Vars(r)
	name := vars["name"]
	var worker Worker
	db.Where("name = ?", name).Find(&worker)
	db.Delete(&worker)
	fmt.Fprintf(w, "User deleted")

}

func createuser(w http.ResponseWriter, r *http.Request) {
	dbparam := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", "root", "root", "172.17.0.2", 3306, "gorm")
	// ?parseTime=true. The default internal output type of MySQL DATE and DATETIME values is []byte which allows you to scan the value into a []byte, string or sql.RawBytes variable in your program.
	// You can do that by changing the internal output type from []byte to time.Time with the DSN parameter parseTime=true.
	// Need this because gorm.Model contains fields time.Time type but default mysql driver returns []uint8
	db, err := gorm.Open("mysql", dbparam) // Change databse url
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}
	defer db.Close()
	vars := mux.Vars(r)
	name := vars["name"]
	age, _ := strconv.Atoi(vars["age"])
	db.Create(&Worker{Name: name, Age: age})
	fmt.Fprintf(w, "User created")
}

func updateuser(w http.ResponseWriter, r *http.Request) {
	dbparam := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", "root", "root", "172.17.0.2", 3306, "gorm")
	db, err := gorm.Open("mysql", dbparam) // Change databse url
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}
	defer db.Close()
	var worker Worker
	vars := mux.Vars(r)
	db.Where("name = ?", vars["name"]).Find(&worker) // Find all the user with specified name add put all the info in worker variable
	worker.Name = vars["name"]                       // in the case user was not represented in the table
	worker.Age, _ = strconv.Atoi(vars["age"])
	db.Save(&worker)
	fmt.Fprintf(w, "User updated")

}

func main() {
	fmt.Println("Running gorn_main!")
	Migration()
	HandleRequestsServ()
}
