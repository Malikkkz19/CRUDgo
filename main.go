package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

var users = []User{
	{ID: 1, Name: "Malik", Email: "malik@gmail.com", Role: "admin"},
	{ID: 2, Name: "Danil", Email: "danil@gmail.com", Role: "guess"},
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		return
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, user := range users {
		if user.ID == id {
			err := json.NewEncoder(w).Encode(user)
			if err != nil {
				return
			}
		}
	}
	http.NotFound(w, r)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = len(users) + 1
	users = append(users, user)
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		return
	}

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for index, user := range users {
		if user.ID == id {
			users = append(users[:index], users[index+1:]...)
			var updatedUser User
			_ = json.NewDecoder(r.Body).Decode(&updatedUser)
			updatedUser.ID = id
			users = append(users, updatedUser)
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}
	http.NotFound(w, r)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for index, user := range users {
		if user.ID == id {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		return
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}
