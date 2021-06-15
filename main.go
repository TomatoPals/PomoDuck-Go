package main

import (
	"github.com/TomatoPals/Pomoduck-Go/routes"
	"github.com/TomatoPals/Pomoduck-Go/util"
)

func main() {

	//checks the environment and then uses env variables to establish db connection
	config.InitialMigration()

	//creates a new mux router, mounts routes, and then listening
	router.InitializeRouter()

}
