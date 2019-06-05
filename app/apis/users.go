package apis

import (
	"encoding/json"
	"net/http"

	"../models/users"
)

// ListUsersHandlers returns the list of user contained in database
func ListUsersHandlers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	list := users.ListUsers()
	json.NewEncoder(w).Encode(list)
}

// CreateUserHandlers creates a new user in the database and returns the user
func CreateUserHandlers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	users.CreateUser()
	var user users.User
	json.NewEncoder(w).Encode(&user)
}
