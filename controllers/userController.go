package userController

import (
	"encoding/json"
	"net/http"

	"github.com/TomatoPals/Pomoduck-Go/models"
	"github.com/TomatoPals/Pomoduck-Go/util"

	"github.com/gorilla/mux"
)

//gets a single user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user models.User
	config.DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

//gets all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	config.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

//creates a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	config.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

//updates a user by ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user models.User
	config.DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	config.DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

//creates a timestamp in deleted_at column
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user []models.User
	config.DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The User is Deleted Successfuly")
}
