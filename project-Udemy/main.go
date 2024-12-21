package main

import (
	"github.com/Wong-bui/Udemy-project/db"
	"github.com/Wong-bui/Udemy-project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database := db.InitDB()

	db.SetDatabaseInstance(database)

	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
