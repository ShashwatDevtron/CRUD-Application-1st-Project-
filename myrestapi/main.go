package main
import (
	"encoding/json" // We are working with JSON so we need to bring this inbuilt package
	"log"           //logs errors
	"math/rand"     // to generate random number
	"net/http"      //working with http -> used to create Api's
	"strconv"       // will be used to used to convert integer to string
	"github.com/gorilla/mux" // other packages are imported but this is not imported as it is not an inbuilt package
)

// A struct is like a class in object oriented programming.
// It has properties and methods like one in Java and C++

// User Struct (Model)
type User struct {
	ID     string  `json:"id"`
	Occupation   string  `json:"occupation"`
	Name  string  `json:"name"`
	Contact *Contact `json:"contact"` // we will create another structure of contact 
}

// Contact Struct
type Contact struct {
	Phone string `json:"phone"`
	Email  string `json:"email"`
}

// Init users var as a slice Book struct
// a slice is an array with variable length
var users []User

// every rout handler has to have a response and a request
// get all users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// get a single user
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//getting the user by searching id
	params := mux.Vars(r) //get parameters
	// Loop through users and find with ID
	for _, item := range users { // this is another way of writing for loop
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

// create a new user
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(1000000)) //Mock ID - not safe
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

// update a user
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range users{
		if item.ID == params["id"]{
		users = append(users[:index], users[index+1:]...)
		w.Header().Set("Content-Type", "application/json")
		var user User
		_ = json.NewDecoder(r.Body).Decode(&user)
		user.ID = params["id"]
		users = append(users, user)
		json.NewEncoder(w).Encode(user)
	return
		}
	}
	json.NewEncoder(w).Encode(users)
}

// delete a user
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range users{
		if item.ID == params["id"]{
		users = append(users[:index], users[index+1:]...)
		break
		}
	}
	json.NewEncoder(w).Encode(users)
}


func main(){
	//Init the mux router
	r := mux.NewRouter()

	// creating mock data
	users = append(users, User{ID: "1", Occupation: "farmer", Name: "Prerit", Contact: &Contact{Phone: "9090756467", Email: "prerit@gmailcom"}})
	users = append(users, User{ID: "2", Occupation: "businessman", Name: "Jay", Contact: &Contact{Phone: "9789556467", Email: "jay@gmailcom"}})
	//creating router handlers which will establish endpoints for our api's
	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/api/users/{id}", createUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r)) //To run the server. Here 8000 is the port number

}
