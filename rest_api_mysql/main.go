package main

import(

	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"fmt"
)
type student struct{
	ID string `json:"id"`
	NAME string `json:"name"`
}
var db *gorm.DB

func initDB() {
    var err error
	dataSourceNAME := "root:#Anjali123@tcp(127.0.0.1:3306)/"
	db, err = gorm.Open("mysql", dataSourceNAME)

	if err!= nil {
		fmt.Println(err)
        panic(err)
    }
	//db.Exec("CREATE DATABASE students_db")
	db.Exec("USE students_db")

	db.AutoMigrate(&student{})
}


func poststudent(w http.ResponseWriter, r *http.Request) {
	var newstudent student
	json.NewDecoder(r.Body).Decode(&newstudent)
	db.Create(&newstudent)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newstudent)
	fmt.Println("hello guys")
}

func getstudent(w http.ResponseWriter, r*http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var students []student
	db.Find(&students)
	json.NewEncoder(w).Encode(students)
}

	func main(){
		router:= mux.NewRouter()
		router.HandleFunc("/students",poststudent).Methods("POST")
		router.HandleFunc("/students",getstudent).Methods("GET")
		initDB()
		log.Fatal(http.ListenAndServe(":8000", router))
	}
