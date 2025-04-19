package main

import (
	"yubi-fullstack-test/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// load env file
	if err := godotenv.Load(); err != nil {
		panic("Failed to load env file: " + err.Error())
	}

	//  connect database
	if err := database.ConnectDatabase(); err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// start gin
	r := gin.Default()

	// cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:3000",
	}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(config))

	// routes

	// start server
	r.Run(":8080")

}
