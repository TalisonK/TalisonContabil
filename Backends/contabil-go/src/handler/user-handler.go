package handler

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/model"
	"github.com/TalisonK/TalisonContabil/src/util"
)

// GetUsers retrieves all users from both databases
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Get users from database

	users, err := model.GetUsers()

	if err != nil {
		util.LogHandler("Failed to get users", err, "handler.GetUsers")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Failed to get users")
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
		return
	}

}

// CreateUser creates a user in both databases
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := new(domain.User)

	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		util.LogHandler("Failed to parse body.", err, "createUser")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Failed to parse body")
		return
	}

	result, err := model.CreateUser(user)

	if err != nil {
		util.LogHandler("Failed to add the new user", err, "handler.CreateUser")
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
		return
	}
}

// UpdateUser updates a user by id in both databases
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Get user from request
	user := new(domain.User)

	json.NewDecoder(r.Body).Decode(user)

	if user.ID == "" {
		util.LogHandler("Empty request body", nil, "updateUser")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Empty request body")
		return
	}

	result, err := model.UpdateUser(user)

	if err != nil {
		util.LogHandler("Failed to update user", err, "handler.UpdateUser")
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
		return
	}
}

// DeleteUser deletes a user by id in both databases
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Get user from request
	id := r.PathValue("id")

	if id == "" {
		util.LogHandler("Empty id passed.", nil, "deleteUser")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Empty id passed.")
		return
	}

	err := model.DeleteUser(id)

	if err != nil {
		util.LogHandler("Failed to delete user", err, "handler.DeleteUser")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Failed to delete user")
		return
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "User deleted successfully")
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	// Get user from request
	user := new(domain.User)
	json.NewDecoder(r.Body).Decode(user)

	if user.Name == "" || user.Password == "" {
		util.LogHandler("Empty request body", nil, "login")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Empty request body")
		return
	}

	result, err := model.LoginUser(*user)

	if err != nil {
		util.LogHandler("Failed to login user", err, "handler.Login")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Failed to login user")
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
		return
	}

}
