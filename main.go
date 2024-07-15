package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

var connection *sql.DB

const connectionString = "host=127.0.0.1 port=5432 user=postgres password=12345 dbname=postgres sslmode=disable"

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	router := gin.Default()

	router.GET("/", handlerIndex)

	router.Run(":8080")
}

func handlerIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"Error": nil,
	})
}
