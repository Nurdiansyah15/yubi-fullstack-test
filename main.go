package main

import (
	"yubi-fullstack-test/database"
	"yubi-fullstack-test/routes"

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

	// seeder
	// seeders.SeedSalesOrders()
	// seeders.SeedSoDts()

	// start gin
	r := gin.Default()

	// cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:5173",
	}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	r.Use(cors.New(config))

	// routes
	routes.InitRoutes(database.DB, r)

	// start server
	r.Run(":8080")

}
