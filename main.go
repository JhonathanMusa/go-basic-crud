package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// array that pretends to be a database
type Users struct {
	ID        string `json: "id", omitempty`
	FirstName string `json: "firstname", omitempty`
	LastName  string `json: "lastname", omitempty`
}

var users []Users

func GetUsers(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func CreateNewUser(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var user Users

	_ = json.NewDecoder(req.Body).Decode(&user)
	user.ID = params["id"]
	users = append(users, user)
	json.NewEncoder(w).Encode(users)
}

func main() {

	router := mux.NewRouter()

	users = append(users, Users{ID: "1", FirstName: "jhonathan", LastName: "salazar"})

	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/new-user/{id}", CreateNewUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))

}
