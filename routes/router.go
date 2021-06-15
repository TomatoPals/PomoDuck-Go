package router

import (
	"log"
	"net/http"

	"github.com/TomatoPals/Pomoduck-Go/routes/api"
	"github.com/gorilla/mux"
)

//creates new router with mounted routes
func InitializeRouter() {
	r := mux.NewRouter()

	API.Mount(r, "/v1", API.Routes())

	log.Fatal((http.ListenAndServe(":8080", r)))
}
