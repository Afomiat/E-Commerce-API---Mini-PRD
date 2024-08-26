package main

import (
	"time"

	"github.com/Afomiat/E-Commerce-API---Mini-PRD/config" // Correct import path
	"github.com/Afomiat/E-Commerce-API---Mini-PRD/delivery/router"
	"github.com/gin-gonic/gin"
)

func main() {
    // Load environment variables
    env := config.LoadEnv()

    // Connect to MongoDB
    dbClient := config.ConnectMongoDB(env) // Assuming ConnectMongoDB is in the config package
    db := dbClient.Database(env.DBName)

    // Set up server timeout
    timeout := time.Duration(env.ContextTimeout) * time.Second

    // Initialize Gin router
    gin := gin.Default()

    // Set up routes
    router.Setup(env, timeout, db, gin)

    // Start the server
    gin.Run(env.LocalServerPort)
}
