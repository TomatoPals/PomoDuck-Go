package API

import (
	"github.com/TomatoPals/Pomoduck-Go/controllers"
	"github.com/gorilla/mux"
)

//routes for CRUD methods
func AddTasksHandler(r *mux.Router) {

	r.HandleFunc("/tasks", userController.GetUsers).Methods("GET")
	r.HandleFunc("/tasks/{id}", userController.GetUser).Methods("GET")
	r.HandleFunc("/tasks", userController.CreateUser).Methods("POST")
	r.HandleFunc("/tasks/{id}", userController.UpdateUser).Methods("PUT")
	r.HandleFunc("/tasks/{id}", userController.DeleteUser).Methods("DELETE")

}
