package API

import (
	"github.com/TomatoPals/Pomoduck-Go/controllers"
	"github.com/gorilla/mux"
)

//routes for CRUD methods
func AddUserHandler(r *mux.Router) {

	r.HandleFunc("/users", userController.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", userController.GetUser).Methods("GET")
	r.HandleFunc("/users", userController.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

}
