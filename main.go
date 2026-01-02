package main

import (
	"log"

	"backend/config"
	"backend/database"
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()
	database.Connect()

	r := gin.Default()

	routes.AuthRoutes(r)

	log.Fatal(r.Run(":8080"))
}
