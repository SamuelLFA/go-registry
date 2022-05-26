package main

import (
	database "github.com/SamuelLFA/goregister/config/database"
	"github.com/SamuelLFA/goregister/routes"
)

func main() {
	database.ConectWithDatabase()
	routes.HandleRequest()
}
