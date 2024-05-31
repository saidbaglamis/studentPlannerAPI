package main

import (
	"studentPlanner/internal/database"
	"studentPlanner/internal/routes"
)

func main() {
	database.InitDB()
	routes.InitAllRoutes()
}
