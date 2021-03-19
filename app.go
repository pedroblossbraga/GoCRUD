package main

/*
A great tutorial: https://tutorialedge.net/golang/golang-orm-tutorial/
Go ORM querie tutorial : https://gorm.io/docs/query.html

 http://localhost:8081/users
 http://localhost:8081/user/pedro/braga/saopaulo/pedro@email.com
 http://localhost:8081/user/joca/caldeira/socorro/joca@script.net


*/
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Name  string
	LastName string
	City string
	Email string

}

func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hi! This is the homepage.\nYou can do CRUD operations with this Rest API :)\n")
	fmt.Fprintf(w, `
	
	/users -> GET (All Users)
	/user/{name}" -> "DELETE
	/user/{name}/{lastname}/{city}/{email} -> PUT
	/user/{name}/{lastname}/{city}/{email} -> POST
	
	`)
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New User Endpoint Hit")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	lastname := vars["lastname"]
	city := vars["city"]
	email := vars["email"]

	fmt.Println(name)
	fmt.Println(lastname)
	fmt.Println(city)
	fmt.Println(email)

	db.Create(&User{Name: name, LastName: lastname, City: city, Email: email})
	fmt.Fprintf(w, "New User Successfully Created")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	//lastname := vars["lastname"]
	//city := vars["city"]
	//email := vars["email"]

	var user User
	db.Where("name = ? ", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "Successfully Deleted User")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	lastname := vars["lastname"]
	city := vars["city"]
	email := vars["email"]

	var user User
	db.Where("name = ? ", name,).Find(&user)

	user.LastName = lastname
	user.Email = email
	user.City = city
	
	db.Save(&user)
	fmt.Fprintf(w, "Successfully Updated User")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage).Methods("GET")
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{lastname}/{city}/{email}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{name}/{lastname}/{city}/{email}", newUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func initialMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})
}

func main() {
	fmt.Println("Go ORM Tutorial")

	initialMigration()
	// Handle Subsequent requests
	handleRequests()
}